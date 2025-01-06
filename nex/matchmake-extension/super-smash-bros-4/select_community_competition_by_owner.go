package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

func SelectCommunityCompetitionByOwner(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	c := matchmake_extension_super_smash_bros_4_types.NewCommunityCompetition()
	c.Description = nex_types.NewString("Test")
	c.Gathering.ID = nex_types.NewPrimitiveU32(250)
	c.Gathering.HostPID = packet.Sender().PID()
	c.Gathering.MaximumParticipants = nex_types.NewPrimitiveU16(30)
	c.Gathering.MinimumParticipants = nex_types.NewPrimitiveU16(2)
	c.Gathering.OwnerPID = packet.Sender().PID()
	c.Gathering.ParticipationPolicy = nex_types.NewPrimitiveU32(1)
	c.Gathering.PolicyArgument = nex_types.NewPrimitiveU32(0)
	c.Gathering.State = nex_types.NewPrimitiveU32(0)
	c.Param1 = nex_types.NewPrimitiveU32(1513547864)
	c.Param2 = nex_types.NewPrimitiveU8(1)
	c.Param3 = nex_types.NewPrimitiveU8(1)
	c.Param4 = nex_types.NewPrimitiveU16(150)
	c.Param5 = nex_types.NewPrimitiveU8(1)
	c.Param6 = nex_types.NewPrimitiveU8(1)
	c.Param7 = nex_types.NewPrimitiveU8(1)
	c.Param8 = nex_types.NewDateTime(0).FromTimestamp(time.Now().Add(5 * time.Minute))
	c.Param9 = nex_types.NewPrimitiveU32(1513547864)
	c.Param10 = nex_types.NewPrimitiveU32(1513547864)
	c.Param11 = nex_types.NewDateTime(0).FromTimestamp(time.Now().Add(10 * time.Minute))

	//write to struct
	emptyList := nex_types.NewList[*matchmake_extension_super_smash_bros_4_types.CommunityCompetition]()
	emptyList.Append(c)
	emptyList.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodSelectCommunityCompetitionByOwner
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

/*
== INCOMING PARAMETERS ==
uint32	ownerPid
bool	isParticipant?

== OUTGOING PARAMETERS ==
unknown

*/
