package cast

import (
	"math"
	"reflect"
	"testing"
)

func TestNewBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name  string
		args  args
		wantV bool
	}{
		{name: "NewBoolFalse", args: args{b: false}, wantV: false},
		{name: "NewBoolTrue", args: args{b: true}, wantV: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewBool(tt.args.b); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewBool() = %v, want %v", *gotV, tt.wantV)
			}
		})
	}
}

func TestNewByte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name  string
		args  args
		wantV byte
	}{
		{name: "NewByte_1", args: args{b: 1}, wantV: 1},
		{name: "NewByte_2", args: args{b: 2}, wantV: 2},
		{name: "NewByte_3", args: args{b: 3}, wantV: 3},
		{name: "NewByte_127", args: args{b: 127}, wantV: 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewByte(tt.args.b); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewByte() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewFloat32(t *testing.T) {
	type args struct {
		f32 float32
	}
	tests := []struct {
		name  string
		args  args
		wantV float32
	}{
		{name: "NewFloat32_0.0", args: args{f32: 0.0}, wantV: 0.0},
		{name: "NewFloat32_3.14", args: args{f32: 3.14}, wantV: 3.14},
		{name: "NewFloat32_0.00001", args: args{f32: 0.00001}, wantV: 0.00001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewFloat32(tt.args.f32); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewFloat32() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewFloat64(t *testing.T) {
	type args struct {
		f64 float64
	}
	tests := []struct {
		name  string
		args  args
		wantV float64
	}{
		{name: "NewFloat64_0.0", args: args{f64: 0.00}, wantV: 0.00},
		{name: "NewFloat64_3.14", args: args{f64: 3.14}, wantV: 3.14},
		{name: "NewFloat64_0.00001", args: args{f64: 0.00001}, wantV: 0.00001},
		{name: "NewFloat64_0.0000001", args: args{f64: 0.0000001}, wantV: 0.0000001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewFloat64(tt.args.f64); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewFloat64() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewInt16(t *testing.T) {
	type args struct {
		i16 int16
	}
	tests := []struct {
		name  string
		args  args
		wantV int16
	}{
		{name: "New16_0", args: args{i16: 0}, wantV: 0},
		{name: "New16_127", args: args{i16: 127}, wantV: 127},
		{name: "New16_-128", args: args{i16: -128}, wantV: 0 - 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewInt16(tt.args.i16); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewInt16() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewInt32(t *testing.T) {
	type args struct {
		i32 int32
	}
	tests := []struct {
		name  string
		args  args
		wantV int32
	}{
		{name: "NewInt32_0", args: args{i32: 0}, wantV: 0},
		{name: "NewInt32_2147483647", args: args{i32: math.MaxInt32}, wantV: math.MaxInt32},
		{name: "NewInt32_-2147483648", args: args{i32: math.MinInt32}, wantV: math.MinInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewInt32(tt.args.i32); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewInt32() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewInt64(t *testing.T) {
	type args struct {
		i64 int64
	}
	tests := []struct {
		name  string
		args  args
		wantV int64
	}{
		{name: "NewInt64_0", args: args{i64: 0}, wantV: 0},
		{name: "NewInt64_9223372036854775807", args: args{i64: math.MaxInt64}, wantV: math.MaxInt64},
		{name: "NewInt64_-9223372036854775808", args: args{i64: math.MinInt64}, wantV: math.MinInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewInt64(tt.args.i64); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewInt64() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewInt8(t *testing.T) {
	type args struct {
		i8 int8
	}
	tests := []struct {
		name  string
		args  args
		wantV int8
	}{
		{name: "NewInt8_0", args: args{i8: 0}, wantV: 0},
		{name: "NewInt8_127", args: args{i8: math.MaxInt8}, wantV: math.MaxInt8},
		{name: "NewInt8_-128", args: args{i8: math.MinInt8}, wantV: math.MinInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewInt8(tt.args.i8); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewInt8() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{name: "NewString", args: args{str: "hello world"}, wantV: "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewString(tt.args.str); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewUint16(t *testing.T) {
	type args struct {
		u16 uint16
	}
	tests := []struct {
		name  string
		args  args
		wantV uint16
	}{
		{name: "NewUint16_0", args: args{u16: 0}, wantV: 0},
		{name: "NewUint16_32767", args: args{u16: math.MaxInt16}, wantV: math.MaxInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewUint16(tt.args.u16); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewUint16() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewUint32(t *testing.T) {
	type args struct {
		u32 uint32
	}
	tests := []struct {
		name  string
		args  args
		wantV uint32
	}{
		{name: "NewUint32_0", args: args{u32: 0}, wantV: 0},
		{name: "NewUint32_4294967295", args: args{u32: math.MaxUint32}, wantV: math.MaxUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewUint32(tt.args.u32); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewUint32() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewUint64(t *testing.T) {
	type args struct {
		u64 uint64
	}
	tests := []struct {
		name  string
		args  args
		wantV uint64
	}{
		{name: "NewUint64_0", args: args{u64: 0}, wantV: 0},
		{name: "NewUint64_9223372036854775808L", args: args{u64: math.MaxUint64}, wantV: math.MaxUint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewUint64(tt.args.u64); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewUint64() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestNewUint8(t *testing.T) {
	type args struct {
		u8 uint8
	}
	tests := []struct {
		name  string
		args  args
		wantV uint8
	}{
		{name: "NewUint8_0", args: args{u8: 0}, wantV: 0},
		{name: "NewUint8_255", args: args{u8: math.MaxUint8}, wantV: math.MaxUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewUint8(tt.args.u8); !reflect.DeepEqual(*gotV, tt.wantV) {
				t.Errorf("NewUint8() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
