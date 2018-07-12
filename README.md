# tcpundump

[![License](https://img.shields.io/badge/License-BSD%202--Clause-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Build Status](https://travis-ci.org/wataash/tcpundump.svg?branch=master)](https://travis-ci.org/wataash/tcpundump)
[![codecov](https://codecov.io/gh/wataash/tcpundump/branch/master/graph/badge.svg)](https://codecov.io/gh/wataash/tcpundump)

WIP

```sh
# example 1
ssh host 'tcpdump -i eth0 -x' | tcpundump | wireshark -k - 
# stderr: tcpundump: neither -w nor `command` specified, reading from stdin.
# -q to supress

# example 2
mkfifo dump
tcpundump --type juniper -w dump.pcapng -- ssh -p 10022 juniper
wireshark -k - < dump.pcapng

# example 3

```

usage

```sh
tcpundump
    [-q] [--type {cisco | juniper | seil | tcpdump-x | tcpdump-xx}]
    {
        [-r <file>] [-w <file>]
        | -w <file> [--] command ...
    }

-r, command: exclusive. if neigher specified, read from stdin
-w: if not specified, write out to stdout.
command: require -w
```
