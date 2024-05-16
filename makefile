protoc:
	cd api/proto && protoc \
	--go_out=../protogen/golang \
	--go_opt=paths=source_relative \
	*.proto



protoco:
	cd api/proto && protoc \
    	    --go-grpc_out=../protogen/golang \
    	    --go-grpc_opt=paths=source_relative \
    	    *.proto


progatway:
	cd api/proto && protoc \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	*.proto



build:
	@go build -o ./bin/main cmd/main.go

run: build
	@./bin/main

test:
	@go test  ./... -v

commit:
	@echo "commit message:"
	@read msg; \
	git add . && git commit -m "$$msg" && git push origin master
