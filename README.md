# Go Profiling Samples

## First One: cpuprofile flag on tests 

This is the simplest one, it's only about use `-cpuprofile` flag of 'go test' command.

Check out, running: `make cpuprofile`

Will print a list of the hottest functions.

There are several output types available, the most useful ones are: `--text`, `--web` and `--list`. Run `go tool pprof` to get the complete list.

Cons: Only works only for tests.

## Second One: 