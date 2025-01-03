// *********************************************************************************************************************

// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***

// *********************************************************************************************************************

// * Auth: ColeCai                                                                                                     *

// * Date: 2023/11/29 21:53:50                                                                                         *

// * Proj: codec                                                                                                       *

// * Pack: base100                                                                                                     *

// * File: base100.go                                                                                                  *

// *-------------------------------------------------------------------------------------------------------------------*

// * Overviews:                                                                                                        *

// *-------------------------------------------------------------------------------------------------------------------*

// * Functions:                                                                                                        *

// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base100

import (
	"reflect"
	"testing"
)

var (
	tests = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base100-1",
			plainText: []byte("this is encode."),
			encodText: []byte("👫👟👠👪🐗👠👪🐗👜👥👚👦👛👜🐥"),
		},
		{
			name:      "base100-2",
			plainText: []byte("this is base100 encode."),
			encodText: []byte("👫👟👠👪🐗👠👪🐗👙👘👪👜🐨🐧🐧🐗👜👥👚👦👛👜🐥"),
		},
		{
			name:      "base100-3",
			plainText: []byte("这是一次 base100 编码/解码测试。"),
			encodText: []byte("📟💶💐📝💏💦📛💯👷📝💣💘🐗👙👘👪👜🐨🐧🐧🐗📞💳💍📞💗👸🐦📟💞💚📞💗👸📝💬💂📟💦💌📚👷👹"),
		},
	}
)

func Test_base100Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base100Codec.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Log("base100.Encode() failed!")
			} else {
				t.Log("base100.Encode() success!")
			}
			t.Logf("got:%v", string(got))
			t.Logf("want: %v", string(tt.encodText))
		})
	}
}
