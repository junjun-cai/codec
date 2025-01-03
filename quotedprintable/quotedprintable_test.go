package quotedprintable

import (
	"reflect"
	"testing"
)

func Test_quotedprintableCodec_Encode(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		q       *quotedprintableCodec
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Encode(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("quotedprintableCodec.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quotedprintableCodec.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
