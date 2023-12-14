// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/14 22:21:12                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base24                                                                                                      *
// * File: base64_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base24

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
			name:      "base24-1",
			plainText: []byte("this is encode."),
			encodText: []byte("H3ETZ75C9FZAG44W82WYF4STR7CP"),
		},
		{
			name:      "base24-2",
			plainText: []byte("this is base24 encode."),
			encodText: []byte("H3ETZ75C9FZAG447EWP6WBGRP6GKGP6C78T4WBHPZZ"),
		},
		{
			name:      "base24-3",
			plainText: []byte("这是一次 base24 编码/解码测试。"),
			encodText: []byte("9HG6G86KG8AK7P5F6E29AC936KC24W3E6EP94EH5XY6CBGBGP6833BEP5GPSXWS9HEH5T25ER4W4Z"),
		},
	}
)

func TestBase24Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base24.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base24.Encode() failed!")
			} else {
				t.Log("base24.Encode() success!")
			}
			t.Logf(" got: %v", string(got))
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase24Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base24.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base24.Decode() failed!")
			} else {
				t.Log("base24.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
