build:
	go build -o build/dp-dd-gp

debug: build
	HUMAN_LOG=1 ./build/dp-dd-gp

.PHONY: build debug
