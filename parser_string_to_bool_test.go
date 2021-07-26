package parser

import "testing"

func TestToBool(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "",
			args: args{
				s: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s: "false",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s: "invalid-value",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}
