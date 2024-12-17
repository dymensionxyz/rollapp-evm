import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts:{
        mnemonic: "penalty useful movie rookie toilet album abuse rude sing size meadow noodle wise pen castle trust proud chalk loud era universe can reflect clarify"
      }
    },
  },
};

export default config;