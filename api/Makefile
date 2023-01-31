PB_ROOT = ./pb

protoc:
	protoc --proto_path=. --go_out=. --go_opt=module=github.com/yuki/api $(PB_ROOT)/$(project)/*.proto

test:
	protoc --go_out=. --go_opt=paths=source_relative ./pb/demo/v1/demo.proto