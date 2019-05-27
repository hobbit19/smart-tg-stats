pragma solidity ^0.5.8;

contract SmartTgStats {

  address                     owner;
  bool                 public status;
  uint256              public statusRenew;

  enum ProcErr {
    None,
    UnknownError,
    ChannelAccessError,
    PostAccessError
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

  mapping (uint256 => Request)    public Requests;
  uint256                         public maxRequestID;

  event                       NewRequest(uint);
  event                       NewResponse(uint);

  constructor(bool _isOn) public payable {
      require(owner == address(0));
      owner = msg.sender;
      status = _isOn;
      statusRenew = block.timestamp;
  }

  function CEORenewStatus(bool _isOn) public {
    require(owner != address(0));
    require(owner == msg.sender);
    status = _isOn;
    statusRenew = block.timestamp;
  }

  function AddRequest(string memory _channel, uint32 _postID) public payable returns(uint256 requestID) {
    require(owner != address(0));

    require(status == true);
    require(block.timestamp - statusRenew >= 0);
    require(block.timestamp - statusRenew < 1209600);

    maxRequestID += 1;
    requestID = maxRequestID;
    Requests[requestID] = Request(
          _channel,
          _postID,
          msg.value,
          true,
          block.timestamp, 0, 0,
          0, 0, 0,
          ProcErr.None
    );

    emit NewRequest(requestID);
  }

  function CEOAddResponse(uint _requestID, bool _willUpdate, uint32 _subscribers, uint32 _postViews, uint256 _postTime, ProcErr _error) public {
    require(owner != address(0));
    require(owner == msg.sender);
    require(_requestID > 0 && _requestID <= maxRequestID);
    require(Requests[_requestID].willUpdate == true);

    Requests[_requestID].willUpdate = _willUpdate;
    Requests[_requestID].lastSubscribers = _subscribers;
    Requests[_requestID].lastPostViews = _postViews;
    Requests[_requestID].postTime = _postTime;
    Requests[_requestID].error = _error;

    emit NewResponse(_requestID);
  }

  function CEOWithdraw(address payable _to, uint256 _amount) public {
    require(owner != address(0));
    require(msg.sender == owner);

    _to.transfer(_amount);
  }
}
