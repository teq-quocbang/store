package replace

import (
	"fmt"
	"reflect"
)

// Replace is replace value of field x.field -> y.field if that two fields is same json tag
// and same type
//
// *note: generics type not support for pointer
func Replace[T1 any, T2 any](t1 T1, t2 T2) (T2, error) {
	if reflect.TypeOf(*new(T1)).Kind() == reflect.Pointer ||
		reflect.TypeOf(*new(T2)).Kind() == reflect.Pointer {
		return t2, fmt.Errorf("can not send pointer type")
	}
	valueT1 := reflect.ValueOf(t1)
	typeT1 := reflect.TypeOf(t1)
	valueT2 := reflect.ValueOf(&t2)
	typeT2 := reflect.TypeOf(t2)

	for i := 0; i < valueT1.NumField(); i++ {
		fieldT1 := valueT1.Field(i)
		tag1 := typeT1.Field(i).Tag.Get("json")

		// if is different
		if !fieldT1.IsZero() {
			// check all field of model if found json tag is same let replace at t2(model) value
			for i := 0; i < valueT2.Elem().NumField(); i++ {
				tag2 := typeT2.Field(i).Tag.Get("json")
				if tag1 == tag2 {
					if valueT2.Elem().Field(i).CanSet() {
						if valueT2.Elem().Field(i).Type() != fieldT1.Type() {
							break // TODO: need to same type
						}
						valueT2.Elem().Field(i).Set(fieldT1)
					}
					break
				}
			}
		}
	}

	return t2, nil // return model with replaced value
}
