// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"encoding/binary"
	"fmt"

	"github.com/panjf2000/gnet"
)

// IntermediateCodec
// https://core.telegram.org/mtproto#tcp-transport
//
// In case 4-byte data alignment is needed,
// an intermediate version of the original protocol may be used:
// if the client sends 0xeeeeeeee as the first int (four bytes),
// then packet length is encoded always by four bytes as in the original version,
// but the sequence number and CRC32 are omitted,
// thus decreasing total packet size by 8 bytes.
//
type IntermediateCodec struct {
	*AesCTR128Crypto
	state     int
	packetLen uint32
}

func newMTProtoIntermediateCodec(crypto *AesCTR128Crypto) *IntermediateCodec {
	return &IntermediateCodec{
		AesCTR128Crypto: crypto,
		state:           WAIT_PACKET_LENGTH,
	}
}

// Encode encodes frames upon server responses into TCP stream.
func (c *IntermediateCodec) Encode(conn gnet.Conn, msg interface{}) ([]byte, error) {
	rawMsg, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		return nil, err
	}

	sb := make([]byte, 4)
	binary.LittleEndian.PutUint32(sb, uint32(len(rawMsg.Payload)))

	b := append(sb, rawMsg.Payload...)
	return c.Encrypt(b), nil
}

// Decode decodes frames from TCP stream via specific implementation.
func (c *IntermediateCodec) Decode(conn gnet.Conn) (interface{}, error) {
	var (
		buf []byte
		n   int
		in  innerBuffer
		err error
	)

	in = conn.Read()

	if c.state == WAIT_PACKET_LENGTH {
		if buf, err = in.readN(4); err != nil {
			return nil, errUnexpectedEOF
		}
		conn.ShiftN(1)
		buf = c.Decrypt(buf)
		c.packetLen = binary.LittleEndian.Uint32(buf)
		c.state = WAIT_PACKET
	}

	needAck := c.packetLen>>31 == 1
	_ = needAck
	n = int(c.packetLen & 0xffffff)
	if n > MAX_MTPRORO_FRAME_SIZE {
		// TODO(@benqi): close conn
		return nil, fmt.Errorf("too large data(%d)", n)
	}

	if buf, err = in.readN(n); err != nil {
		return nil, errUnexpectedEOF
	}
	buf = c.Decrypt(buf)
	conn.ShiftN(n)
	c.state = WAIT_PACKET_LENGTH

	return NewMTPRawMessage(false,
		int64(binary.LittleEndian.Uint64(buf)),
		0,
		buf), nil
}

// Clone ...
func (c *IntermediateCodec) Clone() gnet.ICodec {
	return new(IntermediateCodec)
}

// Release ...
func (c *IntermediateCodec) Release() {
}
