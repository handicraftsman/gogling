br: build run

build:
	go build -o $(PWD)/gogling github.com/handicraftsman/gogling/cmd/gogling

run:
	$(PWD)/gogling

generate:
	bash generate-apis.sh

pull:
	git pull
	go get -v -u github.com/yuin/gopher-lua
	go get -v -u layeh.com/gopher-luar
	go get -v -u github.com/gorilla/mux
	go get -v -u github.com/mdempsky/gocode

setupLibs:
	ruby ./get-gostdlib.rb