// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package datatypes

import (
	"testing"

	"github.com/gopcua/opcua/utils/codectest"
)

func TestExtensionObject(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "anonymous-user-identity-token",
			Struct: NewExtensionObject(
				0x01, NewAnonymousIdentityToken("anonymous"),
			),
			Bytes: []byte{
				// TypeID
				0x01, 0x00, 0x41, 0x01,
				// EncodingMask
				0x01,
				// Length
				0x0d, 0x00, 0x00, 0x00,
				// AnonymousIdentityToken
				0x09, 0x00, 0x00, 0x00, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
			},
		},
	}
	codectest.Run(t, cases)
}
