package nex_datastore_super_smash_bros_4

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	datastore_super_smash_bros_4_protocol_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func PrepareGetReplay(err error, packet nex.PacketInterface, callID uint32, param datastore_super_smash_bros_4_protocol_types.DataStorePrepareGetReplayParam) (*nex.RMCMessage, *nex.Error) {
	fakeData := datastore_types.NewDataStoreReqGetInfo()
	fakeData.DataID = types.NewUInt64(1)
	fakeData.URL = types.NewString("http://127.0.0.1")

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fakeData.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodPrepareGetReplay
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
