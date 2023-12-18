// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/23 23:45:19                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base32                                                                                                      *
// * File: base32_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base32

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
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("ORUGS4ZANFZSAZLOMNXWIZJO"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("ORUGS4ZANFZSAYTBONSTGMRAMVXGG33EMUXA===="),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("5C7ZTZUYV7SLRAHGVSQSAYTBONSTGMRA466JNZ5AQEX6RJ5D46QIDZVVRPUK7FPDQCBA===="),
		},
	}
)

func TestBase32StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base32.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base32.Encode() failed!")
			} else {
				t.Log("base32.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase32StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base32.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base32.Decode() failed!")
			} else {
				t.Log("base32.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
