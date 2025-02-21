package nex_matchmake_extension_super_smash_bros_4_database

import (
	"database/sql"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

var SelectCommunityCompetitionByParamStmt *sql.Stmt

func FindCommunityCompetitionsBySearchParam(param matchmake_extension_super_smash_bros_4_types.SearchCommunityCompetitionParam) types.List[matchmake_extension_super_smash_bros_4_types.CommunityCompetition] {
	// TODO: implement

	c := matchmake_extension_super_smash_bros_4_types.NewCommunityCompetition()

	communityCompetitions := types.NewList[matchmake_extension_super_smash_bros_4_types.CommunityCompetition]()
	communityCompetitions = append(communityCompetitions, c)

	return communityCompetitions
}

func InitSelectCommunityCompetitionByParamStmt() error {
	stmt, err := globals.MatchmakingManager.Database.Prepare(selectCommunityCompetitionObject + ` WHERE data_id = $1 LIMIT 1`)
	if err != nil {
		return err
	}

	SelectCommunityCompetitionByParamStmt = stmt
	return nil
}
