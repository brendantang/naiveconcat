test:
	go test -v ./parse ./eval ./interpret

doc:
	godoc -http=:3000

example:
	go run . -verbose example.txt
