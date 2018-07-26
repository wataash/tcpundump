// Converts hex-dump output of tcpdump to pcap.
package tcpundump

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Args struct {
	Quiet     bool
	DumpType  string
	FileRead  string
	FileWrite string
	Command   []string
}

// validate arguments and open reader and writer
func openArgs(args Args) (io.ReadCloser, io.WriteCloser, dumpType, error) {
	dt, err := parseDumpType(args.DumpType)
	if err != nil {
		return nil, nil, dt, err
	}

	if args.FileRead != "" && len(args.Command) != 0 {
		err := fmt.Errorf("can't specify both `-r` and `command` (%q)",
			args.Command)
		return nil, nil, dt, err
	}

	r := io.ReadCloser(os.Stdin)
	w := io.WriteCloser(os.Stdout)

	if args.FileRead != "" {
		r2, err := os.Open(args.FileRead)
		if err != nil {
			return nil, nil, dt, err
		}
		r = r2
	}

	if args.FileWrite != "" {
		w2, err := os.OpenFile(args.FileWrite,
			os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
		if err != nil {
			return nil, nil, dt, err
		}
		w = w2
	}

	if len(args.Command) != 0 {
		r2, err := openCmdReader(args.Command)
		if err != nil {
			return nil, nil, dt, err
		}
		r = r2
	}

	return r, w, dt, nil
}

// must be deferred after openArgs()
func closeIO(r io.ReadCloser, w io.WriteCloser) {
	// don't close stdio otherwise testing will terminated
	if r != os.Stdin {
		_ = r.Close()
	}
	if w != os.Stdout {
		_ = w.Close()
	}
}

func readLine(rd io.Reader) ([]byte, error) {
	// ioutil.ReadAll()

	r := bufio.NewReader(rd)

	// TODO: if too long line (65536 chars?), descard it

	return r.ReadBytes('\n')
}

// --type juniper -r test/data/tcpdump_seil.txt -w dump.pcapng -- ssh -p 10022 juniper
func Tcpundump(args Args) error {
	r, w, dt, err := openArgs(args)
	if err != nil {
		return err
	}
	defer closeIO(r, w)

	_ = dt // TODO

	for {
		lineBytes, err := readLine(r)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// TODO f = getFile(cobra options)
		// line := f.ReadAll()

		// TODO stay []byte ?
		line := string(lineBytes)
		positions, err := lineToColsByte(line)

		switch err {
		case nil:
			break
		case inferFailed:
			continue
		default:
			bugPanic(nil, "unreachable code executed while processing: %q",
				line)
		}

		dumped, err := undump(line, positions)

		n, err := w.Write(dumped)
		if err != nil {
			return err
		}
		_ = n
	}
}
