// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package ethclient

import (
	"context"
	"encoding/json"

	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Eth1Client is a struct that holds the Ethereum 1 client and
// its configuration.
type Eth1Client[
	ExecutionPayloadT interface {
		json.Marshaler
		json.Unmarshaler
		Version() uint32
		Empty(uint32) ExecutionPayloadT
	},
	PayloadAttributesT interface {
		Version() uint32
		Empty(uint32) PayloadAttributesT
	},
] struct {
	*ethclient.Client
}

// NewEth1Client creates a new Ethereum 1 client with the provided
// context and options.
func NewEth1Client[
	ExecutionPayloadT interface {
		json.Marshaler
		json.Unmarshaler
		Version() uint32
		Empty(uint32) ExecutionPayloadT
	},
	PayloadAttributesT interface {
		Version() uint32
		Empty(uint32) PayloadAttributesT
	},
](client *ethclient.Client) (*Eth1Client[ExecutionPayloadT, PayloadAttributesT], error) {
	c := &Eth1Client[ExecutionPayloadT, PayloadAttributesT]{
		Client: client,
	}
	return c, nil
}

// NewFromRPCClient creates a new Ethereum 1 client from an RPC client.
func NewFromRPCClient[
	ExecutionPayloadT interface {
		json.Marshaler
		json.Unmarshaler
		Version() uint32
		Empty(uint32) ExecutionPayloadT
	},
	PayloadAttributesT interface {
		Version() uint32
		Empty(uint32) PayloadAttributesT
	},
](rpcClient *rpc.Client) (*Eth1Client[ExecutionPayloadT, PayloadAttributesT], error) {
	return NewEth1Client[ExecutionPayloadT, PayloadAttributesT](
		ethclient.NewClient(rpcClient),
	)
}

// ExecutionBlockByHash fetches an execution engine block by hash by calling
// eth_blockByHash via JSON-RPC.
func (s *Eth1Client[ExecutionPayloadT, PayloadAttributesT]) ExecutionBlockByHash(
	ctx context.Context,
	hash common.ExecutionHash,
	withTxs bool,
) (*engineprimitives.Block, error) {
	result := &engineprimitives.Block{}
	err := s.Client.Client().CallContext(
		ctx, result, BlockByHashMethod, hash, withTxs)
	return result, err
}

// ExecutionBlockByNumber fetches an execution engine block by number
// by calling eth_getBlockByNumber via JSON-RPC.
func (s *Eth1Client[ExecutionPayloadT, PayloadAttributesT]) ExecutionBlockByNumber(
	ctx context.Context,
	num rpc.BlockNumber,
	withTxs bool,
) (*engineprimitives.Block, error) {
	result := &engineprimitives.Block{}
	err := s.Client.Client().CallContext(
		ctx, result, BlockByNumberMethod, num, withTxs)
	return result, err
}

// GetClientVersionV1 calls the engine_getClientVersionV1 method via JSON-RPC.
func (s *Eth1Client[ExecutionPayloadT, PayloadAttributesT]) GetClientVersionV1(
	ctx context.Context,
) ([]engineprimitives.ClientVersionV1, error) {
	result := make([]engineprimitives.ClientVersionV1, 0)
	if err := s.Client.Client().CallContext(
		ctx, &result, GetClientVersionV1, nil,
	); err != nil {
		return nil, err
	}
	return result, nil
}

// ExchangeCapabilities calls the engine_exchangeCapabilities method via
// JSON-RPC.
func (s *Eth1Client[ExecutionPayloadT, PayloadAttributesT]) ExchangeCapabilities(
	ctx context.Context,
	capabilities []string,
) ([]string, error) {
	result := make([]string, 0)
	if err := s.Client.Client().CallContext(
		ctx, &result, ExchangeCapabilities, &capabilities,
	); err != nil {
		return nil, err
	}
	return result, nil
}
