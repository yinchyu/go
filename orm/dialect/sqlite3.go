package dialect

import (
	"fmt"
	// 驱动只用导入init 函数就可以 然后 标准库提供database 来处理数据。
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"time"
)

// 主要实现类型的反射， sqlite3 实现了对应的接口
type sqlite3 struct {
}

var _ Dialect = (*sqlite3)(nil)

// 初始化的时候注册对应的对象， 然后再调用
// 得到的是空的结构体， 但是可以调用对应的函数方法就可以了
func init() {

	RegisterDialect("sqlite3", &sqlite3{})
}
func (s *sqlite3) DltaTypeOf(typ reflect.Value) string {

	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.String:
		return "text"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	// 有了panic 就不用return了
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (s *sqlite3) TableExistSql(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
