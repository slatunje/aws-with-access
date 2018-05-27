.DEFAULT_GOAL := build

deps:
	@mkdir -p .private/
	touch .private/{env,config.ini}
	direnv allow

build:
	go build -o bin/with main.go

term:
	go build -i -o bin/sealterm cmd/interactive/main.go && ./bin/sealterm
