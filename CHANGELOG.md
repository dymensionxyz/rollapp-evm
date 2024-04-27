<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the GitHub issue reference in the following format:
* **app:** fixed bech32 on account keeper to not be hardcoded  ([#165](https://github.com/dymensionxyz/rollapp-evm/issues/165)) ([750d1e7](https://github.com/dymensionxyz/rollapp-evm/commit/750d1e70ad052daf7b2942bcecaf0dddfbc17d90))
* **deps:** bumps `block-explorer-rpc-cosmos v1.0.3` & `evm-block-explorer-rpc-cosmos` v1.0.3 ([#142](https://github.com/dymensionxyz/rollapp-evm/issues/142)) ([ea5e5fd](https://github.com/dymensionxyz/rollapp-evm/commit/ea5e5fdc854d5a4fa4079c4d79b79732e78cf9d8))
* **init scripts:** update account-prefix in ibc script ([#190](https://github.com/dymensionxyz/rollapp-evm/issues/190)) ([25be6c3](https://github.com/dymensionxyz/rollapp-evm/commit/25be6c3dda7885870d514438548e10daad45f4d7))
* **local script:** updated default genesis created on extended guide with EIP 3855 ([#183](https://github.com/dymensionxyz/rollapp-evm/issues/183)) ([d201be4](https://github.com/dymensionxyz/rollapp-evm/commit/d201be4ee6757c912ecae568207c1ea358387cae))
* make sure that accounts are not double funded during eibc channel creation ([6e5b6f1](https://github.com/dymensionxyz/rollapp-evm/commit/6e5b6f1ac826ccf926815e7d0d73fe004b5f5842))
* multiple fixes to advance readme features ([#141](https://github.com/dymensionxyz/rollapp-evm/issues/141)) ([469d39f](https://github.com/dymensionxyz/rollapp-evm/commit/469d39fc79591cdae4455839db1546cc5bd9c053))
* **readme:** broken links have been renewed. ([#78](https://github.com/dymensionxyz/rollapp-evm/issues/78)) ([c7df6f2](https://github.com/dymensionxyz/rollapp-evm/commit/c7df6f29c8b9981d7a998be4091d2e96c19647a3))
* **scripts:** add an option to skip evm base fees ([#162](https://github.com/dymensionxyz/rollapp-evm/issues/162)) ([ea51eee](https://github.com/dymensionxyz/rollapp-evm/commit/ea51eee8d66dbba587d6ec00395418a1f08b99a8))
* **scripts:** fix hubgenesis tokens in update genesis ([#172](https://github.com/dymensionxyz/rollapp-evm/issues/172)) ([8d37db8](https://github.com/dymensionxyz/rollapp-evm/commit/8d37db874902eb483293894333416d54bd051e72))
* **scripts:** remove redundant line in setup ibc script ([#161](https://github.com/dymensionxyz/rollapp-evm/issues/161)) ([57d4f17](https://github.com/dymensionxyz/rollapp-evm/commit/57d4f170779dbaeac0e877d1005709d48a4df2f0))
* **test scripts:** make alice and bob keys in .rollap dir ([#147](https://github.com/dymensionxyz/rollapp-evm/issues/147)) ([b6ee646](https://github.com/dymensionxyz/rollapp-evm/commit/b6ee64640af1b528728414bded6a70216a4b5fdf))
* updated block size and evm `no_base_fee` ([#160](https://github.com/dymensionxyz/rollapp-evm/issues/160)) ([876ccad](https://github.com/dymensionxyz/rollapp-evm/commit/876ccad96765d0d3bd279903c552ab483ecf6b9a))

* (<tag>) \#<issue-number> message

Tag must include `sql` if having any changes relate to schema

The issue numbers will later be link-ified during the release process,
so you do not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState
given same genesisState and txList.

If any PR belong to multiple types of change, reference it into all types with only ticket id, no need description (convention)

Ref: https://keepachangelog.com/en/1.0.0/
-->

<!--
Templates for Unreleased:

## Unreleased

### Features

### Improvements

### Bug Fixes

### Client Breaking

### API Breaking

### State Machine Breaking

-->

# Changelog

## Unreleased

### Improvements

- (deps) [#138](https://github.com/dymensionxyz/rollapp-evm/issues/138) Bumps `block-explorer-rpc-cosmos v1.0.2` & `evm-block-explorer-rpc-cosmos v1.0.2`
- (deps) [#151](https://github.com/dymensionxyz/rollapp-evm/issues/151) Bumps `block-explorer-rpc-cosmos v1.0.2` & `evm-block-explorer-rpc-cosmos v1.1.0`

### Bug Fixes

- (deps) [#142](https://github.com/dymensionxyz/rollapp-evm/issues/142) Bumps `block-explorer-rpc-cosmos v1.0.3` & `evm-block-explorer-rpc-cosmos v1.0.3`
