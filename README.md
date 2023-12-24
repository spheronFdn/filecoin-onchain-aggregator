# Filecoin On-Chain Aggregator

## Description

Currently smart contract deals on FVM are not aggregated, and mostly small size (under 32G), which is not very attractive for Storage Providers (note: we refer to these small deals as subpiece deals below, as compared to the aggregated full piece). Aggregating these small smart contract deals would make it much more appealing for Storage Providers to pick up the deal. There needs to be a way for smart contract deals to be aggregated into bigger deals in a decentralized and transparent way, allowing the storage market to function effectively on FVM.

Additionally, without bundling small deals into larger ones, economically it won’t make sense for Storage Provider to store the data too since most of the sector size which could have been used to store verified deals would be wasted if they use the sector to store only small pieces of data (see analysis here: Fil sectors with small deals).

Therefore, we believe it’s crucial to have a decentralized deal aggregator for FVM smart contract deals to work at scale. And with the recent development of FVM, PoDSI, and IPC, the aggregator can perform aggregation onchain and also show to the clients that the subpiece data is included in the bigger deal in a trustless, verifiable manner.

### Main flow:

- Client submits a subpiece CID (commPc) with metadata (download URI) directly to a bundler contract;
- Aggregator(Dealer in our case) watches the bundler contract, when there are enough pieces to produce a 32GiB aggregated piece CID (CommPa), the aggregator downloads all subpieces for the aggregated piece from the fetch link;
- The aggregator will combine the subpiece CIDs (CommPc) into 32GiB CommP (CommPa) by performing hashing within the contract, including the Data Segment Index onchain (DSI is built and hashed onchain as a part of the resulting CommPa, which can later be used for subpiece retrieval);
- Once the aggregated CommPa and fetch link is ready, the aggregator can make a deal directly by calling Boost API (or onchain triggers a Smart Contract deal following deal making FRC which emits dealproposal event for the aggregated piece) on mainnet where Boost SPs are listening to claim and store the aggregated piece.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Installation and Usage

Provide step-by-step instructions on how to install and set up your project. Include any dependencies that need to be installed.

```bash
# Example installation command
$ go install
$ go run cmd/dealer/main.go
```
