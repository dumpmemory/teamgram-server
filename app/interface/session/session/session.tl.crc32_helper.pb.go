/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package session

const (
	CRC32_UNKNOWN                        TLConstructor = 0
	CRC32_sessionClientEvent             TLConstructor = -548007522  // 0xdf56119e
	CRC32_sessionClientData              TLConstructor = -870242788  // 0xcc21261c
	CRC32_httpSessionData                TLConstructor = -606579889  // 0xdbd8534f
	CRC32_session_queryAuthKey           TLConstructor = 1798174801  // 0x6b2df851
	CRC32_session_setAuthKey             TLConstructor = 487672075   // 0x1d11490b
	CRC32_session_createSession          TLConstructor = 1091351053  // 0x410cb20d
	CRC32_session_sendDataToSession      TLConstructor = -2023019028 // 0x876b2dec
	CRC32_session_sendHttpDataToSession  TLConstructor = -1142152274 // 0xbbec23ae
	CRC32_session_closeSession           TLConstructor = 393200211   // 0x176fc253
	CRC32_session_pushUpdatesData        TLConstructor = -1050612110 // 0xc160ee72
	CRC32_session_pushSessionUpdatesData TLConstructor = -434243286  // 0xe61df92a
	CRC32_session_pushRpcResultData      TLConstructor = 1627318120  // 0x60fee768
)