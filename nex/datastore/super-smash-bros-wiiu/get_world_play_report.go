package nex_datastore_super_smash_bros_4

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func GetWorldPlayReport(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	config := types.NewList[types.Int64]()

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	config.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodGetWorldPlayReport
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
