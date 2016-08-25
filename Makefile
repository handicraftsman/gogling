all:
	go build -o "gogling" src/*.go

br:
	go build -o "gogling" src/*.go
	./gogling

test_all:
	go build -o "gogling" src/*.go
	./gogling -test=all

0_echo:
	go build -o "gogling" src/*.go
	./gogling -test=0_echo
