// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

package errors

type Error interface {
	error
	ErrorCode() int
}

type internalError struct {
	code int
	orig error
}

func (e *internalError) ErrorCode() int {
	return e.code
}

func (e *internalError) Error() string {
	return e.orig.Error()
}

func New(err error, code int) Error {
	return &internalError{
		code: code,
		orig: err,
	}
}

func NewInternalError(err error) Error {
	return New(err, CodeInternalError)
}

func NewInvalidParamsError(err error) Error {
	return New(err, CodeInvalidParamsError)
}

func NewInvalidRequestError(err error) Error {
	return New(err, CodeInvalidRequestError)
}

func NewNotFoundError(err error) Error {
	return New(err, CodeNotFoundError)
}

func NewContractCallError(err error) Error {
	return New(err, CodeContractCallError)
}

func NewContractWatchError(err error) Error {
	return New(err, CodeContractWatchError)
}

func NewSafeTransferExpiredError(err error) Error {
	return New(err, CodeSafeTransferExpiredError)
}

func NewSafeTransferWrongProtectionCodeError(err error) Error {
	return New(err, CodeSafeTransferWrongProtectionCodeError)
}

func NewSafeTransferWrongProtectionCodeLimitError(err error) Error {
	return New(err, CodeSafeTransferWrongProtectionCodeLimitError)
}

func NewSafeTransferUndefinedError(err error) Error {
	return New(err, CodeSafeTransferUndefinedError)
}

func NewSafeTransferUnknownReasonError(err error) Error {
	return New(err, CodeSafeTransferUnknownReasonError)
}

func NewEmitEtherError(err error) Error {
	return New(err, CodeEmitEtherError)
}

func NewBalanceTooLowError(err error) Error {
	return New(err, CodeBalanceTooLowError)
}

func HasLockedBalanceError(err error) Error {
	return New(err, CodeHasLockedBalanceError)
}

func NewTransferExistsError(err error) Error {
	return New(err, CodeTransferExistsError)
}

func NewTransferIsNotExistError(err error) Error {
	return New(err, CodeTransferIsNotExistError)
}

func NewUnsuccessfullTxError(err error) Error {
	return New(err, CodeUnsuccessfulTxError)
}

func NewDeleteLastAssetboxError(err error) Error {
	return New(err, CodeDeleteLastAssetboxError)
}
