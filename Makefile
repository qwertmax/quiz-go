# default: deps clean build
default: build

clean:
	rm -rf pt

deps:
	go get github.com/upwork/golang-upwork/api
	go get github.com/gin-gonic/gin
	go get gopkg.in/yaml.v2
	go get github.com/jinzhu/gorm
	go get github.com/bmizerany/pq

start:
	./upwork

format:
	go fmt ./...

build:
	go build -o pt main.go
