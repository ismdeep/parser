package parser

import "testing"

func TestToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "TestToInt-001",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "TestToInt-002",
			args: args{
				s: "-100",
			},
			want:    -100,
			wantErr: false,
		},
		{
			name: "TestToInt-003",
			args: args{
				s: "214748364700",
			},
			want:    214748364700,
			wantErr: false,
		},
		{
			name: "TestToInt-004",
			args: args{
				s: "-214748364700",
			},
			want:    -214748364700,
			wantErr: false,
		},
		{
			name: "TestToInt-005",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "TestToInt32-001",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "TestToInt32-002",
			args: args{
				s: "-100",
			},
			want:    -100,
			wantErr: false,
		},
		{
			name: "TestToInt32-003",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "TestToInt64-001",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "TestToInt64-002",
			args: args{
				s: "-100",
			},
			want:    -100,
			wantErr: false,
		},
		{
			name: "TestToInt64-003",
			args: args{
				s: "214748364700",
			},
			want:    214748364700,
			wantErr: false,
		},
		{
			name: "TestToInt64-004",
			args: args{
				s: "-214748364700",
			},
			want:    -214748364700,
			wantErr: false,
		},
		{
			name: "TestToInt64-005",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "TestToUint-001",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "TestToUint-002",
			args: args{
				s: "214748364700",
			},
			want:    214748364700,
			wantErr: false,
		},
		{
			name: "TestToUint-003",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "TestToUint32-001",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "TestToUint32-002",
			args: args{
				s: "3147483647",
			},
			want:    3147483647,
			wantErr: false,
		},
		{
			name: "TestToUint32-003",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "TestToUint64-001",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "TestToUint64-002",
			args: args{
				s: "214748364700",
			},
			want:    214748364700,
			wantErr: false,
		},
		{
			name: "TestToUint64-003",
			args: args{
				s: "invalid number",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
