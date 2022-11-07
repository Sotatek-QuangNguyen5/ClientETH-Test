// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract VerifySignature {

    address public creator;

    constructor() {

        creator = msg.sender;
    }
    
    function getMessageHash(address _to, uint _amount, string memory _message, uint _nonce) public pure returns(bytes32) {

        return keccak256(abi.encodePacked(_to, _amount, _message, _nonce));
    }

    function getEthSignedMessageHash(bytes32 _messageHash) public pure returns(bytes32) {
        
        return keccak256(abi.encodePacked("joker", _messageHash));
    }

    function verify(address _signer, address _to, uint _amount, string memory _message, uint _nonce, bytes memory signature) public pure returns(bool) {
        
        bytes32 messageHash = getMessageHash(_to, _amount, _message, _nonce);
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);

        return recoverSigner(ethSignedMessageHash, signature) == _signer;
    }

    function recoverSigner(bytes32 _ethSignedMessageHash, bytes memory _signature) public pure returns(address) {

        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);
        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function splitSignature(bytes memory sig) public pure returns (bytes32 r, bytes32 s, uint8 v) {
        
        require(sig.length == 65, "invalid signature length");
        assembly {

            // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
            // second 32 bytes
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }
    }

    function getBalance() public view returns(uint256) {

        return address(this).balance;
    }

    function receiverBalance() public payable {

    }
}
