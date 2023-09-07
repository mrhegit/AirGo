package other_plugin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

// 对长度不足n的数字后面补0
func Sup(i, n int64) string {
	m := fmt.Sprintf("%d", i)
	for int64(len(m)) < n {
		m = fmt.Sprintf("%s0", m)
	}
	return m
}

// struct转map
func StructToMap(data interface{}) map[string]interface{} {

	m := make(map[string]interface{})

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return m
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if tag == "-" || name == "-" {
			continue
		}
		if tag != "" {
			index := strings.Index(tag, ",")
			if index == -1 {
				name = tag
			} else {
				name = tag[:index]
			}
		}
		m[name] = v.Field(i).Interface()
	}
	return m
}

// golang struct 动态创建
func RegisterType(elem ...interface{}) map[string]reflect.Type {
	var typeRegistry = make(map[string]reflect.Type)
	for i := 0; i < len(elem); i++ {
		t := reflect.TypeOf(elem[i])
		typeRegistry[t.Name()] = t
	}
	return typeRegistry
}
func NewStruct(name string, typeRegistry map[string]reflect.Type) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	fmt.Println("elem", elem)
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true

}

// gin.Context中获取user id
func GetUserIDFromGinContext(ctx *gin.Context) (int64, bool) {
	userID, ok := ctx.Get("uID")
	return userID.(int64), ok
}

// gin.Context中获取user id
func GetUserNameFromGinContext(ctx *gin.Context) (string, bool) {
	userName, ok := ctx.Get("uName")
	return userName.(string), ok
}
