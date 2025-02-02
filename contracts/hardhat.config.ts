
require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: '0.8.20',
  networks: {
    rollapp: {
      url: 'https://json-rpc.ra-2.rollapp.network',
      accounts: {
        mnemonic: 'chalk excess welcome pool sea session pencil health region lamp library today',
      },
    },
  },
};
