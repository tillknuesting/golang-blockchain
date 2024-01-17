# Blockchain in Go

This is a simple blockchain implementation in Go. It consists of:

- Transactions with a sender, receiver, and amount
- Blocks that contain transactions and the hash of the previous block
- A Blockchain that consists of a series of blocks

## Features

- Generate new blocks containing transactions
- Calculate the hash for blocks 
- Validate blocks by comparing hashes 
- Maintain an immutable transaction history 

## Code Overview

The key components:

- `Transaction` struct - A basic transaction between a sender and receiver 
- `Block` struct - Contains transactions, metadata, and cryptographic hash 
- `Blockchain` struct - Sequence of blocks with hash pointers  

Core functionality:

- `NewBlock()` - Create and initialize new blocks
- `calculateHash()` - Generate SHA-256 hash for blocks
- `AddBlock()` - Append new blocks to the blockchain

The `main()` function demonstrates creating a blockchain, adding blocks, and printing the results.

## Running the Code

To run the example:

```
go run blockchain.go
```

This will initialize a blockchain, add some sample blocks, and print details to stdout.

## Next Steps

Possible improvements:

- Persist the blockchain data structure 
- Implement consensus algorithms
- Add transaction signatures and verification
- Optimize hashing and serialization
- Create miners that compete to add blocks  
- Enable peer-to-peer network communication
- Add RPC interface for queries and updates

## License

This code is released under the MIT License. Feel free to modify and reuse according to the terms.
