BINARY_NAME=pikman
build:
 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}
