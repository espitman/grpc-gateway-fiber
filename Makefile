# Variables
OLD_MODULE_NAME := github.com/espitman/grpc-gateway-fiber


.PHONY: rename
rename: rename-ask rename-module rename-update-imports rename-cleanup


rename-ask:
	@read -p "Enter the new module name: " new_module_name; \
	echo $$new_module_name > new_module_name.txt


rename-module:
	@NEW_MODULE_NAME=$$(cat new_module_name.txt); \
	echo "Renaming module to $$NEW_MODULE_NAME"; \
	go mod edit -module $$NEW_MODULE_NAME

rename-update-imports:
	@NEW_MODULE_NAME=$$(cat new_module_name.txt); \
	echo "Updating import paths in Go source files with module name $$NEW_MODULE_NAME"; \
	find . -name "*.go" -type f -exec sed -i'' -e 's|$(OLD_MODULE_NAME)|'"$$NEW_MODULE_NAME"'|g' {} +

rename-cleanup:
	@echo "Cleaning up backup files"
	find . -name "*.go-e" -type f -delete
	rm new_module_name.txt

.PHONY: generate
generate:
	cd maker && make generate