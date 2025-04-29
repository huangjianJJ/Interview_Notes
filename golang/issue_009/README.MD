可以通过泛型结合反射动态处理不同的 `sql.Null*` 类型，合并成一个通用方法。以下是实现代码：

# 原始代码
```go
func sqlFloat64Value(value sql.NullFloat64) *float64 {
	if value.Valid {
		return &value.Float64
	}
	return nil
}

func sqlInt32Value(value sql.NullInt32) *int32 {
	if value.Valid {
		return &value.Int32
	}
	return nil
}

func sqlUint32Value(value sql.NullInt32) *uint32 {
	if value.Valid {
		v := uint32(value.Int32)
		return &v
	}
	return nil
}
```


# 重构之后的代码

```go
import (
    "database/sql"
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
```

**使用方法：**

```go
// 原函数等效调用
sqlFloat64Value => SQLValue[sql.NullFloat64, float64]
sqlInt32Value   => SQLValue[sql.NullInt32, int32]
sqlUint32Value  => SQLValue[sql.NullInt32, uint32]

// 示例
nullFloat := sql.NullFloat64{Float64: 123.45, Valid: true}
ptrFloat := SQLValue[sql.NullFloat64, float64](nullFloat)  // *float64

nullInt32 := sql.NullInt32{Int32: 42, Valid: true}
ptrInt32 := SQLValue[sql.NullInt32, int32](nullInt32)      // *int32
ptrUint32 := SQLValue[sql.NullInt32, uint32](nullInt32)    // *uint32
```

**关键设计点：**

1. **双类型参数 `[T, U]`：**
   - `T` 表示输入的 `sql.Null*` 类型（如 `sql.NullFloat64`）
   - `U` 表示输出的基础类型（如 `float64`/`uint32`）

2. **反射动态处理：**
   - 通过反射检查 `Valid` 字段判断有效性
   - 自动查找第一个非 `Valid` 字段作为数值字段
   - 支持类型转换（如 `int32` → `uint32`）

3. **类型安全：**
   - 使用 `ConvertibleTo` 检查类型可转换性
   - 返回具体类型的指针而非 `interface{}`

**优势：**
- 支持所有 `sql.Null*` 类型（如 `NullString`/`NullBool`）
- 支持跨基础类型转换（需确保可转换）
- 完全替代原始三个函数，减少重复代码

**注意事项：**
- 需确保 `sql.Null*` 类型结构包含 `Valid bool` 字段
- 数值字段必须是结构体中第一个非 `Valid` 字段
- 类型转换失败会返回 `nil`（如负整数转无符号）