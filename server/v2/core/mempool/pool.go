package mempool

import (
	"context"

	"github.com/cosmos/cosmos-sdk/server/v2/core/transaction"
)

// Mempool defines the required methods of an application's mempool.
type Mempool[T any] interface {
	// Insert attempts to insert a Tx into the app-side mempool returning
	// an error upon failure. Insert will validate the transaction using the txValidator
	Insert(context.Context, T) error

	// GetTxs returns a list of transactions to add in a block
	// size specifies the size of the block left for transactions
	GetTxs(context.Context, uint32, TxSizeFn) ([]T, error)

	// Remove attempts to remove a transaction from the mempool, returning an error
	// upon failure.
	Remove(T) error
}

// TxSizeFn defines the function type for calculating the size of a transaction.
type TxSizeFn func(context.Context, transaction.Tx) (uint64, error)