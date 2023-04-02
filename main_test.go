package main

import (
	"reflect"
	"testing"
)

func TestAddBlock(t *testing.T) {
	testCases := []struct {
		name          string
		transactions  []Transaction
		expectedLen   int
		expectedTxLen int
	}{
		{
			name: "Test case 1: Single transaction",
			transactions: []Transaction{
				{
					Sender:   "Alice",
					Receiver: "Bob",
					Amount:   50,
				},
			},
			expectedLen:   2,
			expectedTxLen: 1,
		},
		{
			name: "Test case 2: Multiple transactions",
			transactions: []Transaction{
				{
					Sender:   "Alice",
					Receiver: "Bob",
					Amount:   50,
				},
				{
					Sender:   "Bob",
					Receiver: "Charlie",
					Amount:   30,
				},
			},
			expectedLen:   2,
			expectedTxLen: 2,
		},
		{
			name:          "Test case 3: Empty transactions",
			transactions:  []Transaction{},
			expectedLen:   2,
			expectedTxLen: 0,
		},
		{
			name: "Test case 4: Two transactions with the same sender and receiver",
			transactions: []Transaction{
				{
					Sender:   "Alice",
					Receiver: "Bob",
					Amount:   50,
				},
				{
					Sender:   "Alice",
					Receiver: "Bob",
					Amount:   100,
				},
			},
			expectedLen:   2,
			expectedTxLen: 2,
		},
		{
			name: "Test case 5: Negative amount transaction",
			transactions: []Transaction{
				{
					Sender:   "Alice",
					Receiver: "Bob",
					Amount:   -50,
				},
			},
			expectedLen:   2,
			expectedTxLen: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			blockchain := NewBlockchain()
			blockchain.AddBlock(tc.transactions)

			if len(blockchain.Blocks) != tc.expectedLen {
				t.Errorf("Expected blockchain length after adding a block: %d, got: %d", tc.expectedLen, len(blockchain.Blocks))
			}

			if len(blockchain.Blocks[1].Transactions) != tc.expectedTxLen {
				t.Errorf("Expected transaction length in the new block: %d, got: %d", tc.expectedTxLen, len(blockchain.Blocks[1].Transactions))
			}

			if !reflect.DeepEqual(blockchain.Blocks[1].Transactions, tc.transactions) {
				t.Errorf("Transactions in the block do not match the expected transactions")
			}

			if blockchain.Blocks[1].PrevHash != blockchain.Blocks[0].Hash {
				t.Errorf("Previous hash of the new block does not match the hash of the previous block")
			}
		})
	}
}
