// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

abstract contract Ownable is Context {
    address private _owner;

    /**
     * @dev The caller account is not authorized to perform an operation.
     */
    error OwnableUnauthorizedAccount(address account);

    /**
     * @dev The owner is not a valid owner account. (eg. `address(0)`)
     */
    error OwnableInvalidOwner(address owner);

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the address provided by the deployer as the initial owner.
     */
    constructor(address initialOwner) {
        if (initialOwner == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }
        _transferOwnership(initialOwner);
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        if (owner() != _msgSender()) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby disabling any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        if (newOwner == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}
contract CommonUtil {
    uint256 public nonce = 0;
    struct Election {
        uint256 id;
        uint256 startDate ;   //timestamp
        uint256 duration;    //timestamp
        uint256 numCandidates;
    }
    struct Ballot {
        bytes32 voterId; // hash cccd_id
        uint256 value;   // id of choice (ballot) in offchain db
    }
    struct Result {
        uint256 id;
        uint256 value;
    }

    function bytes32ToString(bytes32 _b) internal pure returns(string memory) {
        bytes memory res = new bytes(32);

        for (uint256 i = 0; i < 32; ++i) {
            res[i] = _b[i];
        }
        return string(res);
    }
    function _stringToByte(string memory _str) internal pure returns (bytes memory) {
        return bytes(_str);
    }
    function encodeVoterId(string memory id, string memory prvKey) internal pure returns (bytes32) {
        return keccak256(abi.encodeWithSignature(prvKey, id));
    }
}
contract Elector is Ownable, CommonUtil{
    string private _signatureString;

    // electionResult[<id election>][<choice value>] = <uint256> số phiếu
    mapping(uint256 => mapping(uint256 => uint256)) public electionToResult;
    mapping(uint256 => mapping(bytes32 => bool)) public isVoted;
    mapping(uint256 => Election) public elections;

    // electionToBallot[<id election>][ballot id]
    // mapping(uint256 => mapping(uint256 => Ballot)) electionToBallot;

    event NewElectionCreated(uint256 _startDate, uint256 _duration, uint256 _numCandidate);
    event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256[] indexed value);
    constructor(string memory prvKey) Ownable(_msgSender()){
        _signatureString = bytes32ToString(keccak256(abi.encodePacked(prvKey)));
    }

    // apply khí khóa k cho đăng ký candidate. => get candidate Length
    function createNewElection(uint256 _startDate, uint256 _duration, uint256 _numCandidate) external onlyOwner {
        require(_startDate + _duration >= block.timestamp);
        nonce++;
        elections[nonce] = Election(nonce, _startDate, _duration, _numCandidate);
    }

    function getSignature() external view onlyOwner returns (bytes memory) {
        return _stringToByte(_signatureString);
    }

    // value is id in offchain database of candidate Id
    function vote(uint256 electionId, string memory _voterId, uint256[] memory value) external onlyOwner {
        // TODO: require time
        Election memory el = elections[electionId];
        require(block.timestamp >= el.startDate && block.timestamp <= el.startDate + el.duration, "[Vote]: Invalid voting time");
        bytes32 hashVoterId = encodeVoterId(_voterId, _signatureString);
        require(!isVoted[electionId][hashVoterId], "You have voted. Each voter can only vote once!");

        // Ballot memory _bal = Ballot(hashVoterId, value);

        for(uint256 i = 0; i < value.length; ++i){
            electionToResult[electionId][value[i]]++;
        }
        isVoted[electionId][hashVoterId] = true;


        emit Voted(electionId, hashVoterId, value);
    }


    function getElectionResult(uint256 electionId) public view returns (Result[] memory) {
        mapping(uint256 => uint256) storage electionResult = electionToResult[electionId];
        Election memory election = elections[electionId];
        uint256 _numCandidate = election.numCandidates;
        Result[] memory res = new Result[](_numCandidate);
        for (uint64 i = 0; i < _numCandidate; ++i) {
            // i - val
            res[i] = Result(i,electionResult[i]);
        }
        return res;
    }
}
