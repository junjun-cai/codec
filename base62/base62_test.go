// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/18 22:41:30                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base62                                                                                                      *
// * File: base62_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base62

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
			name:      "base62-1",
			plainText: []byte("this is encode."),
			encodText: []byte("rCJBhfvjUf12Ocy2NNtO"),
		},
		{
			name:      "base62-2",
			plainText: []byte("this is base62 encode."),
			encodText: []byte("4ZAp5fLvpSOmQjtik54pZatNTlTiik"),
		},
		{
			name:      "base62-3",
			plainText: []byte("这是一次 base62 编码/解码测试。"),
			encodText: []byte("5NZ7edvjXxa5sRWzZTZrJDqdcQo7zNxNLtzptPTTpMWAjhYgtER78b7kA"),
		},
	}
	encoder     = "@#$%&*(){}ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	cusCodec, _ = NewCodec(encoder)
	testsCus    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base62-1",
			plainText: []byte("this is encode."),
			encodText: []byte("rCJBhfvjUf#$Ocy$NNtO"),
		},
		{
			name:      "base62-2",
			plainText: []byte("this is base62 encode."),
			encodText: []byte("&ZAp*fLvpSOmQjtik*&pZatNTlTiik"),
		},
		{
			name:      "base62-3",
			plainText: []byte("这是一次 base62 编码/解码测试。"),
			encodText: []byte("*NZ)edvjXxa*sRWzZTZrJDqdcQo)zNxNLtzptPTTpMWAjhYgtER){b)kA"),
		},
	}
)

func TestBase62StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base62.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base62.Encode() failed!")
			} else {
				t.Log("base62.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase62StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base62.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base62.Decode() failed!")
			} else {
				t.Log("base62.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase62CusCodec_Encode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base62.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base62.Encode() failed!")
			} else {
				t.Log("base62.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase62CusCodec_Decode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base62.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base62.Decode() failed!")
			} else {
				t.Log("base62.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
