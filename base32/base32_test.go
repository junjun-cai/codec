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
	testsHex = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("EHK6ISP0D5PI0PBECDNM8P9E"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("EHK6ISP0D5PI0OJ1EDIJ6CH0CLN66RR4CKN0===="),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("T2VPJPKOLVIBH076LIGI0OJ1EDIJ6CH0SUU9DPT0G4NUH9T3SUG83PLLHFKAV5F3G210===="),
		},
	}
	testsRawStd = []struct {
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
			encodText: []byte("ORUGS4ZANFZSAYTBONSTGMRAMVXGG33EMUXA"),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("5C7ZTZUYV7SLRAHGVSQSAYTBONSTGMRA466JNZ5AQEX6RJ5D46QIDZVVRPUK7FPDQCBA"),
		},
	}
	testsRawHex = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("EHK6ISP0D5PI0PBECDNM8P9E"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("EHK6ISP0D5PI0OJ1EDIJ6CH0CLN66RR4CKN0"),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("T2VPJPKOLVIBH076LIGI0OJ1EDIJ6CH0SUU9DPT0G4NUH9T3SUG83PLLHFKAV5F3G210"),
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

func TestBase32HexCodec_Encode(t *testing.T) {
	for _, tt := range testsHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexCodec.Encode(tt.plainText)
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

func TestBase32HexCodec_Decode(t *testing.T) {
	for _, tt := range testsHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexCodec.Decode(tt.encodText)
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

func TestBase32RawStdCodec_Encode(t *testing.T) {
	for _, tt := range testsRawStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawStdCodec.Encode(tt.plainText)
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

func TestBase32RawStdCodec_Decode(t *testing.T) {
	for _, tt := range testsRawStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawStdCodec.Decode(tt.encodText)
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

func TestBase32RawHexCodec_Encode(t *testing.T) {
	for _, tt := range testsRawHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawHexCodec.Encode(tt.plainText)
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

func TestBase32RawHexCodec_Decode(t *testing.T) {
	for _, tt := range testsRawHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawHexCodec.Decode(tt.encodText)
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
