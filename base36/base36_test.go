// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/26 22:25:07                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base36                                                                                                      *
// * File: base36_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base36

import (
	"reflect"
	"testing"
)

var (
	testsStd = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base36-1",
			plainText: []byte("this is encode."),
			encodText: []byte("YVZYQ7KSD1CTDKDSRV0B0VI"),
		},
		{
			name:      "base36-2",
			plainText: []byte("this is base36 encode."),
			encodText: []byte("J3LX98DQYZ75JFNQHIU8LSYA31C7BX07XQ"),
		},
		{
			name:      "base36-3",
			plainText: []byte("这是一次 base36 编码/解码测试。"),
			encodText: []byte("VQ05TKNUB8X1YCBUKD6HCWV8LBWG05D43RKVRCOGBG78DP7N757OD0DBNR3GXEU0Y"),
		},
	}
	encoder     = "~`!@#$%^&*()_-+[{]}|;:,<.>/?12345678"
	cusCodec, _ = NewCodec(encoder)
	testsCus    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base36-1",
			plainText: []byte("this is encode."),
			encodText: []byte("7487/^;1-`_2-;-1?4~)~4}"),
		},
		{
			name:      "base36-2",
			plainText: []byte("this is base36 encode."),
			encodText: []byte("|@:6*&-/78^$|[</]}3&:17(@`_^)6~^6/"),
		},
		{
			name:      "base36-3",
			plainText: []byte("这是一次 base36 编码/解码测试。"),
			encodText: []byte("4/~$2;<3)&6`7_)3;-%]_54&:)5{~$-#@?;4?_.{){^&->^<^$^.-~-)<?@{6+3~7"),
		},
	}
)

func TestBase36StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base36.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base36.Encode() failed!")
			} else {
				t.Log("base36.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase36StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base36.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base36.Decode() failed!")
			} else {
				t.Log("base36.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase36CusCodec_Encode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base36.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base36.Encode() failed!")
			} else {
				t.Log("base36.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase36SCusCodec_Decode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base36.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base36.Decode() failed!")
			} else {
				t.Log("base36.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
