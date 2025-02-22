package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_database "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/database"
)

/*
== INCOMING PARAMETERS ==
uint32	ownerPid
bool	isOwner?

== OUTGOING PARAMETERS ==
List<CommunityCompetition> competitions
*/

func FindCommunityCompetitionsByParticipant(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(packet.RMCMessage().Parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	participant := types.NewPID(0)
	participant.ExtractFrom(parametersStream)

	unk := types.NewBool(false)
	unk.ExtractFrom(parametersStream)

	communityCompetitions := matchmake_extension_super_smash_bros_4_database.FindCommunityCompetitionsByParticipant(participant, unk)
	communityCompetitions.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodFindCommunityCompetitionsByParticipant
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
