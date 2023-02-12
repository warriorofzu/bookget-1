$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o target/bookget.exe .

$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o target/bookget .
