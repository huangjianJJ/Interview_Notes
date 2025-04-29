为了将两个函数合并为一个通用的数值转字符串方法，我们可以利用反射处理不同的数值类型，并根据类型选择适当的格式化方式。以下是实现代码：

```go
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
```

**步骤解释：**

1. **定义类型约束 `Number`：**
   - 使用泛型约束 `Number` 涵盖所有浮点数和整数类型，包括有符号、无符号及不同位数的变体。

2. **处理 nil 指针：**
   - 如果传入的指针为 nil，直接返回空字符串。

3. **反射获取值类型：**
   - 使用 `reflect.ValueOf` 获取解引用后的值，并检查其类型种类（Kind）。

4. **类型分支处理：**
   - **浮点数类型（float32/float64）：** 使用 `humanize.FormatFloat` 格式化为带千位分隔符的字符串。
   - **有符号整数类型：** 使用 `strconv.FormatInt` 转换为十进制字符串。
   - **无符号整数类型：** 使用 `strconv.FormatUint` 转换为十进制字符串。

**优点：**
- 支持广泛的数值类型，包括所有常见整数和浮点数类型。
- 保持原有函数行为：浮点数保留两位小数并添加千位分隔符，整数直接转换为十进制。
- 通过反射动态处理类型，无需为每种类型编写重复代码。

**注意事项：**
- 反射会带来一定的性能开销，在极高性能要求的场景需谨慎使用。
- 确保引入了 `github.com/dustin/go-humanize` 库。