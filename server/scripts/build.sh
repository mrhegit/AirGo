

cd ../web
npm i
# 前端打包
npm run build
# 将打包文件移动到后端，嵌入到go编译的二进制文件中
mv web ../server/web
cd ../server

# ubuntu/debian 交叉编译 arm64
sudo apt update -y
sudo apt install gcc-aarch64-linux-gnu -y
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o AirGo -ldflags='-s -w --extldflags "-static -fpic"' main.go


# ubuntu/debian 交叉编译 arm
#sudo apt update -y
#sudo apt install gcc-arm-linux-gnueabihf  -y
#CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc go build -o AirGo -ldflags='-s -w --extldflags "-static -fpic"' main.go