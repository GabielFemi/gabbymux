package gabbymux


import (
	"fmt"
	"net"
	"os"
)

func isPortTaken(port string)bool  {
	tryToListen, err := net.Listen("tcp", ":" + port)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
		os.Exit(1)
	}

	tryToListen.Close()
	return true
}