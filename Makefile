all:
	rm -rf find-bytes
	go build find-bytes.go

clean:
	rm -rf find-bytes

install:
	sudo install -m 755 find-bytes /usr/local/bin/find-bytes

