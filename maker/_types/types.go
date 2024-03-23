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
	Name   string
	Route  string
	Enable bool
	In     string
	Out    string
	Method string
	Query  bool
	Param  bool
	Body   bool
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
