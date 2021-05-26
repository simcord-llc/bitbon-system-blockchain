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

package rpc

import (
	"context"
	"errors"
	"io"
	"net"
	"os"
	"time"
)

// DialStdIO creates a client on stdin/stdout.
func DialStdIO(ctx context.Context) (*Client, error) {
	return DialIO(ctx, os.Stdin, os.Stdout)
}

// DialIO creates a client which uses the given IO channels
func DialIO(ctx context.Context, in io.Reader, out io.Writer) (*Client, error) {
	return newClient(ctx, func(_ context.Context) (ServerCodec, error) {
		return NewCodec(stdioConn{
			in:  in,
			out: out,
		}), nil
	})
}

type stdioConn struct {
	in  io.Reader
	out io.Writer
}

func (c stdioConn) Read(b []byte) (n int, err error) {
	return c.in.Read(b)
}

func (c stdioConn) Write(b []byte) (n int, err error) {
	return c.out.Write(b)
}

func (c stdioConn) Close() error {
	return nil
}

func (c stdioConn) RemoteAddr() string {
	return "/dev/stdin"
}

func (c stdioConn) SetWriteDeadline(t time.Time) error {
	return &net.OpError{Op: "set", Net: "stdio", Source: nil, Addr: nil, Err: errors.New("deadline not supported")}
}
