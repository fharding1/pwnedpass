for goos in linux darwin windows
do
    for goarch in amd64
    do
        GOOS=${goos} GOARCH=${goarch} go build -o "build/pwnedpass-${goos}-${goarch}" cmd/pwnedpass/main.go
    done
done