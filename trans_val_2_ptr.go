package cast

func NewBool(b bool) (v *bool) {
	v = new(bool)
	*v = b
	return
}

func NewByte(b byte) (v *byte) {
	v = new(byte)
	*v = b
	return
}

func NewInt8(i8 int8) (v *int8) {
	v = new(int8)
	*v = i8
	return
}

func NewUint8(u8 uint8) (v *uint8) {
	v = new(uint8)
	*v = u8
	return
}

func NewInt16(i16 int16) (v *int16) {
	v = new(int16)
	*v = i16
	return
}

func NewUint16(u16 uint16) (v *uint16) {
	v = new(uint16)
	*v = u16
	return
}

func NewInt32(i32 int32) (v *int32) {
	v = new(int32)
	*v = i32
	return
}

func NewUint32(u32 uint32) (v *uint32) {
	v = new(uint32)
	*v = u32
	return
}

func NewInt64(i64 int64) (v *int64) {
	v = new(int64)
	*v = i64
	return
}

func NewUint64(u64 uint64) (v *uint64) {
	v = new(uint64)
	*v = u64
	return
}

func NewFloat32(f32 float32) (v *float32) {
	v = new(float32)
	*v = f32
	return
}

func NewFloat64(f64 float64) (v *float64) {
	v = new(float64)
	*v = f64
	return
}

func NewString(str string) (v *string) {
	v = new(string)
	*v = str
	return
}
