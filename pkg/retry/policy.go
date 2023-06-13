/*
Copyright 2023 Chainguard, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package retry

import (
	"context"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
)

// Key is used as the key for associating information with a context.Context.
type Key struct{}

type Predicate struct {
	Match   func(method string) bool
	Options []grpc_retry.CallOption
}

type Policy []Predicate

func WithPolicy(ctx context.Context, policy Policy) context.Context {
	return context.WithValue(ctx, Key{}, policy)
}

func PolicyFromContext(ctx context.Context) Policy {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		return nil
	}
	return untyped.(Policy)

}
