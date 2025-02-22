package nex

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	globals "github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func GetNotificationURL(err error, packet nex.PacketInterface, callID uint32, param datastore_types.DataStoreGetNotificationURLParam) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.HPPServer.ByteStreamSettings())

	globals.Logger.Info(hex.EncodeToString(packet.RMCMessage().Parameters))
	globals.Logger.Info(param.PreviousURL.String())

	urlInfo := datastore_types.NewDataStoreReqGetNotificationURLInfo()

	urlInfo.URL = types.NewString("https://example.com")
	urlInfo.Key = types.NewString("some/key")
	urlInfo.Query = types.NewString("?test=test")
	urlInfo.RootCACert = types.NewBuffer(nil)

	urlInfo.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodGetNotificationURL
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

func StartHPPServer() {
	globals.HPPServer = nex.NewHPPServer()

	globals.HPPServer.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.HPPServer.AccountDetailsByUsername = globals.AccountDetailsByUsername

	// TODO: does hpp actually need access keys?
	globals.HPPServer.SetAccessKey("2869ba38")

	globals.HPPServer.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== SSBWIIU - HPP ===")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID)
		fmt.Printf("Method ID: %d\n", request.MethodID)
		fmt.Println("====================")
	})

	datastoreProtocol := datastore.NewProtocol()
	globals.HPPServer.RegisterServiceProtocol(datastoreProtocol)
	datastoreProtocol.GetNotificationURL = GetNotificationURL

	port, _ := strconv.Atoi(os.Getenv("PN_SSBWIIU_HPP_SERVER_PORT"))

	globals.HPPServer.Listen(port)
}
