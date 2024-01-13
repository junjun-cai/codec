// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/13 21:36:31                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base58                                                                                                      *
// * File: base58_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base58

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
			name:      "base58-1",
			plainText: []byte("this is encode."),
			encodText: []byte("4FteP9t8dCkdVyrv9hG5B"),
		},
		{
			name:      "base58-2",
			plainText: []byte("this is base58 encode."),
			encodText: []byte("YbUUaGb5ufH9tFog9KKdeqaP2fqFjs"),
		},
		{
			name:      "base58-3",
			plainText: []byte("这是一次 base58 编码/解码测试。"),
			encodText: []byte("4tJJ9J6xaKpNaa4bBd7EgQxa3v6euKKEFk7vTnaiV1GeA7mjC4siqaNBLu"),
		},
	}
)

func TestBase58StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base58.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base58.Encode() failed!")
			} else {
				t.Log("base58.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase58StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base58.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base58.Decode() failed!")
			} else {
				t.Log("base58.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
