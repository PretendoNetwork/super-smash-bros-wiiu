// Package types implements all the types used by the MatchmakeExtensionSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	matchmaking_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

// CommunityCompetition is a type within the Matchmaking protocol
type CommunityCompetition struct {
	types.Structure
	*matchmaking_types.PersistentGathering
	Param1  *types.PrimitiveU32
	Param2  *types.PrimitiveU8
	Param3  *types.PrimitiveU8
	Param4  *types.PrimitiveU16
	Param5  *types.PrimitiveU8
	Param6  *types.PrimitiveU8
	Param7  *types.PrimitiveU8
	Param8  *types.DateTime
	Param9  *types.PrimitiveU32
	Param10 *types.PrimitiveU32
	Param11 *types.DateTime
	Param12 *types.QBuffer
}

// WriteTo writes the CommunityCompetition to the given writable
func (ms *CommunityCompetition) WriteTo(writable types.Writable) {
	//stream := writable.(*nex.ByteStreamOut)
	//libraryVersion := stream.LibraryVersions.MatchMaking

	ms.PersistentGathering.WriteTo(writable)

	contentWritable := writable.CopyNew()

	ms.Param1.WriteTo(contentWritable)
	ms.Param2.WriteTo(contentWritable)
	ms.Param3.WriteTo(contentWritable)
	ms.Param4.WriteTo(contentWritable)
	ms.Param5.WriteTo(contentWritable)
	ms.Param6.WriteTo(contentWritable)
	ms.Param7.WriteTo(contentWritable)
	ms.Param8.WriteTo(contentWritable)
	ms.Param9.WriteTo(contentWritable)
	ms.Param10.WriteTo(contentWritable)
	ms.Param11.WriteTo(contentWritable)
	ms.Param12.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ms.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CommunityCompetition from the given readable
func (ms *CommunityCompetition) ExtractFrom(readable types.Readable) error {
	/*stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.MatchMaking

	var err error

	err = ms.Gathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Gathering. %s", err.Error())
	}

	err = ms.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition header. %s", err.Error())
	}

	err = ms.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.GameMode. %s", err.Error())
	}

	err = ms.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Attributes. %s", err.Error())
	}

	err = ms.OpenParticipation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.OpenParticipation. %s", err.Error())
	}

	err = ms.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.MatchmakeSystemType. %s", err.Error())
	}

	err = ms.ApplicationBuffer.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.ApplicationBuffer. %s", err.Error())
	}

	err = ms.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.ParticipationCount. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.4.0") {
		err = ms.ProgressScore.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.ProgressScore. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = ms.SessionKey.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.SessionKey. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		err = ms.Option.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.Option. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		err = ms.MatchmakeParam.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.MatchmakeParam. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		err = ms.StartedTime.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.StartedTime. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		err = ms.UserPassword.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.UserPassword. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = ms.ReferGID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.ReferGID. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = ms.UserPasswordEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.UserPasswordEnabled. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = ms.SystemPasswordEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.SystemPasswordEnabled. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = ms.CodeWord.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract CommunityCompetition.CodeWord. %s", err.Error())
		}
	}*/

	return nil
}

// Copy returns a new copied instance of CommunityCompetition
func (ms *CommunityCompetition) Copy() types.RVType {
	copied := NewCommunityCompetition()

	copied.StructureVersion = ms.StructureVersion
	copied.PersistentGathering = ms.PersistentGathering.Copy().(*matchmaking_types.PersistentGathering)
	/*copied.GameMode = ms.GameMode.Copy().(*types.PrimitiveU32)
	copied.Attributes = ms.Attributes.Copy().(*types.List[*types.PrimitiveU32])
	copied.OpenParticipation = ms.OpenParticipation.Copy().(*types.PrimitiveBool)
	copied.MatchmakeSystemType = ms.MatchmakeSystemType.Copy().(*types.PrimitiveU32)
	copied.ApplicationBuffer = ms.ApplicationBuffer.Copy().(*types.Buffer)
	copied.ParticipationCount = ms.ParticipationCount.Copy().(*types.PrimitiveU32)
	copied.ProgressScore = ms.ProgressScore.Copy().(*types.PrimitiveU8)
	copied.SessionKey = ms.SessionKey.Copy().(*types.Buffer)
	copied.Option = ms.Option.Copy().(*types.PrimitiveU32)
	copied.MatchmakeParam = ms.MatchmakeParam.Copy().(*MatchmakeParam)
	copied.StartedTime = ms.StartedTime.Copy().(*types.DateTime)
	copied.UserPassword = ms.UserPassword.Copy().(*types.String)
	copied.ReferGID = ms.ReferGID.Copy().(*types.PrimitiveU32)
	copied.UserPasswordEnabled = ms.UserPasswordEnabled.Copy().(*types.PrimitiveBool)
	copied.SystemPasswordEnabled = ms.SystemPasswordEnabled.Copy().(*types.PrimitiveBool)
	copied.CodeWord = ms.CodeWord.Copy().(*types.String)*/

	return copied
}

// Equals checks if the given CommunityCompetition contains the same data as the current CommunityCompetition
func (ms *CommunityCompetition) Equals(o types.RVType) bool {
	if _, ok := o.(*CommunityCompetition); !ok {
		return false
	}

	other := o.(*CommunityCompetition)

	if ms.StructureVersion != other.StructureVersion {
		return false
	}

	if !ms.PersistentGathering.Equals(other.PersistentGathering) {
		return false
	}

	/*if !ms.GameMode.Equals(other.GameMode) {
		return false
	}

	if !ms.Attributes.Equals(other.Attributes) {
		return false
	}

	if !ms.OpenParticipation.Equals(other.OpenParticipation) {
		return false
	}

	if !ms.MatchmakeSystemType.Equals(other.MatchmakeSystemType) {
		return false
	}

	if !ms.ApplicationBuffer.Equals(other.ApplicationBuffer) {
		return false
	}

	if !ms.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if !ms.ProgressScore.Equals(other.ProgressScore) {
		return false
	}

	if !ms.SessionKey.Equals(other.SessionKey) {
		return false
	}

	if !ms.Option.Equals(other.Option) {
		return false
	}

	if !ms.MatchmakeParam.Equals(other.MatchmakeParam) {
		return false
	}

	if !ms.StartedTime.Equals(other.StartedTime) {
		return false
	}

	if !ms.UserPassword.Equals(other.UserPassword) {
		return false
	}

	if !ms.ReferGID.Equals(other.ReferGID) {
		return false
	}

	if !ms.UserPasswordEnabled.Equals(other.UserPasswordEnabled) {
		return false
	}

	if !ms.SystemPasswordEnabled.Equals(other.SystemPasswordEnabled) {
		return false
	}

	return ms.CodeWord.Equals(other.CodeWord)*/
	return true
}

// String returns the string representation of the CommunityCompetition
func (ms *CommunityCompetition) String() string {
	return ms.FormatToString(0)
}

// FormatToString pretty-prints the CommunityCompetition using the provided indentation level
func (ms *CommunityCompetition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CommunityCompetition{\n")
	b.WriteString(fmt.Sprintf("%sPersistentGathering (parent): %s,\n", indentationValues, ms.PersistentGathering.FormatToString(indentationLevel+1)))
	/*b.WriteString(fmt.Sprintf("%sGameMode: %s,\n", indentationValues, ms.GameMode))
	b.WriteString(fmt.Sprintf("%sAttributes: %s,\n", indentationValues, ms.Attributes))
	b.WriteString(fmt.Sprintf("%sOpenParticipation: %s,\n", indentationValues, ms.OpenParticipation))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %s,\n", indentationValues, ms.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %s,\n", indentationValues, ms.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %s,\n", indentationValues, ms.ParticipationCount))
	b.WriteString(fmt.Sprintf("%sProgressScore: %s,\n", indentationValues, ms.ProgressScore))
	b.WriteString(fmt.Sprintf("%sSessionKey: %s,\n", indentationValues, ms.SessionKey))
	b.WriteString(fmt.Sprintf("%sOption: %s,\n", indentationValues, ms.Option))
	b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, ms.MatchmakeParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStartedTime: %s,\n", indentationValues, ms.StartedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUserPassword: %s,\n", indentationValues, ms.UserPassword))
	b.WriteString(fmt.Sprintf("%sReferGID: %s,\n", indentationValues, ms.ReferGID))
	b.WriteString(fmt.Sprintf("%sUserPasswordEnabled: %s,\n", indentationValues, ms.UserPasswordEnabled))
	b.WriteString(fmt.Sprintf("%sSystemPasswordEnabled: %s,\n", indentationValues, ms.SystemPasswordEnabled))
	b.WriteString(fmt.Sprintf("%sCodeWord: %s,\n", indentationValues, ms.CodeWord))*/
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCommunityCompetition returns a new CommunityCompetition
func NewCommunityCompetition() *CommunityCompetition {
	ms := &CommunityCompetition{
		PersistentGathering: matchmaking_types.NewPersistentGathering(),
		Param1:              types.NewPrimitiveU32(0),
		Param2:              types.NewPrimitiveU8(0),
		Param3:              types.NewPrimitiveU8(0),
		Param4:              types.NewPrimitiveU16(0),
		Param5:              types.NewPrimitiveU8(0),
		Param6:              types.NewPrimitiveU8(0),
		Param7:              types.NewPrimitiveU8(0),
		Param8:              types.NewDateTime(0),
		Param9:              types.NewPrimitiveU32(0),
		Param10:             types.NewPrimitiveU32(0),
		Param11:             types.NewDateTime(0),
		Param12:             types.NewQBuffer(nil),
	}

	return ms
}
