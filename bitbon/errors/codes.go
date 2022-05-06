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

const (
	CodeInternalError       = -1
	CodeInvalidParamsError  = -2
	CodeInvalidRequestError = -3
	CodeNotFoundError       = -4

	CodeContractCallError  = 1
	CodeContractWatchError = 2

	CodeSafeTransferExpiredError                  = 3
	CodeSafeTransferWrongProtectionCodeError      = 4
	CodeSafeTransferWrongProtectionCodeLimitError = 5
	CodeSafeTransferUndefinedError                = 6
	CodeSafeTransferUnknownReasonError            = 7

	CodeEmitEtherError = 8

	CodeBalanceTooLowError      = 9
	CodeTransferExistsError     = 10
	CodeTransferIsNotExistError = 14

	CodeUnsuccessfulTxError     = 18
	CodeDeleteLastAssetboxError = 19
	CodeHasLockedBalanceError   = 20
)
