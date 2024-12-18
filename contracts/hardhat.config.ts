import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts:{
        mnemonic: "depend version wrestle document episode celery nuclear main penalty hundred trap scale candy donate search glory build valve round athlete become beauty indicate hamster",
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