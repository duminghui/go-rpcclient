// Package cmdjson provides ...
package cmdjson

func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

func Uint(v uint) *uint {
	p := new(uint)
	*p = v
	return p
}

func Int32(v int32) *int32 {
	p := new(int32)
	*p = v
	return p
}

func Uint32(v uint32) *uint32 {
	p := new(uint32)
	*p = v
	return p
}

func Int64(v int64) *int64 {
	p := new(int64)
	*p = v
	return p
}

func Uint64(v uint64) *uint64 {
	p := new(uint64)
	*p = v
	return p
}

func Float64(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

func String(v string) *string {
	p := new(string)
	*p = v
	return p
}
