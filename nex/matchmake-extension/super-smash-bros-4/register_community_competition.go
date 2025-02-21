package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	mmdatabase "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension/database"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	globals "github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_database "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/database"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

func RegisterCommunityCompetition(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(packet.RMCMessage().Parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	communityCompetition := matchmake_extension_super_smash_bros_4_types.NewCommunityCompetition()
	communityCompetition.ExtractFrom(parametersStream)
	fmt.Println(communityCompetition.FormatToString(0))

	globals.MatchmakingManager.Mutex.Lock()

	createdPersistentGatherings, nexError := mmdatabase.GetCreatedPersistentGatherings(globals.MatchmakingManager, connection.PID())
	if nexError != nil {
		globals.MatchmakingManager.Mutex.Unlock()
		return nil, nexError
	}

	if createdPersistentGatherings >= globals.MatchmakeExtensionCommonProtocol.PersistentGatheringCreationMax {
		globals.MatchmakingManager.Mutex.Unlock()
		return nil, nex.NewError(nex.ResultCodes.RendezVous.PersistentGatheringCreationMax, "change_error")
	}

	nexError = matchmake_extension_super_smash_bros_4_database.CreateCommunityCompetition(&communityCompetition, connection)
	if nexError != nil {
		globals.MatchmakingManager.Mutex.Unlock()
		return nil, nexError
	}

	globals.MatchmakingManager.Mutex.Unlock()

	//communityCompetition.ID.WriteTo(rmcResponseStream)
	communityCompetitions := types.NewList[matchmake_extension_super_smash_bros_4_types.CommunityCompetition]()
	communityCompetitions = append(communityCompetitions, communityCompetition)
	communityCompetitions.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodSearchCommunityCompetition
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
