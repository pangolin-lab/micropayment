pragma solidity >=0.4.24 <0.7.0;

import "./Token.sol";

contract PayToCheck {
    address owner = msg.sender;
    mapping(uint256 => checkState) usedNonces;
    Token public tokenContract;
    enum checkState {valid,checked}

    constructor(Token _Token) public {
        tokenContract = _Token;
    }

    modifier onlyOwner(){
            require(msg.sender == owner,"only owned is allowed.");
            _;
    }
    modifier nonceValid(uint256 nonce){
           require(usedNonces[nonce]==checkState.valid,"not a valide nonce");
           _;
    }
    function returnCheckState(uint256 nonce) external view returns (checkState){
        return usedNonces[nonce];
    }
    function claimPayment(uint256 amount, uint256 nonce,  uint8 v, bytes32 r, bytes32 s) public nonceValid(nonce){
        usedNonces[nonce] = checkState.checked;
        // this recreates the message that was signed on the client
        bytes32 message = recreateMsg(msg.sender, amount, nonce,address(this));
        require(testRecoveryNoPrefix(message,v,r,s) == owner,"signature check failed");
        tokenContract.transfer(msg.sender,amount);
    }

    function splitSignature(bytes memory sig) public pure returns (uint8 v, bytes32 r, bytes32 s)
    {
        require(sig.length == 65,"invalid sig length");
        assembly {
        // first 32 bytes, after the length prefix.
            r := mload(add(sig, 32))
        // second 32 bytes.
            s := mload(add(sig, 64))
        // final byte (first byte of the next 32 bytes).
            v := byte(0, mload(add(sig, 96)))
        }
        return (v, r, s);
    }

    function claimPaymentAsm(uint256 amount, uint256 nonce,  bytes memory sig) public nonceValid(nonce){
        usedNonces[nonce] = checkState.checked;
        // this recreates the message that was signed on the client
        bytes32 message = recreateMsg(msg.sender, amount, nonce,address(this));
        (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);
        if(v<27){
            v+=27;
        }
        require(testRecoveryNoPrefix(message,v,r,s) == owner,"signature check failed");
        tokenContract.transfer(msg.sender,amount);
    }

    function recreateMsg(address add, uint256 amount, uint256 nonce,address contractAddress) public pure returns (bytes32){
        return keccak256(abi.encodePacked(add, amount, nonce,contractAddress));
    }
    function testHash(address t,uint v) external pure returns (bytes32){
        return  keccak256(abi.encodePacked(t,v));
    }
    function testRecoveryNoPrefix(bytes32 h, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        address addr = ecrecover(h, v, r, s);
        return addr;
    }

    function testOwnerSigned(bytes32 h, uint8 v, bytes32 r, bytes32 s) public view returns (bool){
        address addr = ecrecover(h, v, r, s);
        return addr==owner;
    }

    function testRecovery(bytes32 h, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 t = keccak256(abi.encodePacked(prefix,h));
        address addr = ecrecover(t, v, r, s);
        return addr;
    }
    function kill() public onlyOwner{
        selfdestruct(msg.sender);
    }
}