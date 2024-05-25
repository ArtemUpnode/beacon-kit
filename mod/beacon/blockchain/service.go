// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package blockchain

import (
	"context"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
)

// Service is the blockchain service.
type Service[
	AvailabilityStoreT AvailabilityStore[
		types.BeaconBlockBody, BlobSidecarsT,
	],
	ReadOnlyBeaconStateT ReadOnlyBeaconState[ReadOnlyBeaconStateT],
	BlobSidecarsT BlobSidecars,
	DepositStoreT DepositStore,
] struct {
	// sb represents the backend storage for beacon states and associated
	// sidecars.
	sb StorageBackend[
		AvailabilityStoreT,
		ReadOnlyBeaconStateT,
		BlobSidecarsT,
		DepositStoreT,
	]
	// logger is used for logging messages in the service.
	logger log.Logger[any]
	// cs holds the chain specifications.
	cs primitives.ChainSpec
	// ee is the execution engine responsible for processing execution payloads.
	ee ExecutionEngine
	// dc is a connection to the deposit contract.
	dc DepositContract
	// lb is a local builder for constructing new beacon states.
	lb LocalBuilder[ReadOnlyBeaconStateT]
	// bp is the blob processor for processing incoming blobs.
	bp BlobProcessor[AvailabilityStoreT, BlobSidecarsT]
	// sp is the state processor for beacon blocks and states.
	sp StateProcessor[
		types.BeaconBlock,
		ReadOnlyBeaconStateT,
		BlobSidecarsT,
		*transition.Context,
	]
	// metrics is the metrics for the service.
	metrics *chainMetrics
}

// NewService creates a new validator service.
func NewService[
	AvailabilityStoreT AvailabilityStore[
		types.BeaconBlockBody, BlobSidecarsT,
	],
	ReadOnlyBeaconStateT ReadOnlyBeaconState[ReadOnlyBeaconStateT],
	BlobSidecarsT BlobSidecars,
	DepositStoreT DepositStore,
](
	sb StorageBackend[
		AvailabilityStoreT,
		ReadOnlyBeaconStateT, BlobSidecarsT, DepositStoreT],
	logger log.Logger[any],
	cs primitives.ChainSpec,
	ee ExecutionEngine,
	lb LocalBuilder[ReadOnlyBeaconStateT],
	bp BlobProcessor[
		AvailabilityStoreT,
		BlobSidecarsT,
	],
	sp StateProcessor[
		types.BeaconBlock, ReadOnlyBeaconStateT,
		BlobSidecarsT, *transition.Context,
	],
	dc DepositContract,
	ts TelemetrySink,
) *Service[
	AvailabilityStoreT, ReadOnlyBeaconStateT,
	BlobSidecarsT, DepositStoreT,
] {
	return &Service[
		AvailabilityStoreT, ReadOnlyBeaconStateT,
		BlobSidecarsT, DepositStoreT,
	]{
		sb:      sb,
		logger:  logger,
		cs:      cs,
		ee:      ee,
		lb:      lb,
		bp:      bp,
		sp:      sp,
		dc:      dc,
		metrics: newChainMetrics(ts),
	}
}

// Name returns the name of the service.
func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) Name() string {
	return "blockchain"
}

func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) Start(
	context.Context,
) error {
	return nil
}

func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) Status() error {
	return nil
}

func (s *Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) WaitForHealthy(
	context.Context,
) {
}

// TODO: Remove
func (s Service[
	AvailabilityStoreT,
	ReadOnlyBeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
]) StateFromContext(
	ctx context.Context,
) ReadOnlyBeaconStateT {
	return s.sb.StateFromContext(ctx)
}
