package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

/*
== INCOMING PARAMETERS ==
uint32	ownerPid
bool	isParticipant?

== OUTGOING PARAMETERS ==
List<CommunityCompetition> competitions
*/

func FindCommunityCompetitionsByParticipant(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	c := matchmake_extension_super_smash_bros_4_types.NewCommunityCompetition()
	c.SetDebugFields(packet)

	//write to struct
	emptyList := nex_types.NewList[*matchmake_extension_super_smash_bros_4_types.CommunityCompetition]()
	emptyList.Append(c)
	emptyList.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodFindCommunityCompetitionsByParticipant
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
