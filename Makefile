default: build

build:
	mkdir -p bin
	go build -o bin/mvr .

clean:
	rm -f bin/mvr

run:
	go run . 

