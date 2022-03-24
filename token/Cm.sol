pragma solidity ^0.8.0;

import "./contracts/token/ERC20/ERC20.sol";

contract Cm is ERC20 {
    constructor() ERC20("CM", "cm") {
        _mint(msg.sender, 100000000 * 10 ** decimals());
    }
}
