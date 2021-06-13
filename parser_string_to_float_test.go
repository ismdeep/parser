package parser

import "testing"

func TestToFloat32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "TestToFloat32-001",
			args: args{
				s: "0.01",
			},
			want:    0.01,
			wantErr: false,
		},
		{
			name: "TestToFloat32-002",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "TestToFloat64-001",
			args: args{
				s: "0.01",
			},
			want:    0.01,
			wantErr: false,
		},
		{
			name: "TestToFloat64-002",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
