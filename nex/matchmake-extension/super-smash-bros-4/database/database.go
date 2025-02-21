package nex_matchmake_extension_super_smash_bros_4_database

import (
	"database/sql"
	"time"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	matchmake_extension_super_smash_bros_4_types "github.com/PretendoNetwork/super-smash-bros-wiiu/nex/matchmake-extension/super-smash-bros-4/types"
)

// inspired by/based on meganex datastore database setup
func InitDatabase() error {
	inits := []func() error{
		InitTables,
		InitSelectCommunityCompetitionByParamStmt,
	}

	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}

	return nil
}

func InitTables() error {
	// matchmaking schema should be here already so just dont bother

	_, err := globals.Postgres.Exec(`CREATE TABLE IF NOT EXISTS matchmaking.community_competitions (
		id bigint NOT NULL PRIMARY KEY,
		param_1 bigint NOT NULL DEFAULT 0,
		param_2 smallint NOT NULL DEFAULT 0,
		param_3 smallint NOT NULL DEFAULT 0,
		param_4 integer NOT NULL DEFAULT 0,
		param_5 smallint NOT NULL DEFAULT 0,
		country_id smallint NOT NULL DEFAULT 0,
		param_7 smallint NOT NULL DEFAULT 0,
		param_8 timestamp,
		param_9 bigint NOT NULL DEFAULT 0,
		param_10 bigint NOT NULL DEFAULT 0,
		param_11 timestamp,
		param_12 bytea
	)`)

	// country id is very likely correct
	// i dont know about language id or region id

	if err != nil {
		return err
	}

	globals.Logger.Success("Postgres tables (SSB4 matchmaking) created")

	return nil
}

// part of this query was stolen from https://github.com/PretendoNetwork/nex-protocols-common-go/blob/main/matchmake-extension/database/get_persistent_gatherings_by_participant.go
const selectCommunityCompetitionObject = `
	SELECT
		g.id,
		g.owner_pid,
		g.host_pid,
		g.min_participants,
		g.max_participants,
		g.participation_policy,
		g.policy_argument,
		g.flags,
		g.state,
		g.description,
		pg.community_type,
		pg.password,
		pg.attribs,
		pg.application_buffer,
		pg.participation_start_date,
		pg.participation_end_date,
		(SELECT COUNT(ms.id)
			FROM matchmaking.matchmake_sessions AS ms
			INNER JOIN matchmaking.gatherings AS gms ON ms.id = gms.id
			WHERE gms.registered=true
			AND ms.matchmake_system_type=5 -- matchmake_system_type=5 is only used in matchmake sessions attached to a persistent gathering
			AND ms.attribs[1]=g.id) AS matchmake_session_count,
		COALESCE((SELECT cp.participation_count
			FROM matchmaking.community_participations AS cp
			WHERE cp.user_pid=$4
			AND cp.gathering_id=g.id), 0) AS participation_count,
		cc.param_1,
		cc.param_2,
		cc.param_3,
		cc.param_4,
		cc.param_5,
		cc.country_id,
		cc.param_7
		cc.param_8,
		cc.param_9,
		cc.param_10,
		cc.param_11,
		cc.param_12
	FROM matchmaking.gatherings AS g
	INNER JOIN matchmaking.persistent_gatherings AS pg ON g.id = pg.id
	INNER JOIN matchmaking.community_competitions AS cc ON cc.id = g.id`

// code repurposed from https://github.com/TraceEntertains/meganex/blob/main/nex/datastore/database.go
// Helper to unpack things selected with (selectObject + ` WHERE ....`)
func getCommunityCompetitionObjects(stmt *sql.Stmt, args ...any) ([]matchmake_extension_super_smash_bros_4_types.CommunityCompetition, error) {
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var results []matchmake_extension_super_smash_bros_4_types.CommunityCompetition

	for rows.Next() {
		result := matchmake_extension_super_smash_bros_4_types.NewCommunityCompetition()

		var attribs []uint32
		var participationStartDate time.Time
		var participationEndDate time.Time
		var param8 time.Time
		var param11 time.Time

		err := rows.Scan(
			// gathering
			&result.ID,
			&result.OwnerPID,
			&result.HostPID,
			&result.MinimumParticipants,
			&result.MaximumParticipants,
			&result.ParticipationPolicy,
			&result.PolicyArgument,
			&result.Flags,
			&result.State,
			&result.Description,
			// persistent gathering
			&result.CommunityType,
			&result.Password,
			&attribs,
			&result.ApplicationBuffer,
			&participationStartDate,
			&participationEndDate,
			&result.MatchmakeSessionCount,
			&result.ParticipationCount,
			// community competition
			&result.Param1,
			&result.Param2,
			&result.Param3,
			&result.Param4,
			&result.Param5,
			&result.CountryId,
			&result.Param7,
			&param8,
			&result.Param9,
			&result.Param10,
			&param11,
			&result.Param12,
		)

		if err != nil {
			return nil, err
			//globals.Logger.Error(err.Error())
			//continue
		}

		result.ParticipationStartDate.FromTimestamp(participationStartDate)
		result.ParticipationEndDate.FromTimestamp(participationEndDate)

		result.Param8.FromTimestamp(param8)
		result.Param11.FromTimestamp(param11)

		result.Attribs = make(types.List[types.UInt32], 0, len(attribs))
		for i := range attribs {
			result.Attribs = append(result.Attribs, types.NewUInt32(attribs[i]))
		}

		results = append(results, result)
	}

	return results, rows.Err()
}
