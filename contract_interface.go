package main

import (
  "math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"context"
  "fmt"
  "errors"
  "time"
  "sort"
)

type ProcErr uint8

const (

  // solidity:
  // << enum ProcErr { >>
  //
  ERROR_NONE      ProcErr = 0
  ERROR_UNKNOWN   ProcErr = 1
  ERROR_CHANACC   ProcErr = 2
  ERROR_POSTACC   ProcErr = 3

)

var (
  ErrNoCEO       = errors.New("No CEO session initialized")
  ErrNoBackend   = errors.New("No request processing backend is configured")
)

type contract_request struct {
  ID * big.Int
  Channel string
  PostID uint32
  Bid * big.Int
  WillUpdate bool
  TimeCreated * big.Int
  TimeUpdated * big.Int
  Subscribers uint32
  PostViews uint32
  PostTime * big.Int
  Error ProcErr
}

type contract_ceo interface {
  RenewStatus(is_on bool) error
  UpdateResponse(contract_request) error
}

type contract_backend interface {
  ProcessRequest(contract_request, contract_ceo) error
}

type contract_interface struct {
  node_address      string
  contract_address  string
  business_backend  contract_backend
  address           common.Address
  eth_client        * ethclient.Client
  eth_contract      * SmartTgStats
  ceo_session       * SmartTgStatsSession
}

func NewContractInterface(node_address string, contract_address string, backend contract_backend) * contract_interface {
  return &contract_interface{
    node_address: node_address,
    contract_address: contract_address,
    business_backend: backend,
  }
}

func (self * contract_interface) setup() error {

  	var err error

  	self.address = common.HexToAddress(self.contract_address)

  	self.eth_client, err = ethclient.Dial(self.node_address)
  	if err != nil {
  		return err
  	}

  	self.eth_contract, err = NewSmartTgStats(self.address, self.eth_client)
  	if err != nil {
  		return err
  	}

    fmt.Println("Contract is ready. Be ambitious!")

  	return nil
}

func (self * contract_interface) setup_ceo(ceo_keystore []byte, password string) error {

	key, err := keystore.DecryptKey(ceo_keystore, password)
  if err != nil {
    return err
  }

  auth := bind.NewKeyedTransactor(key.PrivateKey)

	gasPrice, err := self.eth_client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// If use something other than getters, ceoSession will be required
	self.ceo_session = &SmartTgStatsSession{
		Contract: self.eth_contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
			GasLimit: uint64(2750000),
		},
	}

  fmt.Println("Contract CEO is initialized.")

  return nil
}

func (self * contract_interface) RenewStatus(is_on bool) error {
  if self.ceo_session == nil {
    return ErrNoCEO
  }

  tx, err := self.ceo_session.CEORenewStatus(is_on)
  if err == nil {
    fmt.Printf("CEORenewStatus transaction with tx %s\n", tx.Hash().String())
  }

	return err
}


func (self * contract_interface) UpdateResponse(req contract_request) error {
  if self.ceo_session == nil {
    return ErrNoCEO
  }

  tx, err := self.ceo_session.CEOAddResponse(req.ID, req.WillUpdate, req.Subscribers, req.PostViews, req.PostTime, uint8(req.Error))
  if err == nil {
    fmt.Printf("CEOAddResponse transaction with tx %s\n", tx.Hash().String())
  }

  return err
}

func (self * contract_interface) renew_if_needed() bool {
  pubStatus, err := self.ceo_session.Status()
  if err == nil {

      pubStatusTime, err := self.ceo_session.StatusRenew()
      if err == nil {

        maxActualTime := big.NewInt(0)
        maxActualTime.Add(maxActualTime, pubStatusTime)
        maxActualTime.Add(maxActualTime, big.NewInt(86400 - 200))

        if !pubStatus || pubStatusTime.Cmp(maxActualTime) >= 0 {
          fmt.Printf("Status renew is required: status %d from %s\n", pubStatus, pubStatusTime.String())
           if err := self.RenewStatus(true); err == nil {
             time.Sleep(60)
           } else {
             fmt.Printf("[Warning] error transaction [CEORenewStatus] (retry): %s\n", err.Error())
             return false
           }
        }
      } else {
        fmt.Printf("[Warning] error getting contract [StatusRenew] (skip): %s\n", err.Error())
        return false
      }

  } else {
    fmt.Printf("[Warning] error getting contract [Status] (skip): %s\n", err.Error())
    return false
  }

  return true

}

func (self * contract_interface) poll_loop() {

  if self.ceo_session == nil {
    // Set up non-ceo session for getters?
    panic(ErrNoCEO)
  } else if self.business_backend == nil {
    panic(ErrNoBackend)
  }

  var (
      step = big.NewInt(1)
      min = big.NewInt(1)  // minimum valid request ID is 1, not 0
  )

  for {

    time.Sleep(70)

    if !self.renew_if_needed() {
      continue
    }

    // ======================== GET TOP REQUEST ID =============================

    max, err := self.ceo_session.MaxRequestID()

    if err != nil {
      fmt.Printf("MaxRequestID() ethereum getter error (stop): %s\n", err.Error())
      continue
    }

    if max.Cmp(min) == 0 {
      fmt.Println("There are no requests at all (retry after pause)")
      break
    }

    // ===================== BUILD ACTIVE REQUESTS LIST ========================

    requests := make([]contract_request, 0)

    for i := new(big.Int).Set(min); i.Cmp(max) <= 0; i.Add(i, step) {

      eth_req, err := self.ceo_session.Requests(i)
      if err != nil {
          fmt.Printf("Requests(%d) ethereum getter error (ignore): %s\n", i.Uint64(), err.Error())
          continue
      }

      if !eth_req.WillUpdate {
        // This request is archived
        continue
      }

      requests = append(requests, contract_request{
        ID:         i,
        Channel:    eth_req.Channel,
        PostID:     eth_req.PostID,
        Bid:        eth_req.Bid,
        WillUpdate: eth_req.WillUpdate,
        TimeCreated:eth_req.TimeCreated,
        TimeUpdated:eth_req.TimeLast,
        Subscribers:eth_req.LastSubscribers,
        PostViews:  eth_req.LastPostViews,
        PostTime:   eth_req.PostTime,
        Error:      ProcErr(eth_req.Error),
      })

    }

    if len(requests) == 0 {
      fmt.Println("There are no active requests (retry after pause)")
      break
    }

    fmt.Printf("Active requests: %d\n", len(requests))

    // ================== SORT REQUESTS BY BID AND TIME ========================

    sort.Slice(requests[:], func(i, j int) bool {
        return requests[i].Bid != nil &&
                requests[j].Bid != nil &&
                requests[i].Bid.Cmp(requests[j].Bid) >= 0
    })

    // =================== FEED REQUESTS TO THE BACKEND ========================

    for _, req := range requests {
      err := self.business_backend.ProcessRequest(req, self)
      if err != nil {
        fmt.Printf("Backend [ProcessRequest] failed (%d) (ignore): %s", req.ID.Uint64(), err.Error())
      }
    }

  }

}
