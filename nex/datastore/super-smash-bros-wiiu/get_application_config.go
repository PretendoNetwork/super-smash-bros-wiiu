package nex_datastore_super_smash_bros_4

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func GetApplicationConfig(err error, packet nex.PacketInterface, callID uint32, applicationID types.UInt32) (*nex.RMCMessage, *nex.Error) {
	fmt.Printf("########## Application ID: %d\n", applicationID)

	config := types.NewList[types.String]()
	config = append(config, types.NewString("1000000")) // 1000000
	config = append(config, types.NewString("100"))     // 100
	config = append(config, types.NewString("0"))       // 0
	config = append(config, types.NewString("0"))       // 0
	config = append(config, types.NewString("1"))       // 1
	config = append(config, types.NewString("2"))       // 1   ENABLE TOURNAMENTS WITH 2
	config = append(config, types.NewString("0"))       // 0
	config = append(config, types.NewString("30"))      // 30
	config = append(config, types.NewString("50"))
	config = append(config, types.NewString("45"))
	config = append(config, types.NewString("30"))
	config = append(config, types.NewString("5"))
	config = append(config, types.NewString("8"))
	config = append(config, types.NewString("50"))
	config = append(config, types.NewString("87"))
	config = append(config, types.NewString("50"))
	config = append(config, types.NewString("85"))
	config = append(config, types.NewString("60"))
	config = append(config, types.NewString("60"))
	config = append(config, types.NewString("120"))
	config = append(config, types.NewString("600"))
	config = append(config, types.NewString("5"))
	config = append(config, types.NewString("600"))
	config = append(config, types.NewString("0")) // 0
	config = append(config, types.NewString("45"))
	config = append(config, types.NewString("40"))
	config = append(config, types.NewString("100"))
	config = append(config, types.NewString("100"))
	config = append(config, types.NewString("82"))
	config = append(config, types.NewString("100"))
	config = append(config, types.NewString("95"))
	config = append(config, types.NewString("79"))
	config = append(config, types.NewString("10"))
	config = append(config, types.NewString("4"))
	config = append(config, types.NewString("2"))
	config = append(config, types.NewString("2"))
	config = append(config, types.NewString("2"))
	config = append(config, types.NewString("1")) // 1
	config = append(config, types.NewString("5"))
	config = append(config, types.NewString("1")) // 1
	config = append(config, types.NewString("2"))
	config = append(config, types.NewString("30"))
	config = append(config, types.NewString("9"))
	config = append(config, types.NewString("15"))
	config = append(config, types.NewString("0")) // 0
	config = append(config, types.NewString("10"))
	config = append(config, types.NewString("10"))
	config = append(config, types.NewString("0")) // 0
	config = append(config, types.NewString("0")) // 0
	config = append(config, types.NewString("0")) // 0
	config = append(config, types.NewString("7"))
	config = append(config, types.NewString("20"))
	config = append(config, types.NewString("0")) // 0
	config = append(config, types.NewString("0")) // 0

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	config.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodGetApplicationConfig
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
