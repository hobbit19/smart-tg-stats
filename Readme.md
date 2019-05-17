**Smart Tg Stats** is an Ethereum-based robot that accepts requests and responds with
statistic information regarding posts in Telegram Messenger.

# Interface
Here is the public contract interface:

```sol

  Request[]                       requests;

  event                       NewRequest(uint);
  event                       NewResponse(uint);

  function AddRequest(bytes32 _channel, uint32 _postID)
                        public payable returns (uint256 requestID);

  // use
  function GetLastRequestID() public view returns (uint256) ;

  bool                        status;      // must be true if robot is online
  uint256                     statusRenew; // timestamp; must be bigger than UTCNOW-24h
```

### Send Request
Use method `AddRequest() payable`:

- Channel username
- Post ID

Also records **Bid** - amount of ether sent with request, if any.
Bid affects prioritization of the request and speed of updates.

This amount of ether _may be returned_ to the caller (`msg.sender`) later - only, if there was no concurrency.
In general cases, bid is non-refundable.
Zero bids are fine but such requests may take infinity to be processed by the robot.

After transaction mined on the node,
immediately use getter `GetLastRequestID()` to get `uin256` handle of the created request
and store it.

### Response info
The data behind `Request` struct has packing

```sol
enum ProcErr {
  None,               // no error
  UnknownError,       // there was technical difficulties
  ChannelAccessError, // channel not found or bot blocked (temporarily or not)
  PostAccessError     // invalid post ID
}

struct Request {
    bytes32   channel;          /* Channel username */
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

### ABI
```
none
```

Mainnet address: `not deployed`
Rinkeby address: `not deployed`
