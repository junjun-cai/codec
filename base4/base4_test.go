// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/21 22:34:18                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base4                                                                                                       *
// * File: base4_test.go                                                                                               *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base4

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
			name:      "base4-1",
			plainText: []byte("this is encode."),
			encodText: []byte("131012201221130302001221130302001211123212031233121012110232"),
		},
		{
			name:      "base4-2",
			plainText: []byte("this is base4 encode."),
			encodText: []byte("131012201221130302001221130302001202120113031211031002001211123212031233121012110232"),
		},
		{
			name:      "base4-3",
			plainText: []byte("这是一次 base4 编码/解码测试。"),
			encodText: []byte("3220233321213212212022333210232020003212223022010200120212011303121103100200321323302" +
				"1123213220020010233322022132203321322002001321223112023322022332111320320002002"),
		},
	}
)

func TestBase4Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base4.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base4.Encode() failed!")
			} else {
				t.Log("base4.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase4Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base4.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base4.Decode() failed!")
			} else {
				t.Log("base4.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
