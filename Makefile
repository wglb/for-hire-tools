all:
	touch all

install: ~/bin/jrnl

~/bin/jrnl: jrnl
	cp -vb jrnl ~/bin/jrnl
jrnl: jrnl.go
	go build jrnl.go

fmt: 
	gofmt -w jrnl.go
