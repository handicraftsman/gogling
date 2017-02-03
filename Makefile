br: build run
build: upd
	go build -o "gogling" cmd/gogling/*.go
run:
	./gogling
pull:
	git pull
	go get -v -u github.com/layeh/gopher-luar
	go get -v -u github.com/mattn/go-sqlite3
upd:
	go-bindata -o cmd/gogling/runtime.go err/ runtime/
push: upd
	git push
  
