

contract SmartTgStats {

    address                     owner;
    bool                        status;
    uint256                     statusRenew;
    uint256                     balance;

    struct Request {
        bytes32 channel;
        uint32 postID;
        uint256 bid;
        bool willUpdate;
        uint256 started;
        uint256 updated;
        uint32 lastSubscribers;
        uint32 lastPostViews;
        uint256 postTime;
    }

    Request[]                   requests;

    event                       NewRequest(uint);

    constructor(bool _isOn) public payable {
        owner = msg.sender;
        status = _isOn;
        statusRenew = block.timestamp;
    }

   function RenewStatus(bool _isOn) public {
        require(owner != address(0));
        require(owner == msg.sender);
        status = _isOn;
        statusRenew = block.timestamp;
   }

   function AddRequest(bytes32 _channel, uint32 _postID) public payable {
       require(owner != address(0));

       require(status == true);
       require(block.timestamp - statusRenew >= 0);
       require(block.timestamp - statusRenew < 1209600);

       require(msg.value >= 0);
       balance += msg.value;

       requests.push(Request(
            _channel,
            _postID,
            msg.value,
            true,
            block.timestamp, 0,
            0, 0, 0));
       emit NewRequest(requests.length - 1);
   }

  function AddResponse(uint _requestID, bool _willUpdate, uint32 _subscribers, uint32 _postViews, uint256 _postTime) public {
    require(owner != address(0));
    require(owner == msg.sender);
    requests[_requestID].willUpdate = _willUpdate;
    requests[_requestID].lastSubscribers = _subscribers;
    requests[_requestID].lastPostViews = _postViews;
    requests[_requestID].postTime = _postTime;
  }

    function Withdraw(address payable _to, uint256 _amount) public {

        require(owner != address(0));
        require(msg.sender == owner);

        require(balance >= _amount);
        balance -= _amount;

        _to.transfer(_amount);
  }
}
