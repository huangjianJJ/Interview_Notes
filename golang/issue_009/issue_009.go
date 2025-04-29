package issue009

//可以通过泛型结合反射动态处理不同的 sql.Null* 类型
import (
	"reflect"
)

func SQLValue[T any, U any](value T) *U {
	rv := reflect.ValueOf(value)

	// 检查是否为结构体且包含 Valid 字段
	if rv.Kind() != reflect.Struct {
		return nil
	}
	validField := rv.FieldByName("Valid")
	if !validField.IsValid() || validField.Kind() != reflect.Bool {
		return nil
	}
	if !validField.Bool() {
		return nil
	}

	// 查找数值字段（第一个非 Valid 的字段）
	var numVal reflect.Value
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Name == "Valid" {
			continue
		}
		numVal = rv.Field(i)
		break
	}

	// 类型转换检查
	if !numVal.IsValid() || !numVal.Type().ConvertibleTo(reflect.TypeOf((*U)(nil)).Elem()) {
		return nil
	}

	// 执行类型转换并返回指针
	converted := numVal.Convert(reflect.TypeOf((*U)(nil)).Elem())
	result := converted.Interface().(U)
	return &result
}
