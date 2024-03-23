.PHONY: service
service:
	cd generator && go run generator.go
	cd generator && go run main.go
	cd generator && mv handler.yaml ../service
	cd service && go run main.go

.PHONY: app
app:
	cd app && go run main.go
	cd _files gofmt -w *.go
	cd _files && mv *.go ../..

.PHONY: swagger
swagger:
	cd .. && swag init -g ./main.go -o ./docs --parseDependency

.PHONY: ask
ask:
	@read -p "Customize handler_custom.yaml and press enter: "; \
	echo "Create App started!"

.PHONY: generate
generate: service ask app swagger

