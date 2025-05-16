.PHONY: scaffold

scaffold:
	@echo "Generating scaffold for ${name}"
	@cd tools && go run main.go ${name}