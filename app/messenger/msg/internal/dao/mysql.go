// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dao/mysql_dao"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.MessagesDAO
	*mysql_dao.ChatParticipantsDAO
	*mysql_dao.HashTagsDAO
	*mysql_dao.DialogsDAO
	*sqlx.CommonDAO
}

func NewMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                  db,
		MessagesDAO:         mysql_dao.NewMessagesDAO(db),
		ChatParticipantsDAO: mysql_dao.NewChatParticipantsDAO(db),
		HashTagsDAO:         mysql_dao.NewHashTagsDAO(db),
		DialogsDAO:          mysql_dao.NewDialogsDAO(db),
		CommonDAO:           sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
