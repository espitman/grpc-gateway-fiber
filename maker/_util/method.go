package _util

import (
	"github.com/espitman/grpc-gateway-fiber/maker/_types"
	"reflect"
	"strings"
)

func GetMethods(ifs reflect.Type) []_types.Method {
	var methods []_types.Method
	for i := 0; i < ifs.NumMethod(); i++ {
		method := ifs.Method(i)
		if method.IsExported() {
			methods = append(methods, _types.Method{
				Name:   method.Name,
				Route:  method.Name,
				Enable: true,
				In:     strings.Replace(method.Type.In(1).String(), "*price_service.", "", 1),
				Out:    strings.Replace(method.Type.Out(0).String(), "*price_service.", "", 1),
				Method: "Post",
				Query:  true,
				Param:  true,
				Body:   true,
			})
		}
	}
	return methods
}
