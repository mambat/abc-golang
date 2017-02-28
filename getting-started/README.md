# Install Go tools by HomeBrew

```
Welcome to fish, the friendly interactive shell
Type help for instructions on how to use fish
wanglei@wangleideMacBook-Pro:~$ brew install go
```

# Set GOPATH environment variable

```
Welcome to fish, the friendly interactive shell
Type help for instructions on how to use fish
wanglei@wangleideMacBook-Pro:~$ vi ~/.config/fish/config.fish

set -x GOPATH /usr/local/Cellar/go/1.7.5/bin
```
# Test your installation

```
Welcome to fish, the friendly interactive shell
Type help for instructions on how to use fish
wanglei@wangleideMacBook-Pro:~$ touch ./hello.go
wanglei@wangleideMacBook-Pro:~$ go build
```
