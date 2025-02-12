// Package types implements all the types used by the MatchmakeExtensionSuperSmashBros.4 protocol
package types

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	matchmaking_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

// CommunityCompetition is a type within the Matchmaking protocol
type CommunityCompetition struct {
	types.Structure
	matchmaking_types.PersistentGathering
	Param1     types.UInt32
	Param2     types.UInt8
	Param3     types.UInt8
	Param4     types.UInt16
	LanguageId types.UInt8
	CountryId  types.UInt8
	RegionId   types.UInt8
	Param8     types.DateTime
	Param9     types.UInt32
	Param10    types.UInt32
	Param11    types.DateTime
	Param12    types.QBuffer
}

// WriteTo writes the CommunityCompetition to the given writable
func (ms CommunityCompetition) WriteTo(writable types.Writable) {
	ms.PersistentGathering.WriteTo(writable)

	contentWritable := writable.CopyNew()

	ms.Param1.WriteTo(contentWritable)
	ms.Param2.WriteTo(contentWritable)
	ms.Param3.WriteTo(contentWritable)
	ms.Param4.WriteTo(contentWritable)
	ms.LanguageId.WriteTo(contentWritable)
	ms.CountryId.WriteTo(contentWritable)
	ms.RegionId.WriteTo(contentWritable)
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
	var err error

	err = ms.PersistentGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.PersistentGathering. %s", err.Error())
	}

	err = ms.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition header. %s", err.Error())
	}

	err = ms.Param1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param1. %s", err.Error())
	}

	err = ms.Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param2. %s", err.Error())
	}

	err = ms.Param3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param3. %s", err.Error())
	}

	err = ms.Param4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param4. %s", err.Error())
	}

	err = ms.LanguageId.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.LanguageId. %s", err.Error())
	}

	err = ms.CountryId.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.CountryId. %s", err.Error())
	}

	err = ms.RegionId.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.RegionId. %s", err.Error())
	}

	err = ms.Param8.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param8. %s", err.Error())
	}

	err = ms.Param9.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param9. %s", err.Error())
	}

	err = ms.Param10.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param10. %s", err.Error())
	}

	err = ms.Param11.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param11. %s", err.Error())
	}

	err = ms.Param12.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param12. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CommunityCompetition
func (ms CommunityCompetition) Copy() types.RVType {
	copied := NewCommunityCompetition()

	copied.StructureVersion = ms.StructureVersion
	copied.PersistentGathering = ms.PersistentGathering.Copy().(matchmaking_types.PersistentGathering)
	copied.Param1 = ms.Param1.Copy().(types.UInt32)
	copied.Param2 = ms.Param2.Copy().(types.UInt8)
	copied.Param3 = ms.Param3.Copy().(types.UInt8)
	copied.Param4 = ms.Param4.Copy().(types.UInt16)
	copied.LanguageId = ms.LanguageId.Copy().(types.UInt8)
	copied.CountryId = ms.CountryId.Copy().(types.UInt8)
	copied.RegionId = ms.RegionId.Copy().(types.UInt8)
	copied.Param8 = ms.Param8.Copy().(types.DateTime)
	copied.Param9 = ms.Param9.Copy().(types.UInt32)
	copied.Param10 = ms.Param10.Copy().(types.UInt32)
	copied.Param11 = ms.Param11.Copy().(types.DateTime)
	copied.Param12 = ms.Param12.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given CommunityCompetition contains the same data as the current CommunityCompetition
func (ms CommunityCompetition) Equals(o types.RVType) bool {
	if _, ok := o.(CommunityCompetition); !ok {
		return false
	}

	other := o.(CommunityCompetition)

	if ms.StructureVersion != other.StructureVersion {
		return false
	}

	if !ms.PersistentGathering.Equals(other.PersistentGathering) {
		return false
	}

	if !ms.Param1.Equals(other.Param1) {
		return false
	}

	if !ms.Param2.Equals(other.Param2) {
		return false
	}

	if !ms.Param3.Equals(other.Param3) {
		return false
	}

	if !ms.Param4.Equals(other.Param4) {
		return false
	}

	if !ms.LanguageId.Equals(other.LanguageId) {
		return false
	}

	if !ms.CountryId.Equals(other.CountryId) {
		return false
	}

	if !ms.RegionId.Equals(other.RegionId) {
		return false
	}

	if !ms.Param8.Equals(other.Param8) {
		return false
	}

	if !ms.Param9.Equals(other.Param9) {
		return false
	}

	if !ms.Param10.Equals(other.Param10) {
		return false
	}

	if !ms.Param11.Equals(other.Param11) {
		return false
	}

	return ms.Param12.Equals(other.Param12)
}

// String returns the string representation of the CommunityCompetition
func (ms CommunityCompetition) String() string {
	return ms.FormatToString(0)
}

// FormatToString pretty-prints the CommunityCompetition using the provided indentation level
func (ms CommunityCompetition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CommunityCompetition{\n")
	b.WriteString(fmt.Sprintf("%sPersistentGathering (parent): %s,\n", indentationValues, ms.PersistentGathering.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sParam1: %s,\n", indentationValues, ms.Param1.String()))
	b.WriteString(fmt.Sprintf("%sParam2: %s,\n", indentationValues, ms.Param2.String()))
	b.WriteString(fmt.Sprintf("%sParam3: %s,\n", indentationValues, ms.Param3.String()))
	b.WriteString(fmt.Sprintf("%sParam4: %s,\n", indentationValues, ms.Param4.String()))
	b.WriteString(fmt.Sprintf("%sLanguageId: %s,\n", indentationValues, ms.LanguageId.String()))
	b.WriteString(fmt.Sprintf("%sCountryId: %s,\n", indentationValues, ms.CountryId.String()))
	b.WriteString(fmt.Sprintf("%sRegionId: %s,\n", indentationValues, ms.RegionId.String()))
	b.WriteString(fmt.Sprintf("%sParam8: %s,\n", indentationValues, ms.Param8.String()))
	b.WriteString(fmt.Sprintf("%sParam9: %s,\n", indentationValues, ms.Param9.String()))
	b.WriteString(fmt.Sprintf("%sParam10: %s,\n", indentationValues, ms.Param10.String()))
	b.WriteString(fmt.Sprintf("%sParam11: %s,\n", indentationValues, ms.Param11.String()))
	b.WriteString(fmt.Sprintf("%sParam12: %s,\n", indentationValues, ms.Param12.String()))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCommunityCompetition returns a new CommunityCompetition
func NewCommunityCompetition() CommunityCompetition {
	return CommunityCompetition{
		PersistentGathering: matchmaking_types.NewPersistentGathering(),
		Param1:              types.NewUInt32(0),
		Param2:              types.NewUInt8(0),
		Param3:              types.NewUInt8(0),
		Param4:              types.NewUInt16(0),
		LanguageId:          types.NewUInt8(0),
		CountryId:           types.NewUInt8(0),
		RegionId:            types.NewUInt8(0),
		Param8:              types.NewDateTime(0),
		Param9:              types.NewUInt32(0),
		Param10:             types.NewUInt32(0),
		Param11:             types.NewDateTime(0),
		Param12:             types.NewQBuffer(nil),
	}
}

func (ms *CommunityCompetition) SetDebugFields(packet nex.PacketInterface) {
	ms.Gathering.ID = types.NewUInt32(10000)
	ms.Gathering.OwnerPID = types.NewPID(1446831925)
	ms.Gathering.HostPID = types.NewPID(1446831925)
	ms.Gathering.MinimumParticipants = types.NewUInt16(2)
	ms.Gathering.MaximumParticipants = types.NewUInt16(30)
	ms.Gathering.ParticipationPolicy = types.NewUInt32(95)
	ms.Gathering.PolicyArgument = types.NewUInt32(0)
	ms.Gathering.Flags = types.NewUInt32(641)
	ms.Gathering.State = types.NewUInt32(0)
	ms.Gathering.Description = types.NewString("test5")

	s := "0033d01000010000ff01000a02000000ffffffffffffffffffffffffbfbffbfe00ffab7f0000000000000000bfbffbfe00ffab7f000000000000000000000000010607000a000000000000000000000000000000000000007e914214000060010102010000000000"

	data, _ := hex.DecodeString(s)

	ms.PersistentGathering.ApplicationBuffer = types.NewBuffer(data)
	// ApplicationBuffer dumped from a request:
	// 0033d01000010000ff01000a02000000ffffffffffffffffffffffffbfbffbfe00ffab7f0000000000000000bfbffbfe00ffab7f
	// 000000000000000000000000010607000a000000000000000000000000000000000000007e914214000060010102010000000000
	ms.PersistentGathering.ParticipationStartDate.FromTimestamp(time.Now().UTC().Add(-1 * time.Hour))
	ms.PersistentGathering.ParticipationEndDate.FromTimestamp(time.Now().UTC().Add(10 * time.Hour))
	ms.Param1 = types.NewUInt32(1446831925)
	ms.Param2 = types.NewUInt8(0)
	ms.Param3 = types.NewUInt8(0)
	ms.Param4 = types.NewUInt16(10200)
	ms.LanguageId = types.NewUInt8(1)
	ms.CountryId = types.NewUInt8(49)
	ms.RegionId = types.NewUInt8(2)
	ms.Param8 = types.NewDateTime(135900726528)
	ms.Param9 = types.NewUInt32(1446831925)
	ms.Param10 = types.NewUInt32(1446831925)
	ms.Param11 = types.NewDateTime(0)
	ms.Param12 = types.NewQBuffer(nil)
}
