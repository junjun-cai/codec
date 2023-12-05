// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/22 23:47:56                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base8                                                                                                       *
// * File: base8_test.go                                                                                               *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base8

import (
	"github.com/caijunjun/codec/base"
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
			name:      "base8-1",
			plainText: []byte("this is encode."),
			encodText: []byte("3506415134620151346201453346155731062456"),
		},
		{
			name:      "base8-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("35064151346201513462014230271545160201453346155731062456"),
		},
		{
			name:      "base8-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("72137631715142577113420071526241100611413466247010163674455636404022775051721747501007465530575053712743401010=="),
		},
	}
	codecr, _ = NewCodec("ABCDEFGH", base.StdPadding)
	tests1    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base8-1",
			plainText: []byte("this is encode."),
			encodText: []byte("DFAGEBFBDEGCABFBDEGCABEFDDEGBFFHDBAGCEFG"),
		},
		{
			name:      "base8-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("DFAGEBFBDEGCABFBDEGCABECDACHBFEFBGACABEFDDEGBFFHDBAGCEFG"),
		},
		{
			name:      "base8-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("HCBDHGDBHBFBECFHHBBDECAAHBFCGCEBBAAGBBEBDEGGCEHABABGDGHEEFFGDGEAEACCHHFAFBHCBHEHFABAAHEGFFDAFHFAFDHBCHEDEABABA=="),
		},
	}
)

func TestBase8Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base8.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base8.Encode() failed!")
			} else {
				t.Log("base8.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase8Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base8.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base8.Decode() failed!")
			} else {
				t.Log("base8.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase8CusCodec_Encode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base8.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base8.Encode() failed!")
			} else {
				t.Log("base8.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase8CusCodec_Decode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base8.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base8.Decode() failed!")
			} else {
				t.Log("base8.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
