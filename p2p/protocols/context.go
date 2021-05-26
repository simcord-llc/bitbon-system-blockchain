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

package protocols

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"

	"github.com/ethersphere/swarm/spancontext"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"
)

// msgWithContext is used to propagate marshalled context alongside message payloads
type msgWithContext struct {
	Context []byte
	Msg     []byte
}

func encodeWithContext(ctx context.Context, msg interface{}) (interface{}, int, error) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	tracer := opentracing.GlobalTracer()
	sctx := spancontext.FromContext(ctx)
	if sctx != nil {
		err := tracer.Inject(
			sctx,
			opentracing.Binary,
			writer)
		if err != nil {
			return nil, 0, err
		}
	}
	writer.Flush()
	msgBytes, err := rlp.EncodeToBytes(msg)
	if err != nil {
		return nil, 0, err
	}

	return &msgWithContext{
		Context: b.Bytes(),
		Msg:     msgBytes,
	}, len(msgBytes), nil
}

func decodeWithContext(msg p2p.Msg) (context.Context, []byte, error) {
	var wmsg msgWithContext
	err := msg.Decode(&wmsg)
	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()

	if len(wmsg.Context) == 0 {
		return ctx, wmsg.Msg, nil
	}

	tracer := opentracing.GlobalTracer()
	sctx, err := tracer.Extract(opentracing.Binary, bytes.NewReader(wmsg.Context))
	if err != nil {
		return nil, nil, err
	}
	ctx = spancontext.WithContext(ctx, sctx)
	return ctx, wmsg.Msg, nil
}

func encodeWithoutContext(ctx context.Context, msg interface{}) (interface{}, int, error) {
	return msg, 0, nil
}

func decodeWithoutContext(msg p2p.Msg) (context.Context, []byte, error) {
	b, err := ioutil.ReadAll(msg.Payload)
	if err != nil {
		return nil, nil, err
	}
	return context.Background(), b, nil
}
