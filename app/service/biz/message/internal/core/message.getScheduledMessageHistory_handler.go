/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetScheduledMessageHistory
// message.getScheduledMessageHistory user_id:long peer_type:int peer_id:long = Vector<MessageBox>;
func (c *MessageCore) MessageGetScheduledMessageHistory(in *message.TLMessageGetScheduledMessageHistory) (*message.Vector_MessageBox, error) {
	c.Logger.Errorf("message.getScheduledMessageHistory blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return &message.Vector_MessageBox{
		Datas: []*mtproto.MessageBox{},
	}, nil
}
