package nex

import (
	"fmt"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)
	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.SecureEndpoint.DefaultStreamSettings.MaxSilenceTime = 20000
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)
	globals.SecureEndpoint.ByteStreamSettings().UseStructureHeader = true

	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 6, 0))
	globals.SecureServer.AccessKey = "2869ba38"

	globals.SecureEndpoint.OnConnectionEnded(func(connection *nex.PRUDPConnection) {
		fmt.Printf("Kicked dumbass: %d", connection.PID().LegacyValue())
	})

	globals.SecureEndpoint.OnDisconnect(func(packet nex.PacketInterface) {
		fmt.Printf("Dumbass left: %d", packet.Sender().PID().LegacyValue())
	})

	globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== SSBWIIU - Secure ===")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID)
		fmt.Printf("Method ID: %d\n", request.MethodID)
		fmt.Println("====================")
	})

	registerCommonSecureServerProtocols()
	registerSecureServerNEXProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_SSBWIIU_SECURE_SERVER_PORT"))

	globals.SecureServer.Listen(port)
}
