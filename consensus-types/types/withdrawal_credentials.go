// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (c) 2025 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions
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

package types

import (
	"bytes"

	"github.com/berachain/beacon-kit/primitives/common"
)

const (
	// EthSecp256k1CredentialPrefix is the prefix for an Ethereum secp256k1.
	EthSecp256k1CredentialPrefix = byte(1)

	// numZeroBytesInEth1WithdrawalCredentials is the number of zero bytes in a valid eth1
	// withdrawal credentials.
	numZeroBytesInEth1WithdrawalCredentials = 11
)

// WithdrawalCredentials is a staking credential that is used to identify a validator.
type WithdrawalCredentials common.Bytes32

// NewCredentialsFromExecutionAddress creates a new valid Eth1 WithdrawalCredentials
// from an execution address.
func NewCredentialsFromExecutionAddress(address common.ExecutionAddress) WithdrawalCredentials {
	credentials := WithdrawalCredentials{}
	credentials[0] = EthSecp256k1CredentialPrefix
	copy(credentials[12:], address[:])
	return credentials
}

// IsValidEth1WithdrawalCredentials checks if the withdrawal credentials are valid
// eth1 withdrawal credentials as defined in the Ethereum 2.0 specification:
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/validator.md#eth1_address_withdrawal_prefix
func (wc WithdrawalCredentials) IsValidEth1WithdrawalCredentials() bool {
	zeroBytes := make([]byte, numZeroBytesInEth1WithdrawalCredentials)
	return wc[0] == EthSecp256k1CredentialPrefix && bytes.Equal(wc[1:12], zeroBytes)
}

// ToExecutionAddress converts the WithdrawalCredentials to an ExecutionAddress.
// Returns error if the withdrawal credentials are not valid.
func (wc WithdrawalCredentials) ToExecutionAddress() (common.ExecutionAddress, error) {
	if !wc.IsValidEth1WithdrawalCredentials() {
		return common.ExecutionAddress{}, ErrInvalidWithdrawalCredentials
	}
	return common.ExecutionAddress(wc[12:]), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Bytes32.
// TODO: Figure out how to not have to do this.
func (wc *WithdrawalCredentials) UnmarshalJSON(input []byte) error {
	return (*common.Bytes32)(wc).UnmarshalJSON(input)
}

// String returns the hex string representation of Bytes32.
// TODO: Figure out how to not have to do this.
func (wc WithdrawalCredentials) String() string {
	return common.Bytes32(wc).String()
}

// MarshalText implements the encoding.TextMarshaler interface for Bytes32.
// TODO: Figure out how to not have to do this.
func (wc WithdrawalCredentials) MarshalText() ([]byte, error) {
	return common.Bytes32(wc).MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Bytes32.
// TODO: Figure out how to not have to do this.
func (wc *WithdrawalCredentials) UnmarshalText(text []byte) error {
	return (*common.Bytes32)(wc).UnmarshalText(text)
}
