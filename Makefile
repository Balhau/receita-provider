provider_directory=~/.terraform.d/plugins/terraform.local/balhau/receita/1.0.0/darwin_amd64/

.PHONY: clean-bin
clean-bin:
	rm -rf bin

.PHONY: clean-terraform-state
clean-terraform-state:
	rm -rf example/.terraform terraform.tfstate.backup .terraform.lock.hcl terraform.tfstate

.PHONY: clean
clean: clean-bin clean-terraform-state

.PHONY: build
build:
	mkdir -p bin
	go build -o bin/terraform-provider-receita 
	go build -o bin/backend api/backend.go

.PHONY: install-provider
install-provider:
	mkdir -p $(provider_directory)
	cp bin/terraform-provider-receita $(provider_directory)

