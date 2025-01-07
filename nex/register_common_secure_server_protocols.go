package nex

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	mm_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/super-smash-bros-4"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	local_matchmake_extension_super_smash_bros_4 "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4"

	nex_matchmake_extension_common "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/common"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	secure := common_secure.NewCommonProtocol(secureProtocol)
	secure.CreateReportDBRecord = func(pid *types.PID, reportID *types.PrimitiveU32, reportData *types.QBuffer) error {
		return nil
	}

	natTraversalProtocol := nat_traversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	common_match_making.NewCommonProtocol(matchMakingProtocol)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)

	matchmakeExtensionProtocol := matchmake_extension_super_smash_bros_4.NewProtocol(globals.SecureEndpoint)
	matchmakeExtensionProtocol.SearchCommunityCompetition = local_matchmake_extension_super_smash_bros_4.SearchCommunityCompetition
	matchmakeExtensionProtocol.FindCommunityCompetitionsByParticipant = local_matchmake_extension_super_smash_bros_4.FindCommunityCompetitionsByParticipant
	matchmakeExtensionProtocol.GetTournamentCompetitions = local_matchmake_extension_super_smash_bros_4.GetTournamentCompetitions
	matchmakeExtensionProtocol.SelectCommunityCompetitionByOwner = local_matchmake_extension_super_smash_bros_4.SelectCommunityCompetitionByOwner
	matchmakeExtensionProtocol.SimpleFindByID = local_matchmake_extension_super_smash_bros_4.SimpleFindByID
	matchmakeExtensionProtocol.RegisterCommunityCompetition = local_matchmake_extension_super_smash_bros_4.RegisterCommunityCompetition
	matchmakeExtensionProtocol.JoinOrCreateMatchmakeSession = local_matchmake_extension_super_smash_bros_4.JoinOrCreateMatchmakeSession
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = nex_matchmake_extension_common.CleanupSearchMatchmakeSession
	commonMatchmakeExtensionProtocol.OnAfterAutoMatchmakeWithSearchCriteriaPostpone = func(packet nex.PacketInterface, lstSearchCriteria *types.List[*mm_types.MatchmakeSessionSearchCriteria], anyGathering *types.AnyDataHolder, strMessage *types.String) {
		fmt.Println(anyGathering)
	}
}
