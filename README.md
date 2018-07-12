# tcpundump

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
