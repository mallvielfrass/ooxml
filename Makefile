ts := `/bin/date "+%Y-%m-%d---%H-%M-%S"`
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
run:
	go run .
test:
	go test *.go  -v -run='$(func)'
cover:
	go test *.go  -coverprofile=cover.out
	go tool cover -html=cover.out

