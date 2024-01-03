// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/02 22:52:06                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base45                                                                                                      *
// * File: base45_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base45

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
			name:      "base45-1",
			plainText: []byte("this is encode."),
			encodText: []byte("AWE+EDH44.OE1$CUPC6VC11"),
		},
		{
			name:      "base45-2",
			plainText: []byte("this is base45 encode."),
			encodText: []byte("AWE+EDH44.OEUJCLQE0R6D44:.DV3ERZC"),
		},
		{
			name:      "base45-3",
			plainText: []byte("这是一次 base45 编码/解码测试。"),
			encodText: []byte("3JTNKJRDJ7-SDDG3$LA44HECXZCAW6EDTL3J4DKO26U8LVCT:IGZ.MWITV.I3BG"),
		},
	}
	encoder     = "&$%*+-./:ABCDEFGHIJKL6789MNO34PQRSTUVWXYZ0125"
	cusCodec, _ = NewCodec(encoder)
	testsCus    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base45-1",
			plainText: []byte("this is encode."),
			encodText: []byte("BRFZFEI++19F$WDPMD.QD$$"),
		},
		{
			name:      "base45-2",
			plainText: []byte("this is base45 encode."),
			encodText: []byte("BRFZFEI++19FPKD6NF&O.E++51EQ*FOUD"),
		},
		{
			name:      "base45-3",
			plainText: []byte("这是一次 base45 编码/解码测试。"),
			encodText: []byte("*K48LKOEK/03EEH*W6B++IFDSUDBR.FE46*K+EL9%.P:6QD45JHU17RJ4Q1J*CH"),
		},
	}
)

func TestBase45StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base45.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base45.Encode() failed!")
			} else {
				t.Log("base45.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase45StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base45.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base45.Decode() failed!")
			} else {
				t.Log("base45.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase45CusCodec_Encode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base45.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base45.Encode() failed!")
			} else {
				t.Log("base45.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase45SCusCodec_Decode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base45.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base45.Decode() failed!")
			} else {
				t.Log("base45.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}
