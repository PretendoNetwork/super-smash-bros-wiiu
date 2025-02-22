package main

import (
	"sync"

	"github.com/PretendoNetwork/super-smash-bros-wiiu/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server
	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()
	//go nex.StartHPPServer()

	wg.Wait()
}
