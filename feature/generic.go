package feature

import (
	"fmt"
)

type Orderd interface {
}

// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

func add[T ~int | ~float64](a, b T) T {
	return a + b
}

func compare[T Ordered](a, b T) bool {
	return a > b
}

// 在闭包中导入变量会提示未使用的变量
func good() {
	func() {
		x := 2
		fmt.Println(x)
	}()
}
