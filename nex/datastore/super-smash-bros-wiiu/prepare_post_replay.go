package nex_datastore_super_smash_bros_4

import (
	"fmt"
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	datastore_super_smash_bros_4_protocol_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
)

var ServerReplaysTemp []*datastore_super_smash_bros_4_protocol_types.DataStoreReplayMetaInfo

func PreparePostReplay(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_protocol_types.DataStorePreparePostReplayParam) (*nex.RMCMessage, *nex.Error) {
	//fmt.Printf("Post Param: %s\n", param.String())

	if ServerReplaysTemp == nil {
		ServerReplaysTemp = make([]*datastore_super_smash_bros_4_protocol_types.DataStoreReplayMetaInfo, 0)
	}

	replay := datastore_super_smash_bros_4_protocol_types.NewDataStoreReplayMetaInfo()
	replay.Mode = param.Mode
	replay.Players = param.Players
	replay.ReplayID = types.NewPrimitiveU64(uint64(len(ServerReplaysTemp)))
	replay.ReplayType = param.ReplayType
	replay.Rule = param.Rule
	replay.Size = param.Size
	replay.Stage = param.Stage
	replay.Style = param.Style
	replay.Winners = param.Winners

	ServerReplaysTemp = append(ServerReplaysTemp, replay)

	if globals.Presigner == nil {
		fmt.Println("Presigner not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	dataId := replay.ReplayID.Value
	key := fmt.Sprintf("replay_%d.bin", dataId)
	URL, formData, err := globals.Presigner.PostObject(globals.S3Bucket, key, time.Minute*15)
	if err != nil {
		panic(err)
	}

	pPreparePostParam := datastore_types.NewDataStorePreparePostParam()

	pReqPostInfo := datastore_types.NewDataStoreReqPostInfo()
	pReqPostInfo.DataID = types.NewPrimitiveU64(dataId)
	pReqPostInfo.URL = types.NewString(URL.String())
	pReqPostInfo.RequestHeaders = types.NewList[*datastore_types.DataStoreKeyValue]()
	pReqPostInfo.FormFields = types.NewList[*datastore_types.DataStoreKeyValue]()
	pReqPostInfo.RootCACert = types.NewBuffer([]byte{})

	pReqPostInfo.RequestHeaders.Type = datastore_types.NewDataStoreKeyValue()
	pReqPostInfo.RequestHeaders.SetFromData([]*datastore_types.DataStoreKeyValue{})

	pReqPostInfo.FormFields.Type = datastore_types.NewDataStoreKeyValue()

	for key, value := range formData {
		field := datastore_types.NewDataStoreKeyValue()
		field.Key = types.NewString(key)
		field.Value = types.NewString(value)

		pReqPostInfo.FormFields.Append(field)
	}

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	pPreparePostParam.WriteTo(rmcResponseStream)
	pReqPostInfo.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, nil)
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodPreparePostReplay
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
