.PHONY: clean all

all: heap

heap: heap.go
	go build $^

clean:
	rm -f heap

