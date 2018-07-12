package main

import (
	"log"
	"github.com/wataash/tcpundump/internal/app/tcpundump"
)

func main() {
	// TODO --type=cisco,juniper,seil,tcpdump,...
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tcpundump.Tcpundump()
}
