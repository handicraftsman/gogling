all:
	go build -o "gogling" src/*.go

br:
	go build -o "gogling" src/*.go
	./gogling
