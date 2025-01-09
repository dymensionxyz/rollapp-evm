#  (2025-01-09)


### Bug Fixes

* **app:** Changed to use sequencerKeeper instead of stakingKeepr in ibcKeeper ([#195](https://github.com/dymensionxyz/rollapp-evm/issues/195)) ([cec299f](https://github.com/dymensionxyz/rollapp-evm/commit/cec299f16e6dba780c0832e6bea53275922cfdd7))
* **app:** fixed bech32 on account keeper to not be hardcoded  ([#165](https://github.com/dymensionxyz/rollapp-evm/issues/165)) ([750d1e7](https://github.com/dymensionxyz/rollapp-evm/commit/750d1e70ad052daf7b2942bcecaf0dddfbc17d90))
* **build:** Added ledger and netgo support in makefile when make install ([#213](https://github.com/dymensionxyz/rollapp-evm/issues/213)) ([1be4bfc](https://github.com/dymensionxyz/rollapp-evm/commit/1be4bfc05dffcdb9139d3d57b1007179ceb4aa20))
* **ci:** Update changelog workflow ([#224](https://github.com/dymensionxyz/rollapp-evm/issues/224)) ([32b141a](https://github.com/dymensionxyz/rollapp-evm/commit/32b141a77fef558b459307ae712eed9ee1f5aacc))
* cleaning file descriptors ([#69](https://github.com/dymensionxyz/rollapp-evm/issues/69)) ([18f7558](https://github.com/dymensionxyz/rollapp-evm/commit/18f7558455b74ee097c525fc418375456b90bb00))
* **code standards:** cleans up blocked addresses in app.go ([#292](https://github.com/dymensionxyz/rollapp-evm/issues/292)) ([d67d5f1](https://github.com/dymensionxyz/rollapp-evm/commit/d67d5f110a99cb3ca2bd55e989fea714cc6ccf34))
* **deps:** bump dymint to `v1.1.3-rc02` to fix unsync issue when da is down ([#227](https://github.com/dymensionxyz/rollapp-evm/issues/227)) ([6c745ac](https://github.com/dymensionxyz/rollapp-evm/commit/6c745ac8ace927f41b163d662d0c60f5797cde96))
* **deps:** bump dymint to v1.1.0-rc04 to fix account error not causing panic ([#210](https://github.com/dymensionxyz/rollapp-evm/issues/210)) ([5636e41](https://github.com/dymensionxyz/rollapp-evm/commit/5636e413b15d2a3caf0afe2873c4e7fc9c963aad))
* **deps:** bump dymint to v1.1.2 to fix full node sync issues ([#222](https://github.com/dymensionxyz/rollapp-evm/issues/222)) ([bbcbf5a](https://github.com/dymensionxyz/rollapp-evm/commit/bbcbf5ac97d18ec70aa80ff61b79532a32494258))
* **deps:** bump dymint to v1.1.3-rc01 to fix full node p2p initial sync issues ([#225](https://github.com/dymensionxyz/rollapp-evm/issues/225)) ([37c7b0f](https://github.com/dymensionxyz/rollapp-evm/commit/37c7b0f907ea97149856f3d344c0a2255ff81c79))
* **deps:** bump evmos to v0.4.2 to fix vesting msgs not blocked in eip712 ante handler ([#218](https://github.com/dymensionxyz/rollapp-evm/issues/218)) ([8176ee3](https://github.com/dymensionxyz/rollapp-evm/commit/8176ee362189a37c54f47023cc32c011165c7283))
* **deps:** bumped dymint to `d51b961e7` to fix stuck submission bug ([#341](https://github.com/dymensionxyz/rollapp-evm/issues/341)) ([395072b](https://github.com/dymensionxyz/rollapp-evm/commit/395072b3552e00d83727e228caff644e9bcc8b0a))
* **deps:** bumped dymint to v1.1.0-rc05 to fix da sync issue ([#211](https://github.com/dymensionxyz/rollapp-evm/issues/211)) ([2ef1d93](https://github.com/dymensionxyz/rollapp-evm/commit/2ef1d936836f70ca956d5996d6038024b18e75bf))
* **deps:** bumped dymint to v1.1.1 to fix da health event fast emmision ([#221](https://github.com/dymensionxyz/rollapp-evm/issues/221)) ([95743b4](https://github.com/dymensionxyz/rollapp-evm/commit/95743b4648b18bd00224563f2144f0e1960b2c5f))
* **deps:** bumped dymint to v1.3.0-rc03 to fix multiple forks sync issue ([#425](https://github.com/dymensionxyz/rollapp-evm/issues/425)) ([bf8ac62](https://github.com/dymensionxyz/rollapp-evm/commit/bf8ac62ba7ed7d5ad04db9493d75bcc907d4be46))
* **deps:** bumped evmos to fix foreign gas denom rpc call ([#410](https://github.com/dymensionxyz/rollapp-evm/issues/410)) ([268dc22](https://github.com/dymensionxyz/rollapp-evm/commit/268dc22cb51220d8c81e79b94d460571a22e8b45))
* **deps:** bumped evmos to fix token-pair registration bug ([#429](https://github.com/dymensionxyz/rollapp-evm/issues/429)) ([0226fb4](https://github.com/dymensionxyz/rollapp-evm/commit/0226fb44e3598428354cb2a81236ad074e947710))
* **deps:** bumped rdk to 3fe31b2db to fix denom-metadata transfer on unrelayed packets ([#342](https://github.com/dymensionxyz/rollapp-evm/issues/342)) ([e25a992](https://github.com/dymensionxyz/rollapp-evm/commit/e25a99275f584745a30271fb8f73209cc36d2219))
* **deps:** bumped rdk to v1.8.0-rc02 to include 2 minor fixes to tokenfactory and timeupgrade ([#426](https://github.com/dymensionxyz/rollapp-evm/issues/426)) ([03bea59](https://github.com/dymensionxyz/rollapp-evm/commit/03bea593336cd1f2a4186c87c72557398ee7e36e))
* **deps:** bumps `block-explorer-rpc-cosmos v1.0.3` & `evm-block-explorer-rpc-cosmos` v1.0.3 ([#142](https://github.com/dymensionxyz/rollapp-evm/issues/142)) ([ea5e5fd](https://github.com/dymensionxyz/rollapp-evm/commit/ea5e5fdc854d5a4fa4079c4d79b79732e78cf9d8))
* **deps:** double events on evm transfer ([#367](https://github.com/dymensionxyz/rollapp-evm/issues/367)) ([a11c4d0](https://github.com/dymensionxyz/rollapp-evm/commit/a11c4d092fde6f1223ae07ea5bbc93876ed50319))
* **deps:** dymint bump to include da grpc fix ([#346](https://github.com/dymensionxyz/rollapp-evm/issues/346)) ([00267a3](https://github.com/dymensionxyz/rollapp-evm/commit/00267a363e33dd987a3bcc8867621b3be1c3dd6c))
* **deps:** update evmos tag v0.4.1 to block vesting account ([#203](https://github.com/dymensionxyz/rollapp-evm/issues/203)) ([56eb2c1](https://github.com/dymensionxyz/rollapp-evm/commit/56eb2c1326ebcf7ea3a45a017f252a5868351d24))
* **deps:** update rdk to fix tokenfactory denom-metadata override ([#398](https://github.com/dymensionxyz/rollapp-evm/issues/398)) ([4d6aa3e](https://github.com/dymensionxyz/rollapp-evm/commit/4d6aa3ef3ed5c1b6288e245f186c0348db09f2ed))
* **deps:** updated dymint and rdk by removing the replace which override it ([#344](https://github.com/dymensionxyz/rollapp-evm/issues/344)) ([fe4246e](https://github.com/dymensionxyz/rollapp-evm/commit/fe4246e7ca7f4a636881eb099ebd6e10cd386133))
* **deps:** updated dymint to v1.1.0-rc03 to solve gossip bug ([#205](https://github.com/dymensionxyz/rollapp-evm/issues/205)) ([2bd1daa](https://github.com/dymensionxyz/rollapp-evm/commit/2bd1daa701e0568589b2ff8d235805e3812df59f))
* **deps:** updated protobuf to v1.33.0 due to vulnerability ([#204](https://github.com/dymensionxyz/rollapp-evm/issues/204)) ([ea4d3f1](https://github.com/dymensionxyz/rollapp-evm/commit/ea4d3f14d4828d54f7dff543caddd704e6339ae8))
* **dockerfile:** download wasmvm ([#293](https://github.com/dymensionxyz/rollapp-evm/issues/293)) ([5f2ffeb](https://github.com/dymensionxyz/rollapp-evm/commit/5f2ffebbf7ee015ab7499ea2001ca33db0711c3d))
* **drs:** drs-2 migration default gas denom fix ([#404](https://github.com/dymensionxyz/rollapp-evm/issues/404)) ([23ecd7a](https://github.com/dymensionxyz/rollapp-evm/commit/23ecd7a07f42132656b31af6708034dd209a752c))
* failing tests ([#358](https://github.com/dymensionxyz/rollapp-evm/issues/358)) ([db081bf](https://github.com/dymensionxyz/rollapp-evm/commit/db081bf62fc54c0b76aba3da9be19d72a6729d1f))
* feegrant ante handler wiring ([#378](https://github.com/dymensionxyz/rollapp-evm/issues/378)) ([5d5f322](https://github.com/dymensionxyz/rollapp-evm/commit/5d5f32238fac46bf3b6bf28a2457fb5cf66746b1))
* **genesis transfers:** wires the middleware correctly ([#290](https://github.com/dymensionxyz/rollapp-evm/issues/290)) ([941ac19](https://github.com/dymensionxyz/rollapp-evm/commit/941ac193cfe633501715634421b8d08d727dbe92))
* **genesis-templates:** fixed genesis template generator and drs 1,2 templates ([#402](https://github.com/dymensionxyz/rollapp-evm/issues/402)) ([0be4b39](https://github.com/dymensionxyz/rollapp-evm/commit/0be4b39dac30cd500c6c2c823b2da621c73303d9))
* **github:** use go 1.22.4 or greater in action ([#291](https://github.com/dymensionxyz/rollapp-evm/issues/291)) ([fc01049](https://github.com/dymensionxyz/rollapp-evm/commit/fc01049496a252090a2515e98cdb8cd34e58047a))
* hotfix for ibc-go due to val-set hotfix done on hub for froopyland. ([#73](https://github.com/dymensionxyz/rollapp-evm/issues/73)) ([40c6ec4](https://github.com/dymensionxyz/rollapp-evm/commit/40c6ec4c3f899268bd93cf95dee98aa00410fe18))
* **ibc decorator:** fixed bad comparison for ibc messages. ([#339](https://github.com/dymensionxyz/rollapp-evm/issues/339)) ([91c45a5](https://github.com/dymensionxyz/rollapp-evm/commit/91c45a5cb7cb1605613ac830bb77850a3d13e571))
* **local script:** updated with latest relayer version and extra eips ([#199](https://github.com/dymensionxyz/rollapp-evm/issues/199)) ([b032ae1](https://github.com/dymensionxyz/rollapp-evm/commit/b032ae1d19af005bc5a5f94c7e9e67ea81701a26))
* **makefile:** outdated template and fix for makefile ([#386](https://github.com/dymensionxyz/rollapp-evm/issues/386)) ([521d665](https://github.com/dymensionxyz/rollapp-evm/commit/521d6657f778751cbc3729ecc8599556fdb1d2b7))
* multiple fixes to advance readme features ([#141](https://github.com/dymensionxyz/rollapp-evm/issues/141)) ([469d39f](https://github.com/dymensionxyz/rollapp-evm/commit/469d39fc79591cdae4455839db1546cc5bd9c053))
* **readme:** broken links have been renewed. ([#78](https://github.com/dymensionxyz/rollapp-evm/issues/78)) ([c7df6f2](https://github.com/dymensionxyz/rollapp-evm/commit/c7df6f29c8b9981d7a998be4091d2e96c19647a3))
* **readme:** fix governor creation and block time substition missing in readme ([#206](https://github.com/dymensionxyz/rollapp-evm/issues/206)) ([f612207](https://github.com/dymensionxyz/rollapp-evm/commit/f612207fbc798da1772ea0444f1fa3658a010bf6))
* **readme:** fixed init script with broken consensus params ([#240](https://github.com/dymensionxyz/rollapp-evm/issues/240)) ([302762d](https://github.com/dymensionxyz/rollapp-evm/commit/302762d175c934d906b47537873821f9d92d5979))
* reverted bad deps ([#38](https://github.com/dymensionxyz/rollapp-evm/issues/38)) ([545a367](https://github.com/dymensionxyz/rollapp-evm/commit/545a367643d7b1f6e2dbdce1cdb436fdb56000fe))
* **rollappparams:** Validate gas price gov proposal param change ([#421](https://github.com/dymensionxyz/rollapp-evm/issues/421)) ([af153dc](https://github.com/dymensionxyz/rollapp-evm/commit/af153dc6738f95cc3d9636e0ed4316ce545de0bf))
* **scripts:** add an option to skip evm base fees ([#162](https://github.com/dymensionxyz/rollapp-evm/issues/162)) ([ea51eee](https://github.com/dymensionxyz/rollapp-evm/commit/ea51eee8d66dbba587d6ec00395418a1f08b99a8))
* **scripts:** adjusted IBC script to work with whitelisted relayers ([#365](https://github.com/dymensionxyz/rollapp-evm/issues/365)) ([53027b8](https://github.com/dymensionxyz/rollapp-evm/commit/53027b8e1497954837d13b950b6538269dbd8084))
* **scripts:** fix hubgenesis tokens in update genesis ([#172](https://github.com/dymensionxyz/rollapp-evm/issues/172)) ([8d37db8](https://github.com/dymensionxyz/rollapp-evm/commit/8d37db874902eb483293894333416d54bd051e72))
* **scripts:** init script fix da rollapp param for celestia mocha ([#333](https://github.com/dymensionxyz/rollapp-evm/issues/333)) ([8310803](https://github.com/dymensionxyz/rollapp-evm/commit/8310803d37eb86b8db81ef0ff689dad87404f95b))
* **scripts:** remove redundant line in setup ibc script ([#161](https://github.com/dymensionxyz/rollapp-evm/issues/161)) ([57d4f17](https://github.com/dymensionxyz/rollapp-evm/commit/57d4f170779dbaeac0e877d1005709d48a4df2f0))
* **scripts:** support for latest Hub with light client changes, and new relayer version ([#323](https://github.com/dymensionxyz/rollapp-evm/issues/323)) ([725fd82](https://github.com/dymensionxyz/rollapp-evm/commit/725fd82c54f61ae8ea7824af98856d4a02cc2700))
* **scripts:** update init script to support mock and real da and settlement ([#299](https://github.com/dymensionxyz/rollapp-evm/issues/299)) ([de18992](https://github.com/dymensionxyz/rollapp-evm/commit/de189924583db789469933dfa015fbc2a017d2b3))
* **scripts:** use settlement executable ([#320](https://github.com/dymensionxyz/rollapp-evm/issues/320)) ([7b10b89](https://github.com/dymensionxyz/rollapp-evm/commit/7b10b8935cd15fb84e3d6f394f245c1038d91685))
* **sequencers:** get by cons addr in all scenarios on EVM ([#319](https://github.com/dymensionxyz/rollapp-evm/issues/319)) ([0c9f532](https://github.com/dymensionxyz/rollapp-evm/commit/0c9f53268e4f76c75e2ebb46fdbf3c343f90d1d1))
* **test scripts:** make alice and bob keys in .rollap dir ([#147](https://github.com/dymensionxyz/rollapp-evm/issues/147)) ([b6ee646](https://github.com/dymensionxyz/rollapp-evm/commit/b6ee64640af1b528728414bded6a70216a4b5fdf))
* update local setup scripts ([#100](https://github.com/dymensionxyz/rollapp-evm/issues/100)) ([8e8accc](https://github.com/dymensionxyz/rollapp-evm/commit/8e8accc602449481a3188dfe23c147e4ade48877))
* Update README.md rollapp_id env var ([1cf5795](https://github.com/dymensionxyz/rollapp-evm/commit/1cf57952a39db11623a437b5d9c518da3635aff2))
* updated block size and evm `no_base_fee` ([#160](https://github.com/dymensionxyz/rollapp-evm/issues/160)) ([876ccad](https://github.com/dymensionxyz/rollapp-evm/commit/876ccad96765d0d3bd279903c552ab483ecf6b9a))
* updated init.sh to set gensis operator address ([#123](https://github.com/dymensionxyz/rollapp-evm/issues/123)) ([7cb34f8](https://github.com/dymensionxyz/rollapp-evm/commit/7cb34f84740a244fb168a2c2864303cbdcffe827))
* Updated libp2p to use our fork  ([#90](https://github.com/dymensionxyz/rollapp-evm/issues/90)) ([1bc2110](https://github.com/dymensionxyz/rollapp-evm/commit/1bc21109927dcd649e3c04cd624da36f5b3c232c))
* **upgrade:** drs upgrade from 1 to 3 fix ([#415](https://github.com/dymensionxyz/rollapp-evm/issues/415)) ([44e766b](https://github.com/dymensionxyz/rollapp-evm/commit/44e766b4782904533bb557cd3b4415e544d69ea3))
* **version:** bump dymint and rdk to latest ([6c23d9d](https://github.com/dymensionxyz/rollapp-evm/commit/6c23d9defc4914a0ae1aafaff972b4aa7d18a6b8))
* **version:** bump dymint and rdk to latest ([#366](https://github.com/dymensionxyz/rollapp-evm/issues/366)) ([105a1a0](https://github.com/dymensionxyz/rollapp-evm/commit/105a1a0bad4677993f76f999e323dde20d147c57))
* **version:** bump dymint to c0e39f93d729 ([#375](https://github.com/dymensionxyz/rollapp-evm/issues/375)) ([97c1205](https://github.com/dymensionxyz/rollapp-evm/commit/97c120597d47468b1332347998b4e7bb5966724b))


### Features

* add DRS3 upgrade handler ([#409](https://github.com/dymensionxyz/rollapp-evm/issues/409)) ([6042234](https://github.com/dymensionxyz/rollapp-evm/commit/60422347276e50609b7f3178396f6c1212ae230e))
* add GenesisChecksum to genesis info ([#361](https://github.com/dymensionxyz/rollapp-evm/issues/361)) ([8ff81e3](https://github.com/dymensionxyz/rollapp-evm/commit/8ff81e39a1650801e0dfd50fcf5a8dc0c50b50e0))
* add swagger config and make scripts ([#130](https://github.com/dymensionxyz/rollapp-evm/issues/130)) ([41718e4](https://github.com/dymensionxyz/rollapp-evm/commit/41718e4d4098e6bf18117c31b514b2cc226a331f))
* allow grpc CELESTIA_NETWORK when running setup script ([#441](https://github.com/dymensionxyz/rollapp-evm/issues/441)) ([b69f307](https://github.com/dymensionxyz/rollapp-evm/commit/b69f307e54f73b899aaa816f186952ecabd3456a))
* **ante:** allow doing vesting txs based on whitelist ([#216](https://github.com/dymensionxyz/rollapp-evm/issues/216)) ([e1c968f](https://github.com/dymensionxyz/rollapp-evm/commit/e1c968f3dc53d0d55755858467fff3aa8b8a669f))
* **ante:** Skip fees for IBC messages for existing and new accounts ([#336](https://github.com/dymensionxyz/rollapp-evm/issues/336)) ([1dbc8cc](https://github.com/dymensionxyz/rollapp-evm/commit/1dbc8ccab79645b7c7f02f316a4538db8408495b))
* **ante:** whitelisted relayers ([#357](https://github.com/dymensionxyz/rollapp-evm/issues/357)) ([f850e3b](https://github.com/dymensionxyz/rollapp-evm/commit/f850e3bc84ee450745a81e629ef6017f852a9b23))
* **app:** add v2.2.0 upgrade handler ([#248](https://github.com/dymensionxyz/rollapp-evm/issues/248)) ([c9ffc7f](https://github.com/dymensionxyz/rollapp-evm/commit/c9ffc7f18d8bf6a67f341dcfeb83b8a0ef64f401))
* **app:** removed the wiring of claims module ([#377](https://github.com/dymensionxyz/rollapp-evm/issues/377)) ([99dc8a6](https://github.com/dymensionxyz/rollapp-evm/commit/99dc8a658d6ad422ac016bcf6a45b6fbf03169d2))
* **app:** return genesis bridge data in InitChainResponse ([#370](https://github.com/dymensionxyz/rollapp-evm/issues/370)) ([9bc4283](https://github.com/dymensionxyz/rollapp-evm/commit/9bc4283d3eb97fcbad5188fa3406d61cd7dba061))
* **app:** wire authz and feegrant modules to app ([#270](https://github.com/dymensionxyz/rollapp-evm/issues/270)) ([dad400c](https://github.com/dymensionxyz/rollapp-evm/commit/dad400ce41e210f229ae112531446b37a2d6d981))
* **be:** integrate block explorer Json-RPC server ([#132](https://github.com/dymensionxyz/rollapp-evm/issues/132)) ([d73b1c4](https://github.com/dymensionxyz/rollapp-evm/commit/d73b1c451b93f04a1db5a73c4c8c78fc21729208))
* **build:** add ability to override bech32 with env var  ([#405](https://github.com/dymensionxyz/rollapp-evm/issues/405)) ([70dddc2](https://github.com/dymensionxyz/rollapp-evm/commit/70dddc2adf670dac8c9c315a7afef4d6645b9105))
* consensus messages ([#351](https://github.com/dymensionxyz/rollapp-evm/issues/351)) ([147382a](https://github.com/dymensionxyz/rollapp-evm/commit/147382acf3359a026996df1e8efd7e32a03a3542))
* **denommetadata:** wire the denommetadata ibc middleware ([#283](https://github.com/dymensionxyz/rollapp-evm/issues/283)) ([faa3771](https://github.com/dymensionxyz/rollapp-evm/commit/faa377155842691976eeb72285dd5afa08946aff))
* **deps:** bumped rdk to support tokenless feature ([#411](https://github.com/dymensionxyz/rollapp-evm/issues/411)) ([254cc7f](https://github.com/dymensionxyz/rollapp-evm/commit/254cc7fa8b3d2c15f85d612cd5b434bfbf82c9d4))
* **deps:** command to validate genesis/bridge ([#380](https://github.com/dymensionxyz/rollapp-evm/issues/380)) ([6b900bc](https://github.com/dymensionxyz/rollapp-evm/commit/6b900bc70ffe8cf76efa69c639ee959c33879596))
* **deps:** new genesis-bridge flow and rollappparams upgrade fix ([#394](https://github.com/dymensionxyz/rollapp-evm/issues/394)) ([fb948a9](https://github.com/dymensionxyz/rollapp-evm/commit/fb948a98feb0ff20aec8af9fc363046ba2d8792d))
* **deps:** updated dymint to v1.3.0-rc02 to support skip validation height flag ([#424](https://github.com/dymensionxyz/rollapp-evm/issues/424)) ([74e73b3](https://github.com/dymensionxyz/rollapp-evm/commit/74e73b35511c0c113c2a248e45a9248a5327ed8a))
* **deps:** updated evmos fork to introduce support for custom gas denom ([#399](https://github.com/dymensionxyz/rollapp-evm/issues/399)) ([6555be4](https://github.com/dymensionxyz/rollapp-evm/commit/6555be48c7e413acb66f61f92a0c8af4630c8df2))
* **drs:** Added drs 2 templates ([#401](https://github.com/dymensionxyz/rollapp-evm/issues/401)) ([d0bf365](https://github.com/dymensionxyz/rollapp-evm/commit/d0bf365d97be3fd67bf32792ab043b6460f73749))
* **drs:** Added genesis templates for drs 5 ([#432](https://github.com/dymensionxyz/rollapp-evm/issues/432)) ([a8e28f0](https://github.com/dymensionxyz/rollapp-evm/commit/a8e28f05622b918690d81b40cb4b1ff68d098d59))
* **drs:** changed DRS to be int vs commit hash ([#359](https://github.com/dymensionxyz/rollapp-evm/issues/359)) ([1713976](https://github.com/dymensionxyz/rollapp-evm/commit/17139765cedbb81fbf41dd7530fc8c85f136a446))
* **drs:** updated makefile and templates to DRS 4 ([#423](https://github.com/dymensionxyz/rollapp-evm/issues/423)) ([09d308d](https://github.com/dymensionxyz/rollapp-evm/commit/09d308df2e7b6734dd079ae06dd2052cfa95b3aa))
* **drs:** upgrade handler for DRS 5 ([#431](https://github.com/dymensionxyz/rollapp-evm/issues/431)) ([a656462](https://github.com/dymensionxyz/rollapp-evm/commit/a6564629bd2e1d1be6d644f134b2b11e0f38f9f2))
* **genesis bridge:** genesis transfers ([#279](https://github.com/dymensionxyz/rollapp-evm/issues/279)) ([f0e0909](https://github.com/dymensionxyz/rollapp-evm/commit/f0e0909542891b6ed1444c4d3acae22d824de5de))
* **genesis_bridge:** revised genesis bridge ([#353](https://github.com/dymensionxyz/rollapp-evm/issues/353)) ([0eb27ff](https://github.com/dymensionxyz/rollapp-evm/commit/0eb27ff77c9e09b596a548e24f3562adf291e430))
* **genesis-template:** changed default delegator unbonding time in genesis template ([#437](https://github.com/dymensionxyz/rollapp-evm/issues/437)) ([587ab0c](https://github.com/dymensionxyz/rollapp-evm/commit/587ab0c0875aa16413036913f7ed8fc5f3557465))
* **genesis-template:** changed delegator default min commision rate to be 0 ([#438](https://github.com/dymensionxyz/rollapp-evm/issues/438)) ([1fcdfb3](https://github.com/dymensionxyz/rollapp-evm/commit/1fcdfb32a6de5e57176c7aea39f3cfb34fee7da2))
* **genesis-templates:** add genesis templates for drs 3 ([#412](https://github.com/dymensionxyz/rollapp-evm/issues/412)) ([21acc35](https://github.com/dymensionxyz/rollapp-evm/commit/21acc354cd310b31ba4827883d9e26b39c143fa8))
* **makefile:** create genesis template with DRS from make file ([#385](https://github.com/dymensionxyz/rollapp-evm/issues/385)) ([7877e6d](https://github.com/dymensionxyz/rollapp-evm/commit/7877e6db54341ef55278c61acae216e8f5c9916a))
* migration for mainnet rollapps (nim + mande) to upgrade to 3D ([#442](https://github.com/dymensionxyz/rollapp-evm/issues/442)) ([84aa604](https://github.com/dymensionxyz/rollapp-evm/commit/84aa604882b207761cc7d924b0e1cecbd4234fe4))
* **script:** added replacement for gas_denom ([#403](https://github.com/dymensionxyz/rollapp-evm/issues/403)) ([da6e561](https://github.com/dymensionxyz/rollapp-evm/commit/da6e5617220647519277115cf11c7f960fccd1ab))
* **sequencers:** wire sequencer rewards functionality ([#316](https://github.com/dymensionxyz/rollapp-evm/issues/316)) ([646dfab](https://github.com/dymensionxyz/rollapp-evm/commit/646dfab8601d91941319fe876459fd59047b9877))
* set bech32 prefix without changing source code ([#207](https://github.com/dymensionxyz/rollapp-evm/issues/207)) ([d750eda](https://github.com/dymensionxyz/rollapp-evm/commit/d750eda7ffb96028a70d6159e0c4573e1c2bf10d))
* **upgrade:** add drs 2 upgrade handler ([#396](https://github.com/dymensionxyz/rollapp-evm/issues/396)) ([950d2ba](https://github.com/dymensionxyz/rollapp-evm/commit/950d2ba2b1eb6c616e9ab99ebbc6a49b65326fdd))
* **upgrade:** add upgrade handler for DRS4 ([#420](https://github.com/dymensionxyz/rollapp-evm/issues/420)) ([92a3e85](https://github.com/dymensionxyz/rollapp-evm/commit/92a3e85fc61fe64c4c414f2e1caf2be74e3bf09b))
* **upgrade:** move drs check to begin blocker to allow gov software upgrades ([#393](https://github.com/dymensionxyz/rollapp-evm/issues/393)) ([f9bd1b6](https://github.com/dymensionxyz/rollapp-evm/commit/f9bd1b65d8cd24713b50e580e698ba129e77fc40))
* wire `x/timeupgrade` module ([#331](https://github.com/dymensionxyz/rollapp-evm/issues/331)) ([7778026](https://github.com/dymensionxyz/rollapp-evm/commit/7778026fbfa5bd48f02a24c8f410e24ae68aa351))
* wire rollapp params module ([#307](https://github.com/dymensionxyz/rollapp-evm/issues/307)) ([09bb5c2](https://github.com/dymensionxyz/rollapp-evm/commit/09bb5c2dbea33bcb13ebc3049c7d12ba5e522721))
* wired signless execution ([#416](https://github.com/dymensionxyz/rollapp-evm/issues/416)) ([63216d5](https://github.com/dymensionxyz/rollapp-evm/commit/63216d5e64b6cc89bcc4c4e5c714e75cd7938b3e))



# [](https://github.com/dymensionxyz/rollapp-evm/compare/v2.2.0-rc01...v) (2024-05-20)


### Bug Fixes

* **ci:** Update changelog workflow ([#224](https://github.com/dymensionxyz/rollapp-evm/issues/224)) ([32b141a](https://github.com/dymensionxyz/rollapp-evm/commit/32b141a77fef558b459307ae712eed9ee1f5aacc))
* **readme:** fixed init script with broken consensus params ([#240](https://github.com/dymensionxyz/rollapp-evm/issues/240)) ([302762d](https://github.com/dymensionxyz/rollapp-evm/commit/302762d175c934d906b47537873821f9d92d5979))



# [2.2.0-rc01](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.3-rc02...v2.2.0-rc01) (2024-05-09)


### Features

* **ante:** allow doing vesting txs based on whitelist ([#216](https://github.com/dymensionxyz/rollapp-evm/issues/216)) ([e1c968f](https://github.com/dymensionxyz/rollapp-evm/commit/e1c968f3dc53d0d55755858467fff3aa8b8a669f))



## [2.1.3-rc02](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.3...v2.1.3-rc02) (2024-05-03)


### Bug Fixes

* **deps:** bump dymint to `v1.1.3-rc02` to fix unsync issue when da is down ([#227](https://github.com/dymensionxyz/rollapp-evm/issues/227)) ([6c745ac](https://github.com/dymensionxyz/rollapp-evm/commit/6c745ac8ace927f41b163d662d0c60f5797cde96))



## [2.1.3-rc01](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.2...v2.1.3-rc01) (2024-05-02)


### Bug Fixes

* **deps:** bump dymint to v1.1.3-rc01 to fix full node p2p initial sync issues ([#225](https://github.com/dymensionxyz/rollapp-evm/issues/225)) ([37c7b0f](https://github.com/dymensionxyz/rollapp-evm/commit/37c7b0f907ea97149856f3d344c0a2255ff81c79))



## [2.1.2](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.1...v2.1.2) (2024-05-01)


### Bug Fixes

* **deps:** bump dymint to v1.1.2 to fix full node sync issues ([#222](https://github.com/dymensionxyz/rollapp-evm/issues/222)) ([bbcbf5a](https://github.com/dymensionxyz/rollapp-evm/commit/bbcbf5ac97d18ec70aa80ff61b79532a32494258))



## [2.1.1](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0...v2.1.1) (2024-04-30)


### Bug Fixes

* **deps:** bumped dymint to v1.1.1 to fix da health event fast emmision ([#221](https://github.com/dymensionxyz/rollapp-evm/issues/221)) ([95743b4](https://github.com/dymensionxyz/rollapp-evm/commit/95743b4648b18bd00224563f2144f0e1960b2c5f))



# [2.1.0](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc09...v2.1.0) (2024-04-29)



# [2.1.0-rc09](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc08...v2.1.0-rc09) (2024-04-29)


### Bug Fixes

* **build:** Added ledger and netgo support in makefile when make install ([#213](https://github.com/dymensionxyz/rollapp-evm/issues/213)) ([1be4bfc](https://github.com/dymensionxyz/rollapp-evm/commit/1be4bfc05dffcdb9139d3d57b1007179ceb4aa20))
* **deps:** bump evmos to v0.4.2 to fix vesting msgs not blocked in eip712 ante handler ([#218](https://github.com/dymensionxyz/rollapp-evm/issues/218)) ([8176ee3](https://github.com/dymensionxyz/rollapp-evm/commit/8176ee362189a37c54f47023cc32c011165c7283))



# [2.1.0-rc08](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc07...v2.1.0-rc08) (2024-04-28)


### Bug Fixes

* **deps:** bumped dymint to v1.1.0-rc05 to fix da sync issue ([#211](https://github.com/dymensionxyz/rollapp-evm/issues/211)) ([2ef1d93](https://github.com/dymensionxyz/rollapp-evm/commit/2ef1d936836f70ca956d5996d6038024b18e75bf))



# [2.1.0-rc07](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc06...v2.1.0-rc07) (2024-04-28)


### Bug Fixes

* **deps:** bump dymint to v1.1.0-rc04 to fix account error not causing panic ([#210](https://github.com/dymensionxyz/rollapp-evm/issues/210)) ([5636e41](https://github.com/dymensionxyz/rollapp-evm/commit/5636e413b15d2a3caf0afe2873c4e7fc9c963aad))



# [2.1.0-rc06](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc05...v2.1.0-rc06) (2024-04-28)


### Features

* set bech32 prefix without changing source code ([#207](https://github.com/dymensionxyz/rollapp-evm/issues/207)) ([d750eda](https://github.com/dymensionxyz/rollapp-evm/commit/d750eda7ffb96028a70d6159e0c4573e1c2bf10d))



# [2.1.0-rc05](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc04...v2.1.0-rc05) (2024-04-27)


### Bug Fixes

* **readme:** fix governor creation and block time substition missing in readme ([#206](https://github.com/dymensionxyz/rollapp-evm/issues/206)) ([f612207](https://github.com/dymensionxyz/rollapp-evm/commit/f612207fbc798da1772ea0444f1fa3658a010bf6))



# [2.1.0-rc04](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc03...v2.1.0-rc04) (2024-04-27)


### Bug Fixes

* **deps:** update evmos tag v0.4.1 to block vesting account ([#203](https://github.com/dymensionxyz/rollapp-evm/issues/203)) ([56eb2c1](https://github.com/dymensionxyz/rollapp-evm/commit/56eb2c1326ebcf7ea3a45a017f252a5868351d24))
* **deps:** updated dymint to v1.1.0-rc03 to solve gossip bug ([#205](https://github.com/dymensionxyz/rollapp-evm/issues/205)) ([2bd1daa](https://github.com/dymensionxyz/rollapp-evm/commit/2bd1daa701e0568589b2ff8d235805e3812df59f))
* **deps:** updated protobuf to v1.33.0 due to vulnerability ([#204](https://github.com/dymensionxyz/rollapp-evm/issues/204)) ([ea4d3f1](https://github.com/dymensionxyz/rollapp-evm/commit/ea4d3f14d4828d54f7dff543caddd704e6339ae8))



# [2.1.0-rc03](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc02...v2.1.0-rc03) (2024-04-26)


### Bug Fixes

* **local script:** updated with latest relayer version and extra eips ([#199](https://github.com/dymensionxyz/rollapp-evm/issues/199)) ([b032ae1](https://github.com/dymensionxyz/rollapp-evm/commit/b032ae1d19af005bc5a5f94c7e9e67ea81701a26))



# [2.1.0-rc02](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-rc01...v2.1.0-rc02) (2024-04-26)



# [2.1.0-rc01](https://github.com/dymensionxyz/rollapp-evm/compare/v2.2.0-alpha...v2.1.0-rc01) (2024-04-26)


### Bug Fixes

* **app:** Changed to use sequencerKeeper instead of stakingKeepr in ibcKeeper ([#195](https://github.com/dymensionxyz/rollapp-evm/issues/195)) ([cec299f](https://github.com/dymensionxyz/rollapp-evm/commit/cec299f16e6dba780c0832e6bea53275922cfdd7))
* **app:** fixed bech32 on account keeper to not be hardcoded  ([#165](https://github.com/dymensionxyz/rollapp-evm/issues/165)) ([750d1e7](https://github.com/dymensionxyz/rollapp-evm/commit/750d1e70ad052daf7b2942bcecaf0dddfbc17d90))
* **deps:** bumps `block-explorer-rpc-cosmos v1.0.3` & `evm-block-explorer-rpc-cosmos` v1.0.3 ([#142](https://github.com/dymensionxyz/rollapp-evm/issues/142)) ([ea5e5fd](https://github.com/dymensionxyz/rollapp-evm/commit/ea5e5fdc854d5a4fa4079c4d79b79732e78cf9d8))
* multiple fixes to advance readme features ([#141](https://github.com/dymensionxyz/rollapp-evm/issues/141)) ([469d39f](https://github.com/dymensionxyz/rollapp-evm/commit/469d39fc79591cdae4455839db1546cc5bd9c053))
* **readme:** broken links have been renewed. ([#78](https://github.com/dymensionxyz/rollapp-evm/issues/78)) ([c7df6f2](https://github.com/dymensionxyz/rollapp-evm/commit/c7df6f29c8b9981d7a998be4091d2e96c19647a3))
* **scripts:** add an option to skip evm base fees ([#162](https://github.com/dymensionxyz/rollapp-evm/issues/162)) ([ea51eee](https://github.com/dymensionxyz/rollapp-evm/commit/ea51eee8d66dbba587d6ec00395418a1f08b99a8))
* **scripts:** fix hubgenesis tokens in update genesis ([#172](https://github.com/dymensionxyz/rollapp-evm/issues/172)) ([8d37db8](https://github.com/dymensionxyz/rollapp-evm/commit/8d37db874902eb483293894333416d54bd051e72))
* **scripts:** remove redundant line in setup ibc script ([#161](https://github.com/dymensionxyz/rollapp-evm/issues/161)) ([57d4f17](https://github.com/dymensionxyz/rollapp-evm/commit/57d4f170779dbaeac0e877d1005709d48a4df2f0))
* **test scripts:** make alice and bob keys in .rollap dir ([#147](https://github.com/dymensionxyz/rollapp-evm/issues/147)) ([b6ee646](https://github.com/dymensionxyz/rollapp-evm/commit/b6ee64640af1b528728414bded6a70216a4b5fdf))
* updated block size and evm `no_base_fee` ([#160](https://github.com/dymensionxyz/rollapp-evm/issues/160)) ([876ccad](https://github.com/dymensionxyz/rollapp-evm/commit/876ccad96765d0d3bd279903c552ab483ecf6b9a))


### Features

* add swagger config and make scripts ([#130](https://github.com/dymensionxyz/rollapp-evm/issues/130)) ([41718e4](https://github.com/dymensionxyz/rollapp-evm/commit/41718e4d4098e6bf18117c31b514b2cc226a331f))
* **be:** integrate block explorer Json-RPC server ([#132](https://github.com/dymensionxyz/rollapp-evm/issues/132)) ([d73b1c4](https://github.com/dymensionxyz/rollapp-evm/commit/d73b1c451b93f04a1db5a73c4c8c78fc21729208))



# [2.2.0-alpha](https://github.com/dymensionxyz/rollapp-evm/compare/v2.1.0-alpha...v2.2.0-alpha) (2024-03-26)


### Bug Fixes

* updated init.sh to set gensis operator address ([#123](https://github.com/dymensionxyz/rollapp-evm/issues/123)) ([7cb34f8](https://github.com/dymensionxyz/rollapp-evm/commit/7cb34f84740a244fb168a2c2864303cbdcffe827))



# [2.1.0-alpha](https://github.com/dymensionxyz/rollapp-evm/compare/v2.0.0-beta...v2.1.0-alpha) (2024-03-22)


### Bug Fixes

* update local setup scripts ([#100](https://github.com/dymensionxyz/rollapp-evm/issues/100)) ([8e8accc](https://github.com/dymensionxyz/rollapp-evm/commit/8e8accc602449481a3188dfe23c147e4ade48877))
* Update README.md rollapp_id env var ([1cf5795](https://github.com/dymensionxyz/rollapp-evm/commit/1cf57952a39db11623a437b5d9c518da3635aff2))
* Updated libp2p to use our fork  ([#90](https://github.com/dymensionxyz/rollapp-evm/issues/90)) ([1bc2110](https://github.com/dymensionxyz/rollapp-evm/commit/1bc21109927dcd649e3c04cd624da36f5b3c232c))



# [2.0.0-beta](https://github.com/dymensionxyz/rollapp-evm/compare/v1.0.0-beta...v2.0.0-beta) (2024-01-15)


### Bug Fixes

* cleaning file descriptors ([#69](https://github.com/dymensionxyz/rollapp-evm/issues/69)) ([18f7558](https://github.com/dymensionxyz/rollapp-evm/commit/18f7558455b74ee097c525fc418375456b90bb00))
* hotfix for ibc-go due to val-set hotfix done on hub for froopyland. ([#73](https://github.com/dymensionxyz/rollapp-evm/issues/73)) ([40c6ec4](https://github.com/dymensionxyz/rollapp-evm/commit/40c6ec4c3f899268bd93cf95dee98aa00410fe18))



# [1.0.0-beta](https://github.com/dymensionxyz/rollapp-evm/compare/v0.1.0-rc3...v1.0.0-beta) (2023-10-19)


### Bug Fixes

* reverted bad deps ([#38](https://github.com/dymensionxyz/rollapp-evm/issues/38)) ([545a367](https://github.com/dymensionxyz/rollapp-evm/commit/545a367643d7b1f6e2dbdce1cdb436fdb56000fe))



# [0.1.0-rc2](https://github.com/dymensionxyz/rollapp-evm/compare/v0.1.0-rc1...v0.1.0-rc2) (2023-07-31)



# 0.1.0-rc1 (2023-07-27)



