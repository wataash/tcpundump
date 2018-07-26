package tcpundump

import (
	"io"
	"os"
	"os/exec"
	"fmt"
)

const (
	defaultBufSize = 4096
)

// an io.ReadCloser
type cmdReader struct {
	cmdIn io.WriteCloser
	cmdOut io.ReadCloser
	cmdErr io.ReadCloser
}

func (cr *cmdReader) readWriteErr() {
	p := make([]byte, defaultBufSize)

	for {
		n, err := cr.cmdErr.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Fprint(os.Stderr, string(p[:n]))
	}
}

func (cr *cmdReader) Read(p []byte) (int, error) {
	n, err := os.Stdin.Read(p)
	if err != nil {
		return n, err
	}
	return cr.cmdIn.Write(p)
}

func (cr *cmdReader) Close() error {
	// TODO: step in
	_ = cr.cmdIn.Close()
	_ = cr.cmdOut.Close()
	_ = cr.cmdErr.Close()

	// TODO: and kill it?

	return nil
}

func openCmdReader(command []string) (*cmdReader, error) {
	// TODO no such command
	cmd := exec.Command(command[0], command[1:]...)
	cmdIn, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	cr := &cmdReader{cmdIn, cmdOut, cmdErr}

	err = cmd.Start()
	if err != nil {
		cr.Close()
		return cr, err
	}

	go cr.readWriteErr()

	return cr, nil
}
