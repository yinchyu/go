package dialect

import (
	"reflect"
)

var dialectsMap = map[string]Dialect{}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	//差点忘了这个地方的自动赋值
	return
}

type Dialect interface {
	DltaTypeOf(typ reflect.Value) string
	TableExistSql(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}
