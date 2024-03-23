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

.PHONY: clean
clean:
	cd generator && rm main.go
	cd service && rm handler.yaml
	cd .. && rm client_*.go
	cd .. && rm handler_*.go
	cd .. && rm router_*.go
	cd .. && rm main.go
	cd .. && rm router.go


.PHONY: generate
generate: service ask app swagger

