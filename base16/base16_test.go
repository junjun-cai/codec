// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/10 22:36:52                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base16                                                                                                      *
// * File: base16_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base16

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
			name:      "base16-1",
			plainText: []byte("this is encode."),
			encodText: []byte("7468697320697320656e636f64652e"),
		},
		{
			name:      "base16-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("7468697320697320626173653820656e636f64652e"),
		},
		{
			name:      "base16-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("e8bf99e698afe4b880e6aca120626173653820e7bc96e7a0812fe8a7a3e7a081e6b58be8af95e38082"),
		},
	}
	codecr, _ = NewCodec("ABCDEFGHIJKLMNOP")
	tests1    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base8-1",
			plainText: []byte("this is encode."),
			encodText: []byte("HEGIGJHDCAGJHDCAGFGOGDGPGEGFCO"),
		},
		{
			name:      "base8-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("HEGIGJHDCAGJHDCAGCGBHDGFDICAGFGOGDGPGEGFCO"),
		},
		{
			name:      "base8-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("OILPJJOGJIKPOELIIAOGKMKBCAGCGBHDGFDICAOHLMJGOHKAIBCPOIKHKDOHKAIBOGLFILOIKPJFODIAIC"),
		},
	}
)

func TestBase16Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base16.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base16.Encode() failed!")
			} else {
				t.Log("base16.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase16Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base16.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base16.Decode() failed!")
			} else {
				t.Log("base16.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase16CusCodec_Encode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base16.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base16.Encode() failed!")
			} else {
				t.Log("base16.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase16CusCodec_Decode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base16.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base16.Decode() failed!")
			} else {
				t.Log("base16.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
