ver="1.0"

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/bookget_v${ver}_windows/bookget.exe .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/bookget_v${ver}_linux/bookget .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/bookget_v${ver}_macos/bookget .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/bookget_v${ver}_macos/bookget_m2 .


cp cookie.txt dist/bookget_v${ver}_linux/cookie.txt
cp cookie.txt dist/bookget_v${ver}_macos/cookie.txt
cp cookie.txt dist/bookget_v${ver}_windows/cookie.txt


cd dist/ 
7za a -t7z bookget_v${ver}_windows.7z bookget_v${ver}_windows
tar cjf bookget_v${ver}_linux.tar.bz2 bookget_v${ver}_linux
tar cjf bookget_v${ver}_macos.tar.bz2 bookget_v${ver}_macos
