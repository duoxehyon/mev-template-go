## MEV Template for Go
This repo comes with the following features implemented.

- [x] A robust structure for developing long tail and short tail strategies
- [x] Mempool monitoring, decoding
- [x] Using GraphQL for querying uniswap v2 pairs
- [x] Contract bindings

# Cmd
for all binaries, eg: bot, data analysis tool , etc... should be the logic should be implemented

# Contract modules
implementations of data collection, math functions, etc... per contract

# Executor
for building transaction payload, memory sharing to transfer transaction to geth, etc...

# Recon
for getting live data from blockchain. pending transactions, blocks, contract data, etc...

# Types
global types which require module level access

# Logic
internal bot logic
