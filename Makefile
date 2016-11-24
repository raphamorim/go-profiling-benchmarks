setup:
	go get github.com/pkg/profile

build-n-check:
	go build
	./go-profiling-benchmarks
	go tool pprof —text  /var/folders/d3/cfb18hc12wv6v828d7x9qm2r0000gp/T/profile560133977/cpu.pprof

cpuprofile:
	go test -run=none -bench=ClientServerParallel4 -cpuprofile=cprof net/http
	go tool pprof --text http.test cprof
