# MEV Template for Go

> Inspired by [DeGatchi](https://twitter.com/DeGatchi).

Simple, robust MEV Bot template for Go.This repo comes with the following features and more.

- [x] A robust structure for developing long tail and short tail strategies
- [x] Mempool monitoring, decoding
- [x] Using GraphQL for querying uniswap v2 pairs
- [x] Uniswap v2 implementation 
  - [x] Uniswap v2 math functions (`getAmountIn` and `getAmountOut`)
  - [x] Query of top 1000 pairs in eth sorted by liquidity
  - [x] Transaction input decoding 
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

## Quick Start

- Get rpc urls.
- Import your private key.
- For testing, run `python ./bot-builder.py test` (implementation expected by user)
- For production, run `python ./bog-builder.py build`
