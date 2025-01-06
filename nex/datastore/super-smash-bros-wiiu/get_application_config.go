package nex_datastore_super_smash_bros_4

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func GetApplicationConfig(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error) {
	fmt.Printf("########## Application ID: %d\n", applicationID.Value)

	config := types.NewList[*types.String]()
	config.Append(types.NewString("1000000")) // 1000000
	config.Append(types.NewString("100"))     // 100
	config.Append(types.NewString("0"))       // 0
	config.Append(types.NewString("0"))       // 0
	config.Append(types.NewString("1"))       // 1
	config.Append(types.NewString("2"))       // 1   ENABLE TOURNAMENTS WITH 2
	config.Append(types.NewString("0"))       // 0
	config.Append(types.NewString("30"))      // 30
	config.Append(types.NewString("50"))
	config.Append(types.NewString("45"))
	config.Append(types.NewString("30"))
	config.Append(types.NewString("5"))
	config.Append(types.NewString("8"))
	config.Append(types.NewString("50"))
	config.Append(types.NewString("87"))
	config.Append(types.NewString("50"))
	config.Append(types.NewString("85"))
	config.Append(types.NewString("60"))
	config.Append(types.NewString("60"))
	config.Append(types.NewString("120"))
	config.Append(types.NewString("600"))
	config.Append(types.NewString("5"))
	config.Append(types.NewString("600"))
	config.Append(types.NewString("0")) // 0
	config.Append(types.NewString("45"))
	config.Append(types.NewString("40"))
	config.Append(types.NewString("100"))
	config.Append(types.NewString("100"))
	config.Append(types.NewString("82"))
	config.Append(types.NewString("100"))
	config.Append(types.NewString("95"))
	config.Append(types.NewString("79"))
	config.Append(types.NewString("10"))
	config.Append(types.NewString("4"))
	config.Append(types.NewString("2"))
	config.Append(types.NewString("2"))
	config.Append(types.NewString("2"))
	config.Append(types.NewString("1")) // 1
	config.Append(types.NewString("5"))
	config.Append(types.NewString("1")) // 1
	config.Append(types.NewString("2"))
	config.Append(types.NewString("30"))
	config.Append(types.NewString("9"))
	config.Append(types.NewString("15"))
	config.Append(types.NewString("0")) // 0
	config.Append(types.NewString("10"))
	config.Append(types.NewString("10"))
	config.Append(types.NewString("0")) // 0
	config.Append(types.NewString("0")) // 0
	config.Append(types.NewString("0")) // 0
	config.Append(types.NewString("7"))
	config.Append(types.NewString("20"))
	config.Append(types.NewString("0")) // 0
	config.Append(types.NewString("0")) // 0

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	config.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodGetApplicationConfig
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
