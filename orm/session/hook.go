package session

import (
	"orm/log"
	"reflect"
)

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}

//// Find gets all eligible records
//func (s *Session) Find(values interface{}) error {
//	s.CallMethod(BeforeQuery, nil)
//	// ...
//	for rows.Next() {
//		dest := reflect.New(destType).Elem()
//		// ...
//		s.CallMethod(AfterQuery, dest.Addr().Interface())
//		// ...
//	}
//	return rows.Close()
//}
