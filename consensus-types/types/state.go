// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, Berachain Foundation. All rights reserved.
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
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types

import (
	"github.com/berachain/beacon-kit/primitives/common"
	"github.com/berachain/beacon-kit/primitives/math"
	"github.com/berachain/beacon-kit/primitives/version"
	fastssz "github.com/ferranbt/fastssz"
	"github.com/karalabe/ssz"
)

// BeaconState represents the entire state of the beacon chain.
type BeaconState struct {
	// Versioning
	GenesisValidatorsRoot common.Root `json:"genesis_validators_root,omitempty"`
	Slot                  math.Slot   `json:"slot,omitempty"`
	Fork                  *Fork       `json:"fork,omitempty"`

	// History
	LatestBlockHeader *BeaconBlockHeader `json:"latest_block_header,omitempty"`
	BlockRoots        []common.Root      `json:"block_roots,omitempty"`
	StateRoots        []common.Root      `json:"state_roots,omitempty"`

	// Eth1
	Eth1Data                     *Eth1Data               `json:"eth1_data,omitempty"`
	Eth1DepositIndex             uint64                  `json:"eth1_deposit_index,omitempty"`
	LatestExecutionPayloadHeader *ExecutionPayloadHeader `json:"latest_execution_payload_header,omitempty"`

	// Registry
	Validators []*Validator `json:"validators,omitempty"`
	Balances   []uint64     `json:"balances,omitempty"`

	// Randomness
	RandaoMixes []common.Bytes32 `json:"randao_mixes,omitempty"`

	// Withdrawals
	NextWithdrawalIndex          uint64              `json:"next_withdrawal_index,omitempty"`
	NextWithdrawalValidatorIndex math.ValidatorIndex `json:"next_withdrawal_validator_index,omitempty"`

	// Slashing
	Slashings     []math.Gwei `json:"slashings,omitempty"`
	TotalSlashing math.Gwei   `json:"total_slashing,omitempty"`
}

/* -------------------------------------------------------------------------- */
/*                                     SSZ                                    */
/* -------------------------------------------------------------------------- */

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object.
func (st *BeaconState) SizeSSZ(siz *ssz.Sizer, fixed bool) uint32 {
	var size uint32 = 300

	if fixed {
		return size
	}

	// Dynamic size fields
	size += ssz.SizeSliceOfStaticBytes(siz, st.BlockRoots)
	size += ssz.SizeSliceOfStaticBytes(siz, st.StateRoots)
	size += ssz.SizeDynamicObject(siz, st.LatestExecutionPayloadHeader)
	size += ssz.SizeSliceOfStaticObjects(siz, st.Validators)
	size += ssz.SizeSliceOfUint64s(siz, st.Balances)
	size += ssz.SizeSliceOfStaticBytes(siz, st.RandaoMixes)
	size += ssz.SizeSliceOfUint64s(siz, st.Slashings)

	return size
}

// DefineSSZ defines the SSZ encoding for the BeaconState object.
//
//nolint:mnd // TODO: get from accessible chainspec field params
func (st *BeaconState) DefineSSZ(codec *ssz.Codec) {
	// Versioning
	ssz.DefineStaticBytes(codec, &st.GenesisValidatorsRoot)
	ssz.DefineUint64(codec, &st.Slot)
	ssz.DefineStaticObject(codec, &st.Fork)

	// History
	ssz.DefineStaticObject(codec, &st.LatestBlockHeader)
	ssz.DefineSliceOfStaticBytesOffset(codec, &st.BlockRoots, 8192)
	ssz.DefineSliceOfStaticBytesOffset(codec, &st.StateRoots, 8192)

	// Eth1
	ssz.DefineStaticObject(codec, &st.Eth1Data)
	ssz.DefineUint64(codec, &st.Eth1DepositIndex)
	ssz.DefineDynamicObjectOffset(codec, &st.LatestExecutionPayloadHeader)

	// Registry
	ssz.DefineSliceOfStaticObjectsOffset(codec, &st.Validators, 1099511627776)
	ssz.DefineSliceOfUint64sOffset(codec, &st.Balances, 1099511627776)

	// Randomness
	ssz.DefineSliceOfStaticBytesOffset(codec, &st.RandaoMixes, 65536)

	// Withdrawals
	ssz.DefineUint64(codec, &st.NextWithdrawalIndex)
	ssz.DefineUint64(codec, &st.NextWithdrawalValidatorIndex)

	// // Slashing
	ssz.DefineSliceOfUint64sOffset(codec, &st.Slashings, 1099511627776)
	ssz.DefineUint64(codec, (*uint64)(&st.TotalSlashing))

	// Dynamic content
	ssz.DefineSliceOfStaticBytesContent(codec, &st.BlockRoots, 8192)
	ssz.DefineSliceOfStaticBytesContent(codec, &st.StateRoots, 8192)
	ssz.DefineDynamicObjectContent(codec, &st.LatestExecutionPayloadHeader)
	ssz.DefineSliceOfStaticObjectsContent(codec, &st.Validators, 1099511627776)
	ssz.DefineSliceOfUint64sContent(codec, &st.Balances, 1099511627776)
	ssz.DefineSliceOfStaticBytesContent(codec, &st.RandaoMixes, 65536)
	ssz.DefineSliceOfUint64sContent(codec, &st.Slashings, 1099511627776)
}

// MarshalSSZ marshals the BeaconState into SSZ format.
func (st *BeaconState) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, ssz.Size(st))
	return buf, ssz.EncodeToBytes(buf, st)
}

// UnmarshalSSZ unmarshals the BeaconState from SSZ format.
func (st *BeaconState) UnmarshalSSZ(buf []byte) error {
	return ssz.DecodeFromBytes(buf, st)
}

// HashTreeRoot computes the Merkleization of the BeaconState.
func (st *BeaconState) HashTreeRoot() common.Root {
	return ssz.HashConcurrent(st)
}

/* -------------------------------------------------------------------------- */
/*                                   FastSSZ                                  */
/* -------------------------------------------------------------------------- */

func (st *BeaconState) MarshalSSZTo(
	dst []byte,
) ([]byte, error) {
	bz, err := st.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	dst = append(dst, bz...)
	return dst, nil
}

// HashTreeRootWith ssz hashes the BeaconState object with a hasher.
//
//nolint:mnd,funlen,gocognit // todo fix.
func (st *BeaconState) HashTreeRootWith(
	hh fastssz.HashWalker,
) error {
	indx := hh.Index()

	// Field (0) 'GenesisValidatorsRoot'
	hh.PutBytes(st.GenesisValidatorsRoot[:])

	// Field (1) 'Slot'
	hh.PutUint64(uint64(st.Slot))

	// Field (2) 'Fork'
	if st.Fork == nil {
		st.Fork = &Fork{}
	}
	if err := st.Fork.HashTreeRootWith(hh); err != nil {
		return err
	}

	// Field (3) 'LatestBlockHeader'
	if st.LatestBlockHeader == nil {
		st.LatestBlockHeader = &BeaconBlockHeader{}
	}
	if err := st.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return err
	}

	// Field (4) 'BlockRoots'
	if size := len(st.BlockRoots); size > 8192 {
		return fastssz.ErrListTooBigFn("BeaconState.BlockRoots", size, 8192)
	}
	subIndx := hh.Index()
	for _, i := range st.BlockRoots {
		hh.Append(i[:])
	}
	numItems := uint64(len(st.BlockRoots))
	hh.MerkleizeWithMixin(subIndx, numItems, 8192)

	// Field (5) 'StateRoots'
	if size := len(st.StateRoots); size > 8192 {
		return fastssz.ErrListTooBigFn("BeaconState.StateRoots", size, 8192)
	}
	subIndx = hh.Index()
	for _, i := range st.StateRoots {
		hh.Append(i[:])
	}
	numItems = uint64(len(st.StateRoots))
	hh.MerkleizeWithMixin(subIndx, numItems, 8192)

	// Field (6) 'Eth1Data'
	if st.Eth1Data == nil {
		st.Eth1Data = &Eth1Data{}
	}
	if err := st.Eth1Data.HashTreeRootWith(hh); err != nil {
		return err
	}

	// Field (7) 'Eth1DepositIndex'
	hh.PutUint64(st.Eth1DepositIndex)

	// Field (8) 'LatestExecutionPayloadHeader'
	if st.LatestExecutionPayloadHeader == nil {
		// TODO(pectra): Remove the hardcoded Deneb value and use a retrieved time from beaconState
		st.LatestExecutionPayloadHeader = NewEmptyExecutionPayloadHeaderWithVersion(version.Deneb())
	}
	if err := st.LatestExecutionPayloadHeader.HashTreeRootWith(hh); err != nil {
		return err
	}

	// Field (9) 'Validators'
	subIndx = hh.Index()
	num := uint64(len(st.Validators))
	if num > 1099511627776 {
		return fastssz.ErrIncorrectListSize
	}
	for _, elem := range st.Validators {
		if err := elem.HashTreeRootWith(hh); err != nil {
			return err
		}
	}
	hh.MerkleizeWithMixin(subIndx, num, 1099511627776)

	// Field (10) 'Balances'
	if size := len(st.Balances); size > 1099511627776 {
		return fastssz.ErrListTooBigFn(
			"BeaconState.Balances",
			size,
			1099511627776,
		)
	}
	subIndx = hh.Index()
	for _, i := range st.Balances {
		hh.AppendUint64(i)
	}
	hh.FillUpTo32()
	numItems = uint64(len(st.Balances))
	hh.MerkleizeWithMixin(
		subIndx,
		numItems,
		fastssz.CalculateLimit(1099511627776, numItems, 8),
	)

	// Field (11) 'RandaoMixes'
	if size := len(st.RandaoMixes); size > 65536 {
		return fastssz.ErrListTooBigFn("BeaconState.RandaoMixes", size, 65536)
	}
	subIndx = hh.Index()
	for _, i := range st.RandaoMixes {
		hh.Append(i[:])
	}
	numItems = uint64(len(st.RandaoMixes))
	hh.MerkleizeWithMixin(subIndx, numItems, 65536)

	// Field (12) 'NextWithdrawalIndex'
	hh.PutUint64(st.NextWithdrawalIndex)

	// Field (13) 'NextWithdrawalValidatorIndex'
	hh.PutUint64(uint64(st.NextWithdrawalValidatorIndex))

	// Field (14) 'Slashings'
	if size := len(st.Slashings); size > 1099511627776 {
		return fastssz.ErrListTooBigFn(
			"BeaconState.Slashings",
			size,
			1099511627776,
		)
	}
	subIndx = hh.Index()
	for _, i := range st.Slashings {
		hh.AppendUint64(uint64(i))
	}
	hh.FillUpTo32()
	numItems = uint64(len(st.Slashings))
	hh.MerkleizeWithMixin(
		subIndx,
		numItems,
		fastssz.CalculateLimit(1099511627776, numItems, 8),
	)

	// Field (15) 'TotalSlashing'
	hh.PutUint64(uint64(st.TotalSlashing))

	hh.Merkleize(indx)
	return nil
}

// GetTree ssz hashes the BeaconState object.
func (st *BeaconState) GetTree() (*fastssz.Node, error) {
	return fastssz.ProofTree(st)
}
