package nex

import (
	"fmt"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	serverBuildString = "build:3_9_19_2005_0"

	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)
	globals.AuthenticationEndpoint.ByteStreamSettings().UseStructureHeader = true

	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 6, 0))
	globals.AuthenticationServer.AccessKey = "2869ba38"

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== SSBWIIU - Auth ===")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID)
		fmt.Printf("Method ID: %d\n", request.MethodID)
		fmt.Println("==================")
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_SSBWIIU_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
