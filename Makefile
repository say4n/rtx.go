.PHONY: all

all: clean build

build: main.go
	mkdir build
	go build -o build/rtx main.go

clean:
	rm -rf build/
