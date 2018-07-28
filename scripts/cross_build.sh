#!/bin/sh

# Copyright (c) 2018, tcpundump authors
# All rights reserved.
# Licensed under BSD 2-Clause License.

# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

set -euv

# env GOOS=linux GOARCH=amd64     go bulid -v github.com/wataash/tcpundump
# env GOOS=windows GOARCH=amd64   go build -v github.com/wataash/tcpundump
# => ../../spf13/cobra/command_win.go:9:2: cannot find package "github.com/inconshreveable/mousetrap" in any of:
#
# env GOOS=linux GOARCH=amd64     go get -v github.com/wataash/tcpundump
# env GOOS=windows GOARCH=amd64   go get -v github.com/wataash/tcpundump
# => github.com/inconshreveable/mousetrap (download), success

# TODO: full build when tagged

# env GOOS=darwin GOARCH=386      go get -v github.com/wataash/tcpundump
env GOOS=darwin GOARCH=amd64    go get -v github.com/wataash/tcpundump
# env GOOS=dragonfly GOARCH=amd64 go get -v github.com/wataash/tcpundump
# env GOOS=freebsd GOARCH=386     go get -v github.com/wataash/tcpundump
env GOOS=freebsd GOARCH=amd64   go get -v github.com/wataash/tcpundump
# env GOOS=freebsd GOARCH=arm     go get -v github.com/wataash/tcpundump
env GOOS=linux GOARCH=386       go get -v github.com/wataash/tcpundump
env GOOS=linux GOARCH=amd64     go get -v github.com/wataash/tcpundump
env GOOS=linux GOARCH=arm       go get -v github.com/wataash/tcpundump
env GOOS=linux GOARCH=arm64     go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=mips      go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=mips64    go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=mips64le  go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=mipsle    go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=ppc64     go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=ppc64le   go get -v github.com/wataash/tcpundump
# env GOOS=linux GOARCH=s390x     go get -v github.com/wataash/tcpundump
# env GOOS=nacl GOARCH=386        go get -v github.com/wataash/tcpundump
# env GOOS=nacl GOARCH=amd64p32   go get -v github.com/wataash/tcpundump
# env GOOS=nacl GOARCH=arm        go get -v github.com/wataash/tcpundump
# env GOOS=netbsd GOARCH=386      go get -v github.com/wataash/tcpundump
env GOOS=netbsd GOARCH=amd64    go get -v github.com/wataash/tcpundump
# env GOOS=netbsd GOARCH=arm      go get -v github.com/wataash/tcpundump
# env GOOS=openbsd GOARCH=386     go get -v github.com/wataash/tcpundump
env GOOS=openbsd GOARCH=amd64   go get -v github.com/wataash/tcpundump
# env GOOS=openbsd GOARCH=arm     go get -v github.com/wataash/tcpundump
# env GOOS=plan9 GOARCH=386       go get -v github.com/wataash/tcpundump
# env GOOS=plan9 GOARCH=amd64     go get -v github.com/wataash/tcpundump
# env GOOS=plan9 GOARCH=arm       go get -v github.com/wataash/tcpundump
# env GOOS=solaris GOARCH=amd64   go get -v github.com/wataash/tcpundump
env GOOS=windows GOARCH=386     go get -v github.com/wataash/tcpundump
env GOOS=windows GOARCH=amd64   go get -v github.com/wataash/tcpundump
