package main

import (
	types "github.com/espitman/grpc-gateway-fiber/maker/_types"
	util "github.com/espitman/grpc-gateway-fiber/maker/_util"
	"strings"
)

func createHandler(h types.Handler) {
	tmplFile := "handler.tmpl"
	outputFile := "../_files/handler_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func createRouter(h types.Handler) {
	tmplFile := "router.tmpl"
	outputFile := "../_files/router_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func createClient(h types.Handler) {
	tmplFile := "client.tmpl"
	outputFile := "../_files/client_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func getCustomHandler() map[string]map[string]types.Method {
	custom := make(map[string]map[string]types.Method)
	var data struct {
		Service []types.Service `yaml:"services"`
	}
	util.YamlReader("handler_custom.yaml", &data)
	for _, service := range data.Service {
		if custom[service.Name] == nil {
			custom[service.Name] = make(map[string]types.Method)
		}
		for _, method := range service.Method {
			custom[service.Name][method.Name] = method
		}
	}
	return custom
}

func main() {
	customHandler := getCustomHandler()

	var data struct {
		Service []types.Service `yaml:"services"`
	}
	util.YamlReader("handler.yaml", &data)

	for _, service := range data.Service {
		ServiceNameUpper := util.CapitalizeFirstChar(service.Name)
		var methods []types.Method
		for _, method := range service.Method {
			cm, ok := customHandler[service.Name][method.Name]
			if ok {
				method = cm
			}
			methods = append(methods, method)
		}

		h := types.Handler{
			Name:             service.Name,
			PB:               service.Name + "pb",
			PBPath:           service.Path,
			HandlerName:      service.Name + "ServiceHandler",
			HandlerNameUpper: ServiceNameUpper + "ServiceHandler",
			ClientNameUpper:  ServiceNameUpper + "ServiceClient",
			Methods:          methods,
			RouterName:       service.Name + "ServiceRouter",
		}

		createHandler(h)
		createRouter(h)
		createClient(h)
	}

}
