package main

import (
  "github.com/bsm/redis-lock"
  "github.com/go-redis/redis"
  "errors"
  "fmt"
  "time"
)

var (
  ErrLockedAlready = errors.New("The request ID is locked already")
  lockOptions = lock.Options{
    LockTimeout: time.Second * 60,
    RetryCount: 3,
    RetryDelay: time.Second * 3,
  }
)

type mutex_helper struct {
  keyPrefix       string
  client          * redis.Client
  IsFake          bool
}

type mutex struct {
  lock * lock.Locker
}

func NewMutexHelper(key_prefix, protocol, redis_address string) (* mutex_helper) {
  client := redis.NewClient(&redis.Options{
  		Network:	protocol,
  		Addr:		redis_address,
  	})

  return &mutex_helper{
    client:     client,
    keyPrefix:  key_prefix,
    IsFake:     len(redis_address) <= 1,
  }
}

func (self * mutex_helper) Connect() error {

  var err error

  for i := 0; i < 5; i++ {

    if _, err = self.client.Ping().Result(); err != nil {
      time.Sleep(time.Second * 3)
      fmt.Println("Retrying redis.Ping()...")
    } else {
      break
    }
  }

  return err
}

func (self * mutex_helper) LockOnRequest(requestId string) (mutex, error) {

  var res = mutex{}

  if self.IsFake {
    return res, nil
  }

  key := fmt.Sprintf("[%s]r:%s", self.keyPrefix, requestId)

	lock, err := lock.Obtain(self.client, key, &lockOptions)
	if err != nil {
		return res, err
	} else if lock == nil {
		return res, ErrLockedAlready
	}

  res.lock = lock

  return res, nil
}

func (self * mutex) Release() error {
  if self.lock == nil {
    return nil
  }
  return self.lock.Unlock()
}
