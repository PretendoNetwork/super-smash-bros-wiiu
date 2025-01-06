package nex_matchmake_extension_super_smash_bros_4

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

// This code is so bad it will give you an illness not yet
// discovered or named by humans. Proceed with caution.

func GetTournamentCompetitions(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	fmt.Println(hex.EncodeToString(packet.RMCMessage().Parameters))

	f, _ := os.ReadFile("parameters.txt")
	file := strings.Split(string(f), "\r\n")

	t := matchmake_extension_super_smash_bros_4_types.NewTournamentCompetition()

	p1, _ := strconv.ParseUint(file[0], 16, 32)
	t.Param1 = nex_types.NewPrimitiveU32(uint32(p1)) // TournamentID

	p2, _ := strconv.ParseUint(file[1], 10, 8)
	t.Param2 = nex_types.NewPrimitiveU8(uint8(p2)) // IsFull?

	p3, _ := strconv.ParseUint(file[2], 10, 8)
	t.Param3 = nex_types.NewPrimitiveU8(uint8(p3)) // idk

	t.Param4 = nex_types.NewString(file[3]) // TournamentName

	p4, _ := strconv.ParseUint(file[4], 10, 8)
	t.Param5 = nex_types.NewPrimitiveU8(uint8(p4))

	t.Param6 = nex_types.NewDateTime(0).FromTimestamp(time.Now().Add(-10 * time.Minute)) // StartTime
	t.Param7 = nex_types.NewDateTime(0).FromTimestamp(time.Now().Add(10 * time.Hour))    // EndTime

	p5, _ := strconv.ParseUint(file[5], 10, 32) // ClosingTime (seconds)
	t.Param8 = nex_types.NewPrimitiveU32(uint32(p5))

	p6, _ := strconv.ParseUint(file[6], 10, 32) //
	t.Param9 = nex_types.NewPrimitiveU32(uint32(p6))

	p7, _ := strconv.ParseUint(file[7], 10, 8)
	t.Param10 = nex_types.NewPrimitiveU8(uint8(p7))

	p8, _ := strconv.ParseUint(file[8], 10, 8)
	t.Param11 = nex_types.NewPrimitiveU8(uint8(p8))

	p9, _ := strconv.ParseUint(file[9], 10, 16)
	t.Param12 = nex_types.NewPrimitiveU16(uint16(p9)) // MaxParticipants

	pA, _ := strconv.ParseUint(file[10], 10, 16)
	t.Param13 = nex_types.NewPrimitiveU16(uint16(pA))

	pB, _ := strconv.ParseUint(file[11], 10, 8) //
	t.Param14 = nex_types.NewPrimitiveU8(uint8(pB))

	pC, _ := strconv.ParseUint(file[12], 2, 8) // SomeConfigBitfield?
	t.Param15 = nex_types.NewPrimitiveU8(uint8(pC))

	pD, _ := strconv.ParseUint(file[13], 10, 16) // Time (seconds)
	t.Param16 = nex_types.NewPrimitiveU16(uint16(pD))

	pE, _ := strconv.ParseUint(file[14], 10, 8) // StockCount
	t.Param17 = nex_types.NewPrimitiveU8(uint8(pE))

	//write to struct
	emptyList := nex_types.NewList[*matchmake_extension_super_smash_bros_4_types.TournamentCompetition]()
	emptyList.Append(t)
	emptyList.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = matchmake_extension_super_smash_bros_4.ProtocolID
	rmcResponse.MethodID = matchmake_extension_super_smash_bros_4.MethodGetTournamentCompetitions
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
