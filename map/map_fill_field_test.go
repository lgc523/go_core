package map_test

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	Name string
	Age  int
}

func TestFillNAmeAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 25}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Employee)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func fillBySettings(ori interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(ori).Kind() != reflect.Ptr {
		if reflect.TypeOf(ori).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settings == nil {
		return errors.New("settings is nil")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = reflect.ValueOf(ori).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(ori)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
