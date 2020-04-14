package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types/errors"
)

// Ethermint error codes
const (
	// DefaultCodespace reserves a Codespace for Ethermint.
	DefaultCodespace string = "ethermint"

	CodeInvalidValue   uint32 = 1
	CodeInvalidChainID uint32 = 2
	CodeInvalidSender  uint32 = 3
	CodeVMExecution    uint32 = 4
	CodeInvalidNonce   uint32 = 5
)

// ErrInvalidValue returns a standardized SDK error resulting from an invalid value.
func ErrInvalidValue(msg string) *sdk.Error {
	return sdk.New(DefaultCodespace, CodeInvalidValue, msg)
}

// ErrInvalidChainID returns a standardized SDK error resulting from an invalid chain ID.
func ErrInvalidChainID(msg string) *sdk.Error {
	return sdk.New(DefaultCodespace, CodeInvalidChainID, msg)
}

// ErrInvalidSender returns a standardized SDK error resulting from an invalid transaction sender.
func ErrInvalidSender(msg string) *sdk.Error {
	return sdk.New(DefaultCodespace, CodeInvalidSender, msg)
}

// ErrVMExecution returns a standardized SDK error resulting from an error in EVM execution.
func ErrVMExecution(msg string) *sdk.Error {
	return sdk.New(DefaultCodespace, CodeVMExecution, msg)
}

// ErrVMExecution returns a standardized SDK error resulting from an error in EVM execution.
func ErrInvalidNonce(msg string) *sdk.Error {
	return sdk.New(DefaultCodespace, CodeInvalidNonce, msg)
}
