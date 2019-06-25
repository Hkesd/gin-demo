.PHONY: all, clean

i int:
	mkdir -p ./bin/logs
	if [ ! -f "dev.yaml" ]; then cp .example.yaml dev.yaml ; fi


r run:
	gofmt -w -s ./common ./config ./controller ./middleware ./model ./repository ./route ./service ./util ./main.go
	go run main.go -c dev.yaml

fmt format:
	gofmt -w -s ./common ./config ./controller ./middleware ./model ./repository ./route ./service ./util ./main.go

build:
	go build -o bin/plutus.dev

all: fmt build

build-pro:
	GOOS=linux go build -o release/bin/plutus

c clean:
	go clean -i -n
	if [ -f "./bin/plutus.dev" ]; then rm ./bin/plutus.dev ; fi

