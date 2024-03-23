package _types

type Services struct {
	Services []Service
}

type Service struct {
	Name      string   `yaml:"name"`
	NameUpper string   `yaml:"nameUpper"`
	Port      int      `yaml:"port"`
	Path      string   `yaml:"path"`
	Method    []Method `yaml:"method"`
}

type Method struct {
	Name         string   `yaml:"name"`
	Route        string   `yaml:"route"`
	SwaggerRoute string   `yaml:"swaggerRoute"`
	Enable       bool     `yaml:"enable"`
	In           string   `yaml:"in"`
	Out          string   `yaml:"out"`
	Method       string   `yaml:"method"`
	Query        bool     `yaml:"query"`
	Param        bool     `yaml:"param"`
	Body         bool     `yaml:"body"`
	Params       []Params `yaml:"params"`
}

type Params struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Kind     string `yaml:"kind"`
	Required bool   `yaml:"required"`
}

type Handlers struct {
	Handlers []Handler
}

type Handler struct {
	Name             string
	PB               string
	PBPath           string
	HandlerName      string
	HandlerNameUpper string
	ClientName       string
	ClientNameUpper  string
	Methods          []Method
	RouterName       string
}
