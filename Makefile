setup:
	go get github.com/pkg/profile

build-n-check:
	go build
	./go-profiling-benchmarks
	go tool pprof --text ./mem.pprof

cpuprofile:
	go test -run=none -bench=ClientServerParallel4 -cpuprofile=cprof net/http
	go tool pprof --text http.test cprof
