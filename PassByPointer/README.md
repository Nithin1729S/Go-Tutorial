To create assembly code : go tool compile -N -S -l main.go
To test : go test -bench=. -count=1 -gcflags=-N