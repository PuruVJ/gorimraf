set GOOS=linux
set GOARCH=amd64
go build -o ./bin/gorimraf

set GOOS=windows
set GOARCH=amd64
go build -o ./bin/gorimraf.exe
