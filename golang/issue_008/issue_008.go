package issue008

import (
	"reflect"
	"strconv"

	"github.com/dustin/go-humanize"
)

type Number interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

//Go泛型数值转字符串方法优化

func NumberToString[T Number](value *T) string {
	if value == nil {
		return ""
	}

	v := reflect.ValueOf(*value)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		return humanize.FormatFloat("#,###.##", f)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		return strconv.FormatInt(i, 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		u := v.Uint()
		return strconv.FormatUint(u, 10)
	default:
		return ""
	}
}
