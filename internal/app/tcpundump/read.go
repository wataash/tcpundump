package tcpundump

import (
	"io"
	"os"
	"os/exec"
	"bufio"
)

// an io.ReadCloser
type cmdReader struct {
	cmdIn io.WriteCloser
	cmdOut io.ReadCloser
	cmdErr io.ReadCloser
}

func (cr *cmdReader) Read(p []byte) (int, error) {
	n, err := os.Stdin.Read(p)
	if err != nil {
		return n, err
	}
	return cr.cmdIn.Write(p)
}

func (cr *cmdReader) Close() error {
	_ = os.Stdout.Close()
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

	// TODO here
	err = cmd.Start()
	if err != nil {
		cr.Close()
		return cr, nil
	}

	return cr, nil
}

func readLine(rd io.Reader) ([]byte, error) {
	// ioutil.ReadAll()

	r := bufio.NewReader(rd)

	// TODO: if too long line (65536 chars?), descard it

	for {
		line, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return []byte{}, err
		}

		// TODO: test EOF: with file, terminal ctrl-D

		if err == io.EOF && len(line) == 0 {
			return []byte{}, nil
		}

		return line, nil
	}
}
