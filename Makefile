TARGET_OS = "windows" "linux" "darwin"
TARGET_ARCH = "amd64" "arm64"

clean:
	rm -rf ./build/*

build: clean
	for os in $(TARGET_OS);do   																\
		for arch in $(TARGET_ARCH);do															\
			env GOOS=$$os GOARCH=amd64 go build -o ./build/$$os/$$arch/copy.exe ./cmd/copy.go; 	\
		done; 																					\
	done 

builder:
	docker build --force-rm -t locals/builder .
	docker run  --rm -v $$(pwd):/tmp/builder locals/builder
run:
	go run cmd/copy.go