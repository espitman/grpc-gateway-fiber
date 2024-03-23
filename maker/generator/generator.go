package main

import (
	"github.com/espitman/grpc-gateway-fiber/maker/_types"
	util "github.com/espitman/grpc-gateway-fiber/maker/_util"
)

func createMain(s _types.Services) {
	tmplFile := "main.tmpl"
	outputFile := "main.go"
	util.Render(tmplFile, outputFile, s)
}

func main() {
	var services _types.Services
	var data struct {
		Service []_types.Service `yaml:"services"`
	}
	util.YamlReader("../service.yaml", &data)

	for _, service := range data.Service {
		services.Services = append(services.Services, _types.Service{
			Name:      service.Name,
			NameUpper: service.NameUpper,
			Port:      service.Port,
			Path:      service.Path,
		})
	}

	createMain(services)
}
