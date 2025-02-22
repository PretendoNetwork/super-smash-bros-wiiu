package nex_matchmake_extension_super_smash_bros_4_database

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

func FindCommunityCompetitionsByParticipant(participant types.PID, unk types.Bool) types.List[matchmake_extension_super_smash_bros_4_types.CommunityCompetition] {
	// TODO: implement

	c := matchmake_extension_super_smash_bros_4_types.NewCommunityCompetition()

	emptyList := types.NewList[matchmake_extension_super_smash_bros_4_types.CommunityCompetition]()
	emptyList = append(emptyList, c)
	return emptyList
}
