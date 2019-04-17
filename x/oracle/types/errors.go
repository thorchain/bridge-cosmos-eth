package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Local code type
type CodeType = sdk.CodeType

//Exported code type numbers
const (
	DefaultCodespace sdk.CodespaceType = "oracle"

	CodeInvalidNonce           CodeType = 1
	CodeProphecyNotFound       CodeType = 2
	CodeMinimumPowerTooLow     CodeType = 3
	CodeInvalidIdentifier      CodeType = 4
	CodeNoClaims               CodeType = 5
	CodeInvalidEthereumNonce   CodeType = 6
	CodeInvalidEthereumAddress CodeType = 7
)

func ErrProphecyNotFound(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeProphecyNotFound, "prophecy with given id not found")
}

func ErrMinimumPowerTooLow(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeMinimumPowerTooLow, "minimum number for validator staking power must be greater than 1")
}

func ErrNoClaims(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoClaims, "cannot create prophecy without initial claim")
}

func ErrInvalidIdentifier(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidNonce, "invalid identifier provided, must be a nonempty string")
}

//Ethereum-specific stuff

func ErrInvalidEthereumNonce(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidEthereumNonce, "invalid ethereum nonce provided, must be >= 0")
}

func ErrInvalidEthereumAddress(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidEthereumAddress, "invalid ethereum address provided, must be a valid hex-encoded Ethereum address")
}
