import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts:{
        mnemonic: "inherit jump prison shuffle normal pizza cereal broken fantasy pony mechanic sport stage replace wonder recipe faith stumble pigeon dash smoke what exhaust viable",
      }
    },
/*
    hardhat: {
      chainId: 1337,
    }
*/
  },
};

export default config;