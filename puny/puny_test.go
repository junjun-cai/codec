package puny

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
			name:      "puny-1",
			plainText: []byte("this is encode."),
			encodText: []byte("this is encode."),
		},
		{
			name:      "puny-2",
			plainText: []byte("this is puny encode."),
			encodText: []byte("this is puny encode."),
		},
		{
			name:      "puny-3",
			plainText: []byte("这是一次 puny 编码/解码测试。"),
			encodText: []byte("xn-- puny /-dp3fs18v3s8b6mlfkgo42bba603ycq3b75gipt"),
		},
	}
)

func Test_punyCodec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("punyCodec.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Log("puny.Encode() failed!")
			} else {
				t.Log("puny.Encode() success!")
			}
			t.Logf("gots: %v", string(got))
			t.Logf("want: %v", string(tt.encodText))
		})
	}
}
