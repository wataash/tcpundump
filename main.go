package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wataash/tcpundump/internal/app/tcpundump"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var tuArgs tcpundump.Args
	rootCmd := cobra.Command{
		// Use: "tcpundump",
		Short: "Converts tcpdump hex dump to pcap",
		// Long: `Here is a longer description`,
		Args: cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Printf("%p", cmd) // same as &rootCmd
			tuArgs.Command = args
			if err := tcpundump.Tcpundump(tuArgs); err != nil {
				fmt.Fprintf(os.Stderr, "tcpundump: error: %s\n", err)
				// TODO: os.EXIT_FAILURE?
				os.Exit(1)
			}
			os.Exit(0)
		},
	}

	rootCmd.Flags().BoolVarP(&tuArgs.Quiet, "quiet", "q", false, "suppress messages")
	rootCmd.Flags().StringVar(&tuArgs.DumpType, "type", "", "text file containing hex dump")
	rootCmd.Flags().StringVarP(&tuArgs.FileRead, "read", "r", "", "text file containing hex dump")
	rootCmd.Flags().StringVarP(&tuArgs.FileWrite, "write", "w", "", "pcap file name to be written out\n(standard output if not specified)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
