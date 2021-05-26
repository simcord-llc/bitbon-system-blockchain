// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
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

package common

import (
	"bytes"
	"sort"

	"golang.org/x/crypto/sha3"
)

type SortAddresses []Address

func (s SortAddresses) Len() int      { return len(s) }
func (s SortAddresses) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortAddresses) Less(i, j int) bool {
	return bytes.Compare(s[i][:], s[j][:]) <= 0
}

func Keccak256Addresses(data ...Address) (h Hash) {
	sorted := SortAddresses(data)
	sort.Sort(sorted)

	d := sha3.NewLegacyKeccak256()
	for _, b := range sorted {
		d.Write(b[:]) // nolint:errcheck
	}
	d.Sum(h[:0])
	return h
}

type SortHashes []Hash

func (s SortHashes) Len() int      { return len(s) }
func (s SortHashes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortHashes) Less(i, j int) bool {
	return bytes.Compare(s[i][:], s[j][:]) <= 0
}

func Keccak256Hashes(data ...Hash) (h Hash) {
	sorted := SortHashes(data)
	sort.Sort(sorted)

	d := sha3.NewLegacyKeccak256()
	for _, b := range sorted {
		d.Write(b[:]) // nolint:errcheck
	}
	d.Sum(h[:0])
	return h
}
