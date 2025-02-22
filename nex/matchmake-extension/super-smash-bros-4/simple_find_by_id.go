package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	//matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4/types"
)

// INCOMING:
// List<u32> tournamentIds
//
// OUTGOING:
// List<u32> playerCounts
func SimpleFindByID(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	//write to struct
	tournamentList := nex_types.NewList[nex_types.UInt32]()
	tournamentList = append(tournamentList, nex_types.NewUInt32(uint32(5))) // Current Number of Players
	tournamentList.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodSimpleFindByID
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
