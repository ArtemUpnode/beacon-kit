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

package types

import (
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
)

type BlockHeaderResponse[BlockHeaderT any] struct {
	Root      common.Root                `json:"root"`
	Canonical bool                       `json:"canonical"`
	Header    *BlockHeader[BlockHeaderT] `json:"header"`
}

type BlockHeader[BlockHeaderT any] struct {
	Message   BlockHeaderT `json:"message"`
	Signature bytes.B48    `json:"signature"`
}

type GenesisData struct {
	GenesisTime           string      `json:"genesis_time"`
	GenesisValidatorsRoot common.Root `json:"genesis_validators_root"`
	GenesisForkVersion    string      `json:"genesis_fork_version"`
}

type RootData struct {
	Root common.Root `json:"root"`
}

type ValidatorBalanceData struct {
	Index   uint64 `json:"index"`
	Balance uint64 `json:"balance"`
}

func (vbd ValidatorBalanceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Index   string `json:"index"`
		Balance string `json:"balance"`
	}{
		Index:   strconv.FormatUint(vbd.Index, 10),
		Balance: strconv.FormatUint(vbd.Balance, 10),
	})
}

type ValidatorData struct {
	ValidatorBalanceData
	Status    string    `json:"status"`
	Validator Validator `json:"validator"`
}

// MarshalJSON implements json.Marshal for ValidatorResponse.
//
//nolint:lll
func (vd ValidatorData) MarshalJSON() ([]byte, error) {
	type ValidatorJSON struct {
		PublicKey                  string `json:"pubkey"`
		WithdrawalCredentials      string `json:"withdrawal_credentials"`
		EffectiveBalance           string `json:"effective_balance"`
		Slashed                    bool   `json:"slashed"`
		ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
		ActivationEpoch            string `json:"activation_epoch"`
		ExitEpoch                  string `json:"exit_epoch"`
		WithdrawableEpoch          string `json:"withdrawable_epoch"`
	}

	type ResponseJSON struct {
		Index     string        `json:"index"`
		Balance   string        `json:"balance"`
		Status    string        `json:"status"`
		Validator ValidatorJSON `json:"validator"`
	}

	withdrawalCredentials := vd.Validator.GetWithdrawalCredentials()
	withdrawalCredentialsBytes := withdrawalCredentials[:]

	return json.Marshal(ResponseJSON{
		Index:   strconv.FormatUint(vd.Index, 10),
		Balance: strconv.FormatUint(vd.Balance, 10),
		Status:  vd.Status,
		Validator: ValidatorJSON{
			PublicKey:                  vd.Validator.GetPubkey().String(),
			WithdrawalCredentials:      "0x" + hex.EncodeToString(withdrawalCredentialsBytes),
			EffectiveBalance:           strconv.FormatUint(uint64(vd.Validator.GetEffectiveBalance()), 10),
			Slashed:                    vd.Validator.IsSlashed(),
			ActivationEligibilityEpoch: strconv.FormatUint(uint64(vd.Validator.GetActivationEligibilityEpoch()), 10),
			ActivationEpoch:            strconv.FormatUint(uint64(vd.Validator.GetActivationEpoch()), 10),
			ExitEpoch:                  strconv.FormatUint(uint64(vd.Validator.GetExitEpoch()), 10),
			WithdrawableEpoch:          strconv.FormatUint(uint64(vd.Validator.GetWithdrawableEpoch()), 10),
		},
	})
}

//nolint:staticcheck // todo: figure this out.
type CommitteeData struct {
	Index      uint64   `json:"index,string"`
	Slot       uint64   `json:"slot,string"`
	Validators []uint64 `json:"validators,string"`
}

type BlockRewardsData struct {
	ProposerIndex     uint64 `json:"proposer_index,string"`
	Total             uint64 `json:"total,string"`
	Attestations      uint64 `json:"attestations,string"`
	SyncAggregate     uint64 `json:"sync_aggregate,string"`
	ProposerSlashings uint64 `json:"proposer_slashings,string"`
	AttesterSlashings uint64 `json:"attester_slashings,string"`
}

type RandaoData struct {
	Randao common.Bytes32 `json:"randao"`
}

type ForkResponse struct {
	Fork
}

func (fr ForkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PreviousVersion string `json:"previous_version"`
		CurrentVersion  string `json:"current_version"`
		Epoch           string `json:"epoch"`
	}{
		PreviousVersion: fr.GetPreviousVersion().String(),
		CurrentVersion:  fr.GetCurrentVersion().String(),
		Epoch:           strconv.FormatUint(fr.GetEpoch().Unwrap(), 10),
	})
}
