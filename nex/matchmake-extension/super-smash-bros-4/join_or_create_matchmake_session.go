package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"

	//mmdatabase "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension/database"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

// INCOMING:
// u32 tournamentId
// u32 unk2
// Data<Gathering> anyGathering (only have seen it as MatchmakeSession)
// String message
//
// OUTGOING:
// Data<Gathering> joinedGathering
func JoinOrCreateMatchmakeSession(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	// boilerplate (needs to go in protocols common)
	endpoint := packet.Sender().Endpoint().(*nex.PRUDPEndPoint)
	parametersStream := nex.NewByteStreamIn(packet.RMCMessage().Parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	tournamentId := nex_types.NewUInt32(0)
	tournamentId.ExtractFrom(parametersStream)

	unk2 := nex_types.NewUInt32(0)
	unk2.ExtractFrom(parametersStream)

	anyGathering := nex_types.NewAnyDataHolder()
	anyGathering.ExtractFrom(parametersStream)

	message := nex_types.NewString("")
	message.ExtractFrom(parametersStream)

	// not boilerplate
	fmt.Println(anyGathering.FormatToString(0))

	globals.MatchmakingManager.Mutex.Lock()

	// TODO: somebody please fix this i hate this code this is terrible do NOT let this make it to production
	// matchmakeSession, _, nexError := mmdatabase.GetMatchmakeSessionByID(globals.MatchmakingManager, endpoint, tournamentId.Value)
	// if nexError.ResultCode == nex.ResultCodes.RendezVous.SessionVoid {

	// }

	//resultMatchmakeSession := match_making_types.NewMatchmakeSession()

	anyGathering.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodSearchCommunityCompetition
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
