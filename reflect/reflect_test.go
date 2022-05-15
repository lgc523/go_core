package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type A int

func TestGetType(t *testing.T) {
	var a A = 1
	ret, retP := reflect.TypeOf(a), reflect.TypeOf(&a)
	//name 真实类型，kind 基础类型 elem 指针的基类型
	fmt.Println(ret.Name(), ret.Kind(), retP.Elem(), reflect.ValueOf(a))
}

type student struct {
	name  string `form:"name"`
	skill int    `form:"sk"`
}

func TestGetStruct(t *testing.T) {
	var s student
	ret := reflect.TypeOf(&s)
	if ret.Kind() == reflect.Ptr {
		ret = ret.Elem()
	}
	fmt.Println(ret.NumField())
	for i := 0; i < ret.NumField(); i++ {
		structField := ret.Field(i)
		fmt.Println(structField.Name, structField.Type, structField.Tag, structField.Tag.Get("form"))
	}
}

type animal struct {
	Name string
	age  int
}

func TestModifyValForExportField(t *testing.T) {

	var dog animal
	elem := reflect.ValueOf(&dog).Elem()
	fmt.Printf("v:%v type:%v kind:%v\n", elem, elem.Type(), elem.Kind())
	elem.FieldByName("Name").SetString("super dog")
	fmt.Printf("v:%v type:%v kind:%v\n", elem, elem.Type(), elem.Kind())
}

type member struct {
}

func (member) MemberInfo(name string, score int) string {
	return fmt.Sprintf("姓名: %s 积分: %d\n", name, score)
}

func (member) Skill(name string, sk ...interface{}) string {
	return fmt.Sprintf("姓名: %s 特长: %v\n", name, sk)
}

func TestReflectInvokeMethod(t *testing.T) {
	var m member
	of := reflect.ValueOf(&m)
	method := of.MethodByName("MemberInfo")
	param := []reflect.Value{
		reflect.ValueOf("spider"),
		reflect.ValueOf(523),
	}
	ret := method.Call(param)
	for _, v := range ret {
		fmt.Println(v)
	}

	//variable parameter
	sk := of.MethodByName("Skill")
	strings := make([]string, 5)
	sks := append(strings, "c", "c++", "java", "go")
	skillParam := []reflect.Value{
		reflect.ValueOf("spider"),
		//stupid interface
		reflect.ValueOf([]interface{}{sks}),
	}
	skillRet := sk.CallSlice(skillParam)
	for _, v := range skillRet {
		fmt.Println(v)
	}
}
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	default:
		fmt.Println("Unknown", t)

	}
}

func TestBasicType(t *testing.T) {
	var f float32 = 12
	CheckType(&f)
	newstr
}
