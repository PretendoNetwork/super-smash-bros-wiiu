package nex_matchmake_extension_super_smash_bros_4

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_database "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/database"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

func SearchCommunityCompetition(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(packet.RMCMessage().Parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := matchmake_extension_super_smash_bros_4_types.NewSearchCommunityCompetitionParam()
	param.ExtractFrom(parametersStream)

	fmt.Println(param.FormatToString(0))

	communityCompetitions := matchmake_extension_super_smash_bros_4_database.FindCommunityCompetitionsBySearchParam(param)

	communityCompetitions.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodSearchCommunityCompetition
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
