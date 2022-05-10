# Hello Go!

Example of the talk https://t3easy.github.io/talks/go-fundamentals/

1. ```
   go mod init github.com/t3easy/hello-go
   ```
1. Add main.go and run it
   ```
   go run main.go
   ```
1. Add test and run it
   ```
   go test ./...
   ```
1. Install cobra-cli
   ```
   go install github.com/spf13/cobra-cli@latest
   ```
1. Rename main.go, to temp.go
1. Init module with cobra-cli
   ```
   ~/go/bin/cobra-cli init
   ```
1. Add functionality of previous `main.go` to `cmd/root.go`,  
   move `main_test.go` to `cmd/root_test.go`
1. Test and run
   ```
   go test ./...
   go run main.go
   ```
1. Add `--name` parameter that defaults to `Go`
1. Module housekeeping
   ```
   go mod tidy
   ```
1. Add `serve` command and run it
   ```
   ~/go/bin/cobra-cli add serve
   go run main.go serve
   ````
1. Move `getMessage` function to `github.com/t3easy/hello-go/internal` with capital `G`!
   As we will need it in the `serve` cmd, too.
   Run tests to check if nothing is broken.
   ```
   go test ./...
   go run main.go --name FooBar
   ```
1. Add http server for the "API"
   ```
   go run main.go serve
   ```
   Open http://127.0.0.1:4242/  
   Open http://127.0.0.1:4242/FooBar
1. Add `--port` flag to serve command
