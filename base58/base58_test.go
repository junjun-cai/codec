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
	encoder     = "9ABCDEFGHJ12KLMNPQRST34UVWXYZab56cdefghijkmnopqrstu78vwxyz"
	cusCodec, _ = NewCodec(encoder)
	testsCus    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base58-1",
			plainText: []byte("this is encode."),
			encodText: []byte("CM7g4H7Gf2nfZytvHjND1"),
		},
		{
			name:      "base58-2",
			plainText: []byte("this is base58 encode."),
			encodText: []byte("5dYYcNdD8hPH7MqiHRRfgsc4AhsMmu"),
		},
		{
			name:      "base58-3",
			plainText: []byte("这是一次 base58 编码/解码测试。"),
			encodText: []byte("C7QQHQExcRr3ccCd1fFLiUxcBvEg8RRLMnFvXpckZ9NgJFom2Cuksc31S8"),
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

func TestBase58CusCodec_Encode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Encode(tt.plainText)
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

func TestBase58SCusCodec_Decode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Decode(tt.encodText)
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
