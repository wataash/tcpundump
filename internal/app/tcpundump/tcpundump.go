// Converts hex-dump output of tcpdump to pcap.
package tcpundump

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"bufio"
	"io"
)

type Args struct {
	quiet     bool
	dumpType  dumpType
	fileRead  string
	fileWrite string
	command   []string
}

// TODO move (dump_type.go?)
func parseDumpType(s string) dumpType {
	switch s {
	case "juniper":
		return juniper
	default:
		return unknown
	}
}

func parseArgs() Args {
	var ret Args

	rootCmd := cobra.Command{
		// Use: "tcpundump",
		Short: "Converts tcpdump hex dump to pcap",
		// Long: `Here is a longer description`,
		Args: cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Printf("%p", cmd) // same as &rootCmd
			ret.command = args
		},
	}

	rootCmd.Flags().BoolVarP(&ret.quiet, "quiet", "q", false, "suppress messages")
	var tmp string
	rootCmd.Flags().StringVar(&tmp, "type", "", "text file containing hex dump")
	ret.dumpType = parseDumpType(tmp)
	// TODO if unknown
	rootCmd.Flags().StringVarP(&ret.fileRead, "read", "r", "", "text file containing hex dump")
	rootCmd.Flags().StringVarP(&ret.fileWrite, "write", "w", "", "pcap file name to be written out\n(standard output if not specified)")

	// TODO command (ssh ...)

	// TODO exclusive options

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	return ret
}

// --type juniper -r test/data/tcpdump_seil.txt -w dump.pcapng -- ssh -p 10022 juniper
func Tcpundump() {
	args := parseArgs()

	rd, err := openReader(args)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"tcpundump: error: %s\n", err)
		os.Exit(1) // TODO: os.EXIT_FAILURE?
	}

	for {
		lineBytes, err := readLine(rd)
		_ = err // TODO
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

		fmt.Println(dumped)
		// TODO stub: dumped -> pcap
	}
}

func openReader(args Args) (io.Reader, error) {
	if args.fileRead != "" {
		r, err := os.Open(args.fileRead)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	// TODO if args.command

	return os.Stdin, nil
}

func readLine(rd io.Reader) ([]byte, error) {
	r := bufio.NewReader(rd)

	// TODO: if too long line (65536 chars?), descard it

	for {
		line, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return []byte{}, err
		}

		// TODO: test EOF: with file, terminal ctrl-D

		if err == io.EOF && len(line) == 0 {
			break
		}

		return line, nil
	}

	// TODO: reachable?
	return []byte{}, nil
}
