package nex_datastore_super_smash_bros_4

import (
	"github.com/PretendoNetwork/nex-go/v2"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	datastore_super_smash_bros_4_protocol_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

func PostProfile(err error, packet nex.PacketInterface, callID uint32, param datastore_super_smash_bros_4_protocol_types.DataStorePostProfileParam) (*nex.RMCMessage, *nex.Error) {
	//fmt.Printf("Post Param: %s\n", param.String())

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, nil)
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodPostProfile
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
