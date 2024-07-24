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

package v2

import (
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/karalabe/ssz"
)

// Deposit into the consensus layer from the deposit contract in the execution
// layer.
//
//nolint:lll // struct tags.
type Deposit struct {
	// Public key of the validator specified in the deposit.
	Pubkey crypto.BLSPubkey `json:"pubkey"`
	// A staking credentials with
	// 1 byte prefix + 11 bytes padding + 20 bytes address = 32 bytes.
	Credentials types.WithdrawalCredentials `json:"credentials"`
	// Deposit amount in gwei.
	Amount math.Gwei `json:"amount"`
	// Signature of the deposit data.
	Signature crypto.BLSSignature `json:"signature"`
	// Index of the deposit in the deposit contract.
	Index uint64 `json:"index"`
}

// NewDeposit creates a new Deposit instance.
func NewDeposit(
	pubkey crypto.BLSPubkey,
	credentials types.WithdrawalCredentials,
	amount math.Gwei,
	signature crypto.BLSSignature,
	index uint64,
) *Deposit {
	return &Deposit{
		Pubkey:      pubkey,
		Credentials: credentials,
		Amount:      amount,
		Signature:   signature,
		Index:       index,
	}
}

// New creates a new Deposit instance.
func (d *Deposit) New(
	pubkey crypto.BLSPubkey,
	credentials types.WithdrawalCredentials,
	amount math.Gwei,
	signature crypto.BLSSignature,
	index uint64,
) *Deposit {
	return NewDeposit(
		pubkey, credentials, amount, signature, index,
	)
}

// VerifySignature verifies the deposit data and signature.
func (d *Deposit) VerifySignature(
	forkData *types.ForkData,
	domainType common.DomainType,
	signatureVerificationFn func(
		pubkey crypto.BLSPubkey, message []byte, signature crypto.BLSSignature,
	) error,
) error {
	return (&types.DepositMessage{
		Pubkey:      d.Pubkey,
		Credentials: d.Credentials,
		Amount:      d.Amount,
	}).VerifyCreateValidator(
		forkData, d.Signature,
		domainType, signatureVerificationFn,
	)
}

// DefineSSZ defines the SSZ encoding for the Deposit object.
func (d *Deposit) DefineSSZ(codec *ssz.Codec) {
	ssz.DefineStaticBytes(codec, &d.Pubkey)
	ssz.DefineStaticBytes(codec, &d.Credentials)
	ssz.DefineUint64(codec, &d.Amount)
	ssz.DefineStaticBytes(codec, &d.Signature)
	ssz.DefineUint64(codec, &d.Index)
}

// SizeSSZ returns the size of the Deposit object in SSZ encoding.
func (d *Deposit) SizeSSZ() uint32 {
	return 48 + // Pubkey (BLSPubkey) size
		32 + // Credentials (WithdrawalCredentials) size
		8 + // Amount (Gwei) size
		96 + // Signature (BLSSignature) size
		8 // Index (uint64) size
}

// MarshalSSZ marshals the Deposit object into SSZ format.
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, d.SizeSSZ())
	return buf, ssz.EncodeToBytes(buf, d)
}

// MarshalSSZTo marshals the Deposit object into a pre-allocated byte slice.
func (d *Deposit) MarshalSSZTo(dst []byte) ([]byte, error) {
	return dst, ssz.EncodeToBytes(dst, d)
}

// UnmarshalSSZ unmarshals the Deposit object from SSZ format.
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
	return ssz.DecodeFromBytes(buf, d)
}

// HashTreeRoot computes the Merkleization of the Deposit object.
func (d *Deposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashSequential(d), nil
}

// GetAmount returns the deposit amount in gwei.
func (d *Deposit) GetAmount() math.Gwei {
	return d.Amount
}

// GetPubkey returns the public key of the validator specified in the deposit.
func (d *Deposit) GetPubkey() crypto.BLSPubkey {
	return d.Pubkey
}

// GetIndex returns the index of the deposit in the deposit contract.
func (d *Deposit) GetIndex() uint64 {
	return d.Index
}

// GetSignature returns the signature of the deposit data.
func (d *Deposit) GetSignature() crypto.BLSSignature {
	return d.Signature
}

// GetWithdrawalCredentials returns the staking credentials of the deposit.
func (d *Deposit) GetWithdrawalCredentials() types.WithdrawalCredentials {
	return d.Credentials
}
