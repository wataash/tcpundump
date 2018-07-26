# tcpundump

[![License](https://img.shields.io/badge/License-BSD%202--Clause-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Build Status](https://travis-ci.org/wataash/tcpundump.svg?branch=master)](https://travis-ci.org/wataash/tcpundump)
[![codecov](https://codecov.io/gh/wataash/tcpundump/branch/master/graph/badge.svg)](https://codecov.io/gh/wataash/tcpundump)

WIP

```sh
# example 1
tcpundump > undumped.pcapng
# => tcpundump: reading input...
# paste hex dump, ctrl-D

# example 2
ssh wsh@wataash.com 'tcpdump -i eth0 -xx' | tcpundump | wireshark -k -
# => tcpundump: reading input...
# # -q to surpress message

# example 3
tcpundump -w dump.pcapng -- ssh wsh@wataash.com
tcpdump -r undumped.pcapng
```

usage

```sh
synopsis
tcpundump [-q] [--type <type>] [-r <file>] [-w <file>]
tcpundump [-q] [--type <type>] [-w <file>] [--] command ...

<type> := {cisco | juniper | seil | tcpdump-x | tcpdump-xx}

-r, command: exclusive. if neigher specified, read from stdin
-w: if not specified, write out to stdout.
```
