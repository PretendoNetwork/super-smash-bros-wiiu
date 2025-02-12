package nex_datastore_super_smash_bros_4

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	datastore_super_smash_bros_4_protocol_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func GetNextReplay(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	fakeData := datastore_super_smash_bros_4_protocol_types.NewDataStoreReplayMetaInfo()
	fakeData.Mode = types.NewUInt8(3)
	fakeData.Players = types.NewList[datastore_super_smash_bros_4_protocol_types.DataStoreReplayPlayer]()

	p1 := datastore_super_smash_bros_4_protocol_types.NewDataStoreReplayPlayer()
	p1.PrincipalID = types.NewUInt32(uint32(packet.Sender().PID()))
	p1.Fighter = types.NewUInt8(10)
	p2 := datastore_super_smash_bros_4_protocol_types.NewDataStoreReplayPlayer()
	p2.PrincipalID = types.NewUInt32(uint32(packet.Sender().PID()) + 1)
	p2.Fighter = types.NewUInt8(11)

	fakeData.Players = append(fakeData.Players, p1)
	fakeData.Players = append(fakeData.Players, p2)

	fakeData.ReplayID = types.NewUInt64(5000)
	fakeData.ReplayType = types.NewUInt8(1)
	fakeData.Stage = types.NewUInt8(12)
	fakeData.Winners = types.NewList[types.UInt32]()
	fakeData.Winners = append(fakeData.Winners, types.NewUInt32(uint32(packet.Sender().PID())))

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fakeData.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodGetNextReplay
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
