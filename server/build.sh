cd ../web
npm i
npm run build
mv web ../server/initialize
cd ../server
sudo apt update -y
sudo apt install gcc-aarch64-linux-gnu -y
#CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o airgo -ldflags='-s -w --extldflags "-static -fpic"' main.go
CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc go build -o airgo -ldflags='-s -w --extldflags "-static -fpic"' main.go