package nex

import (
	"fmt"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func StartHPPServer() {
	globals.HPPServer = nex.NewHPPServer()

	globals.HPPServer.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.HPPServer.AccountDetailsByUsername = globals.AccountDetailsByUsername

	globals.HPPServer.SetAccessKey("2869ba38")

	globals.HPPServer.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== SSBWIIU - HPP ===")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID)
		fmt.Printf("Method ID: %d\n", request.MethodID)
		fmt.Println("====================")
	})

	registerHPPServerNEXProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_SSBWIIU_HPP_SERVER_PORT"))

	globals.HPPServer.Listen(port)
}
