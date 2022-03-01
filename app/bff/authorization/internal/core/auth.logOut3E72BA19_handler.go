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

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// AuthLogOut3E72BA19
// auth.logOut#3e72ba19 = auth.LoggedOut;
func (c *AuthorizationCore) AuthLogOut3E72BA19(in *mtproto.TLAuthLogOut3E72BA19) (*mtproto.Auth_LoggedOut, error) {
	// unbind auth_key and user_id
	_, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionUnbindAuthKeyUser(c.ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: c.MD.AuthId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.logOut - error: %v", err)
	} else {
		k, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionQueryAuthKey(c.ctx, &authsession.TLAuthsessionQueryAuthKey{
			AuthKeyId: c.MD.AuthId,
		})
		if err != nil {
			c.Logger.Errorf("auth.logOut - error: %v", err)
		} else {
			_ = k
			// TODO: dispatch enterprise server
			// c.svcCtx.Dao.PushClient.PushDeleteDevice(c.ctx, &push.TLPushDeleteDevice{
			//	UserId:     c.MD.UserId,
			//	AuthKeyIds: []int64{k.PermAuthKeyId, k.TempAuthKeyId, k.MediaTempAuthKeyId},
			// })
		}
	}

	return mtproto.MakeTLAuthLoggedOut(&mtproto.Auth_LoggedOut{
		FutureAuthToken: nil,
	}).To_Auth_LoggedOut(), nil
}
