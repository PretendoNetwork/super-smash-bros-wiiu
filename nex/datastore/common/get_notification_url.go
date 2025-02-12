package nex_datastore_common

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	globals "github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func GetNotificationURL(err error, packet nex.PacketInterface, callID uint32, param datastore_types.DataStoreGetNotificationURLParam) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.HPPServer.ByteStreamSettings())

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))
	fmt.Println(param.PreviousURL.String())

	urlInfo := datastore_types.NewDataStoreReqGetNotificationURLInfo()

	urlInfo.URL = types.NewString("invalid.pretendo.zip/TEst")
	urlInfo.Key = types.NewString("test")
	urlInfo.Query = types.NewString("test=test")
	urlInfo.RootCACert = types.NewBuffer(nil)

	urlInfo.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodGetNotificationURL
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
