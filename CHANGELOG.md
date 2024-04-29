# [](https://github.com/dymensionxyz/rollapp-evm/compare/v2.2.0-alpha...v) (2024-04-29)


### Bug Fixes

* **app:** fixed bech32 on account keeper to not be hardcoded  ([#165](https://github.com/dymensionxyz/rollapp-evm/issues/165)) ([750d1e7](https://github.com/dymensionxyz/rollapp-evm/commit/750d1e70ad052daf7b2942bcecaf0dddfbc17d90))
* **app:** initialize transferkeeper before denommetadatakeeper to avoid nil pointer error ([#194](https://github.com/dymensionxyz/rollapp-evm/issues/194)) ([8050f59](https://github.com/dymensionxyz/rollapp-evm/commit/8050f59de96fde9130b00738048462860ce51be8))
* **deps:** bumps `block-explorer-rpc-cosmos v1.0.3` & `evm-block-explorer-rpc-cosmos` v1.0.3 ([#142](https://github.com/dymensionxyz/rollapp-evm/issues/142)) ([ea5e5fd](https://github.com/dymensionxyz/rollapp-evm/commit/ea5e5fdc854d5a4fa4079c4d79b79732e78cf9d8))
* **init scripts:** update account-prefix in ibc script ([#190](https://github.com/dymensionxyz/rollapp-evm/issues/190)) ([25be6c3](https://github.com/dymensionxyz/rollapp-evm/commit/25be6c3dda7885870d514438548e10daad45f4d7))
* **local script:** updated default genesis created on extended guide with EIP 3855 ([#183](https://github.com/dymensionxyz/rollapp-evm/issues/183)) ([d201be4](https://github.com/dymensionxyz/rollapp-evm/commit/d201be4ee6757c912ecae568207c1ea358387cae))
* make sure that accounts are not double funded during eibc channel creation ([#191](https://github.com/dymensionxyz/rollapp-evm/issues/191)) ([0d6bb2d](https://github.com/dymensionxyz/rollapp-evm/commit/0d6bb2de7c667191352bdac8e00631165b675250))
* multiple fixes to advance readme features ([#141](https://github.com/dymensionxyz/rollapp-evm/issues/141)) ([469d39f](https://github.com/dymensionxyz/rollapp-evm/commit/469d39fc79591cdae4455839db1546cc5bd9c053))
* **readme:** broken links have been renewed. ([#78](https://github.com/dymensionxyz/rollapp-evm/issues/78)) ([c7df6f2](https://github.com/dymensionxyz/rollapp-evm/commit/c7df6f29c8b9981d7a998be4091d2e96c19647a3))
* remove deprecated denommetadata param ([#189](https://github.com/dymensionxyz/rollapp-evm/issues/189)) ([927f7f9](https://github.com/dymensionxyz/rollapp-evm/commit/927f7f92ab7c84d39931a1923780d8c373dfce74))
* **scripts:** add an option to skip evm base fees ([#162](https://github.com/dymensionxyz/rollapp-evm/issues/162)) ([ea51eee](https://github.com/dymensionxyz/rollapp-evm/commit/ea51eee8d66dbba587d6ec00395418a1f08b99a8))
* **scripts:** fix hubgenesis tokens in update genesis ([#172](https://github.com/dymensionxyz/rollapp-evm/issues/172)) ([8d37db8](https://github.com/dymensionxyz/rollapp-evm/commit/8d37db874902eb483293894333416d54bd051e72))
* **scripts:** remove redundant line in setup ibc script ([#161](https://github.com/dymensionxyz/rollapp-evm/issues/161)) ([57d4f17](https://github.com/dymensionxyz/rollapp-evm/commit/57d4f170779dbaeac0e877d1005709d48a4df2f0))
* **test scripts:** make alice and bob keys in .rollap dir ([#147](https://github.com/dymensionxyz/rollapp-evm/issues/147)) ([b6ee646](https://github.com/dymensionxyz/rollapp-evm/commit/b6ee64640af1b528728414bded6a70216a4b5fdf))
* updated block size and evm `no_base_fee` ([#160](https://github.com/dymensionxyz/rollapp-evm/issues/160)) ([876ccad](https://github.com/dymensionxyz/rollapp-evm/commit/876ccad96765d0d3bd279903c552ab483ecf6b9a))


### Features

* add swagger config and make scripts ([#130](https://github.com/dymensionxyz/rollapp-evm/issues/130)) ([41718e4](https://github.com/dymensionxyz/rollapp-evm/commit/41718e4d4098e6bf18117c31b514b2cc226a331f))
* **be:** integrate block explorer Json-RPC server ([#132](https://github.com/dymensionxyz/rollapp-evm/issues/132)) ([d73b1c4](https://github.com/dymensionxyz/rollapp-evm/commit/d73b1c451b93f04a1db5a73c4c8c78fc21729208))
* **ci:** Add changelog log auto update workflow ([#176](https://github.com/dymensionxyz/rollapp-evm/issues/176)) ([f58feaa](https://github.com/dymensionxyz/rollapp-evm/commit/f58feaaea83b17d2258d1025a72c9b832922b7a9))



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



