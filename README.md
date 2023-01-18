# MEV Template for Go
Simple, robust MEV Bot template for Go.This repo comes with the following features and much more.

- [x] A robust structure for developing long tail and short tail strategies
- [x] Mempool monitoring, decoding
- [x] Using GraphQL for querying uniswap v2 pairs
- [x] Uniswap v2 implementation 
- [x] Contract bindings

## `/Cmd`
for all binaries, eg: bot, data analysis tool , etc... 

## `/Contract modules`
implementations of data collection, math functions, etc... per contract

## `/Executor`
for building transaction payload, memory sharing to transfer transaction to geth, etc...

## `/Recon`
for getting live data from blockchain. pending transactions, blocks, contract data, etc...

## `/Types`
global types which require module level access

## `/Logic`
internal bot logic

## Building and testing
run ```python ./bot-builder.py build ``` for building the bot, test and other commands are to be implemented by the user
