SHELL := /bin/bash

# Insecure Client / Server ----------------------------------------------------

.PHONY: run-insecure-server
run-insecure-server:
	go run -v ./src/insecure_server.go

.PHONY: run-insecure-client
run-insecure-client:
	go run -v ./src/insecure_client.go

# Secure Client / Server ------------------------------------------------------

.PHONY: run-insecure-server
run-secure-server:
	go run -v ./src/secure_server.go

.PHONY: run-secure-client
run-secure-client:
	go run -v ./src/secure_client.go

# Configured mTLS Client / Server ---------------------------------------------

.PHONY: run-mtls-server
run-mtls-server:
	go run -v ./src/mtls_server.go

.PHONY: run-mtls-client
run-mtls-client:
	go run -v ./src/mtls_client.go

# Generate Certs --------------------------------------------------------------

.PHONY: create-certs
create-certs:
	openssl req -newkey rsa:2048 \
	  -new -nodes -x509 \
	  -days 3650 \
	  -out cert.pem \
	  -keyout key.pem \
	  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"
	mkdir -p ./pki
	mv cert.pem ./pki
	mv key.pem ./pki

# Clean -----------------------------------------------------------------------

.PHONY: clean
clean:
	rm -Rf ./pki