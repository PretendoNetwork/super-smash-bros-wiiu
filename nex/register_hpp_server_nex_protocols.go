package nex

import (
	common_datastore "github.com/PretendoNetwork/nex-protocols-common-go/v2/datastore"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	local_common_datastore "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/datastore/common"
)

func registerHPPServerNEXProtocols() {
	datastoreProtocol := datastore.NewProtocol()
	globals.HPPServer.RegisterServiceProtocol(datastoreProtocol)
	datastoreProtocol.GetNotificationURL = local_common_datastore.GetNotificationURL
	common_datastore.NewCommonProtocol(datastoreProtocol)
}
