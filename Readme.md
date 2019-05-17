Ethereum service that accepts and processes requests for
stat information for posts in channels inside Telegram Messenger.

# Stats per request
- Subscribers count
- Post views count

# Interface

```sol

  Request[] public            Requests;
  uint256 Public              maxRequestID;

  event                       NewRequest(uint);
  event                       NewResponse(uint);

  function AddRequest(string _channel, uint32 _postID)
                        public payable returns (uint256 requestID);

  bool                        status;      // must be true if robot is online
  uint256                     statusRenew; // timestamp; must be bigger than UTCNOW-24h
```

# Request
Use method `AddRequest() payable`:

- Channel username
- Post ID

Also records **Bid** - amount of ether sent with request, if any.
Sending bid may affect prioritization of the request and speed of updates.

This amount of ether _may be returned_ to the caller (`msg.sender`) later - only, if there was no concurrency.
In general cases, bid is non-refundable.
Zero bids are fine but such requests may take infinity to be processed by the robot.

After transaction mined on the node,
immediately use getter `GetLastRequestID()` to get `uin256` handle of the created request
and store it.

# Response
The data behind `Request` is updated with generated event few times for few hours since request.

Use getters on your node to get struct with packing:

```sol
enum ProcErr {
  None,               // no error
  UnknownError,       // there was technical difficulties
  ChannelAccessError, // channel not found or bot blocked (temporarily or not)
  PostAccessError     // invalid post ID
}

struct Request {
    string    channel;          /* Channel username */
    uint32    postID;           /* Post ID */
    uint256   bid;              /* Ether given for priority */
    bool      willUpdate;       /* true: monitoring is ON for this request (updates are coming later) */
    uint256   timeCreated;      /* Timestamp of the request */
    uint256   timeFirst;        /* Timestamp of the first update of stats */
    uint256   timeLast;         /* Timestamp of the latest update of stats */
    uint32    lastSubscribers;  /* Stats: channel subscribes counter */
    uint32    lastPostViews;    /* Stats: post views counter */
    uint256   postTime;         /* Stats: UTC timestamp post created */
    ProcErr   error;            /* Response: robot last error (if any) */
}
```

# Public instance

Mainnet address:
`not deployed`

Rinkeby address:
`not deployed`

Frontend: https://crackhd.github.io/smart-tg-stats

# Deploy your instance

## Container
`git submodule update -r --init`
`docker build -t smart-tg-stats .`

## Local
Use scripts `./build.sh` (or `./build-mac.sh`) to install/build dependencies.

Then `make` to run the build and app locally.
`make sol` is needed once to re/generate Go bindings for contract.
