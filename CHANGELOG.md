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
- (deps) [#151](https://github.com/dymensionxyz/rollapp-evm/issues/151) Bumps `block-explorer-rpc-cosmos v1.1.0` & `evm-block-explorer-rpc-cosmos v1.1.0`
- (deps) [#167](https://github.com/dymensionxyz/rollapp-evm/issues/167) Bumps `block-explorer-rpc-cosmos v1.1.2`

### Bug Fixes

- (deps) [#142](https://github.com/dymensionxyz/rollapp-evm/issues/142) Bumps `block-explorer-rpc-cosmos v1.0.3` & `evm-block-explorer-rpc-cosmos v1.0.3`
