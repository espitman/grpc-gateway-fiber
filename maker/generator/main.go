package main

import (
    
    pbPrice "git.alibaba.ir/taraaz/salvation2/monorepo/pkg/protos/protogen/price_service"
    
	types "github.com/espitman/grpc-gateway-fiber/maker/_types"
	util "github.com/espitman/grpc-gateway-fiber/maker/_util"
	"reflect"
)

func getServiceMethods(service string) []types.Method {
    
	if service == "price" {
		ifs := reflect.TypeOf((*pbPrice.PriceServiceServer)(nil)).Elem()
		methods := util.GetMethods(ifs)
		return methods
	}
	
	return nil
}

func createHandlerYaml(services types.Services) error {
	filename := "handler.yaml"
	data := types.Services{
		Services: services.Services,
	}
	return util.YamlWriter(filename, data)
}

func main() {

    var services types.Services
    var data struct {
        Service []types.Service `yaml:"services"`
    }
    util.YamlReader("../service.yaml", &data)

    for _, service := range data.Service {
        methods := getServiceMethods(service.Name)
        s := types.Service{
            Name:      service.Name,
            NameUpper: service.NameUpper,
            Port:      service.Port,
            Path:      service.Path,
            Method:    methods,
        }
        services.Services = append(services.Services, s)
    }

    createHandlerYaml(services)
}
