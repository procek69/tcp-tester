mkdir -p build
cd client
env GOOS=windows GOARCH=386 go build .
mv tcp-client.exe ../build/tcp-client-windows.exe
cd ../server
env GOOS=windows GOARCH=386 go build .
mv tcp-server.exe ../build/tcp-server-windows.exe