br: build run
build:
	go build -o "gogling" src/*.go
run:
	./gogling
pull:
	git pull
	go get -v -u github.com/layeh/gopher-luar

test_all:
	go build -o "gogling" src/*.go
	./gogling -test=all

0_echo:
	go build -o "gogling" src/*.go
	./gogling -test=0_echo
