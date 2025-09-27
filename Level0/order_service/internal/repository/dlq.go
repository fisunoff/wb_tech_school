package repository

import "context"

type DeadLetterQueue interface {
	Send(ctx context.Context, message []byte) error
	Close() error
}
