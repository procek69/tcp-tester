# tcp-tester

On receiver:

[Download server program](https://github.com/procek69/tcp-tester/raw/master/build/tcp-server-windows.exe)
```bash
tcp-server -p 8080
```

[Download client program](https://github.com/procek69/tcp-tester/raw/master/build/tcp-client-windows.exe)
On client:
```
tcp-tester -p 8080 -h 127.0.0.1
```