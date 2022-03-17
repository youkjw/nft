pragma solidity ^0.8.0;

import "./contracts/token/ERC721/ERC721.sol";
import "./contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "./contracts/utils/Counters.sol";

contract Nft is ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    mapping (uint256 => uint256) tokenPrice;

    event Buy(address seller, address buyer, uint256 amount);

    string public version;
    constructor(string memory _version) ERC721("NFT", "MCO") {
        version = _version;
    }

    function buy(uint256 _tokenId) external payable {
        price = tokenPrice[_tokenId];
        require(_msgSender.value >= price, "not enough money");
        require(msg.value >= price, "invalid value");

        address owner = ownerOf(_tokenId);
        payable(owner).transfer(price);

        _safeTransfer(owner, msg.sender, _tokenId, "");
        emit Buy(owner, msg.sender, price);
    }

    function mint(address owner, uint256 _tokenId, string memory _tokenURI) {
        require(ownerOf(_tokenId) != address(0x00), "invalid token_id");

        _safeMint(owner, _tokenId);
        _setTokenURI(_tokenId, _tokenURI);
    }

    function setTokenPrice(uint256 _tokenId, uint256 amount) {
        require(amount > 0, "invalid amount");
        tokenPrice[_tokenId] = amount;
    }
}
