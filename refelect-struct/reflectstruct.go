package refelectstruct

import (
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

func Prease(value interface{}) map[string]Field {
	stage := make(map[string]Field)
	indirect := reflect.Indirect(reflect.ValueOf(value)).Type()
	for i := 0; i < indirect.NumField(); i++ {
		field := indirect.Field(i)
		// 解析出来名字，类型，tag 三个属性
		if !field.Anonymous && ast.IsExported(field.Name) {
			f := Field{Name: field.Name,
				Type: field.Type.String(),
			}
			if v, ok := field.Tag.Lookup("orm"); ok {
				f.Tag = v
			}
			stage[field.Name] = f
		}
	}

	return stage
}
