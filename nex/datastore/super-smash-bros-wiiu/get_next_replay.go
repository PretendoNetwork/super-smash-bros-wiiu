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
	fakeData.Mode = types.NewPrimitiveU8(3)
	fakeData.Players = types.NewList[*datastore_super_smash_bros_4_protocol_types.DataStoreReplayPlayer]()

	p1 := datastore_super_smash_bros_4_protocol_types.NewDataStoreReplayPlayer()
	p1.PrincipalID = types.NewPrimitiveU32(packet.Sender().PID().LegacyValue())
	p1.Fighter = types.NewPrimitiveU8(10)
	p2 := datastore_super_smash_bros_4_protocol_types.NewDataStoreReplayPlayer()
	p2.PrincipalID = types.NewPrimitiveU32(packet.Sender().PID().LegacyValue() + 1)
	p2.Fighter = types.NewPrimitiveU8(11)

	fakeData.Players.Append(p1)
	fakeData.Players.Append(p2)

	fakeData.ReplayID = types.NewPrimitiveU64(5000)
	fakeData.ReplayType = types.NewPrimitiveU8(1)
	fakeData.Stage = types.NewPrimitiveU8(12)
	fakeData.Winners = types.NewList[*types.PrimitiveU32]()
	fakeData.Winners.Append(types.NewPrimitiveU32(packet.Sender().PID().LegacyValue()))

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fakeData.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodGetNextReplay
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
