package nex_matchmake_extension_super_smash_bros_4_database

import (
	"github.com/PretendoNetwork/nex-go/v2"
	mmdatabase "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension/database"
	globals "github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

func CreateCommunityCompetition(communityCompetition *matchmake_extension_super_smash_bros_4_types.CommunityCompetition, connection *nex.PRUDPConnection) *nex.Error {
	nexError := mmdatabase.CreatePersistentGathering(globals.MatchmakingManager, connection, &communityCompetition.PersistentGathering)
	if nexError != nil {
		return nexError
	}

	_, err := globals.MatchmakingManager.Database.Exec(`INSERT INTO matchmaking.community_competitions (
		id,
		param_1,
		param_2,
		param_3,
		param_4,
		param_5,
		country_id,
		param_7,
		param_8,
		param_9,
		param_10,
		param_11,
		param_12
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13
	)`,
		uint32(communityCompetition.PersistentGathering.Gathering.ID),
		uint32(communityCompetition.Param1),
		uint8(communityCompetition.Param2),
		uint8(communityCompetition.Param3),
		uint16(communityCompetition.Param4),
		uint8(communityCompetition.Param5),
		uint8(communityCompetition.CountryId),
		uint8(communityCompetition.Param7),
		communityCompetition.Param8.Standard(),
		uint32(communityCompetition.Param9),
		uint32(communityCompetition.Param10),
		communityCompetition.Param11.Standard(),
		[]byte(communityCompetition.Param12),
	)

	if err != nil {
		globals.Logger.Errorf("%s", err.Error())
		return nex.NewError(nex.ResultCodes.Core.Unknown, err.Error())
	}

	return nil
}
