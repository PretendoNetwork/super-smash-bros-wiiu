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
	Param1    types.UInt32
	Param2    types.UInt8
	Param3    types.UInt8
	Param4    types.UInt16
	Param5    types.UInt8
	CountryId types.UInt8
	Param7    types.UInt8
	Param8    types.DateTime
	Param9    types.UInt32
	Param10   types.UInt32
	Param11   types.DateTime
	Param12   types.QBuffer
}

// WriteTo writes the CommunityCompetition to the given writable
func (cc CommunityCompetition) WriteTo(writable types.Writable) {
	cc.PersistentGathering.WriteTo(writable)

	contentWritable := writable.CopyNew()

	cc.Param1.WriteTo(contentWritable)
	cc.Param2.WriteTo(contentWritable)
	cc.Param3.WriteTo(contentWritable)
	cc.Param4.WriteTo(contentWritable)
	cc.Param5.WriteTo(contentWritable)
	cc.CountryId.WriteTo(contentWritable)
	cc.Param7.WriteTo(contentWritable)
	cc.Param8.WriteTo(contentWritable)
	cc.Param9.WriteTo(contentWritable)
	cc.Param10.WriteTo(contentWritable)
	cc.Param11.WriteTo(contentWritable)
	cc.Param12.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	cc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CommunityCompetition from the given readable
func (cc *CommunityCompetition) ExtractFrom(readable types.Readable) error {
	var err error

	err = cc.PersistentGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.PersistentGathering. %s", err.Error())
	}

	err = cc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition header. %s", err.Error())
	}

	err = cc.Param1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param1. %s", err.Error())
	}

	err = cc.Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param2. %s", err.Error())
	}

	err = cc.Param3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param3. %s", err.Error())
	}

	err = cc.Param4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param4. %s", err.Error())
	}

	err = cc.Param5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.LanguageId. %s", err.Error())
	}

	err = cc.CountryId.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.CountryId. %s", err.Error())
	}

	err = cc.Param7.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.RegionId. %s", err.Error())
	}

	err = cc.Param8.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param8. %s", err.Error())
	}

	err = cc.Param9.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param9. %s", err.Error())
	}

	err = cc.Param10.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param10. %s", err.Error())
	}

	err = cc.Param11.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param11. %s", err.Error())
	}

	err = cc.Param12.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param12. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CommunityCompetition
func (cc CommunityCompetition) Copy() types.RVType {
	copied := NewCommunityCompetition()

	copied.StructureVersion = cc.StructureVersion
	copied.PersistentGathering = cc.PersistentGathering.Copy().(matchmaking_types.PersistentGathering)
	copied.Param1 = cc.Param1.Copy().(types.UInt32)
	copied.Param2 = cc.Param2.Copy().(types.UInt8)
	copied.Param3 = cc.Param3.Copy().(types.UInt8)
	copied.Param4 = cc.Param4.Copy().(types.UInt16)
	copied.Param5 = cc.Param5.Copy().(types.UInt8)
	copied.CountryId = cc.CountryId.Copy().(types.UInt8)
	copied.Param7 = cc.Param7.Copy().(types.UInt8)
	copied.Param8 = cc.Param8.Copy().(types.DateTime)
	copied.Param9 = cc.Param9.Copy().(types.UInt32)
	copied.Param10 = cc.Param10.Copy().(types.UInt32)
	copied.Param11 = cc.Param11.Copy().(types.DateTime)
	copied.Param12 = cc.Param12.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given CommunityCompetition contains the same data as the current CommunityCompetition
func (cc CommunityCompetition) Equals(o types.RVType) bool {
	if _, ok := o.(CommunityCompetition); !ok {
		return false
	}

	other := o.(CommunityCompetition)

	if cc.StructureVersion != other.StructureVersion {
		return false
	}

	if !cc.PersistentGathering.Equals(other.PersistentGathering) {
		return false
	}

	if !cc.Param1.Equals(other.Param1) {
		return false
	}

	if !cc.Param2.Equals(other.Param2) {
		return false
	}

	if !cc.Param3.Equals(other.Param3) {
		return false
	}

	if !cc.Param4.Equals(other.Param4) {
		return false
	}

	if !cc.Param5.Equals(other.Param5) {
		return false
	}

	if !cc.CountryId.Equals(other.CountryId) {
		return false
	}

	if !cc.Param7.Equals(other.Param7) {
		return false
	}

	if !cc.Param8.Equals(other.Param8) {
		return false
	}

	if !cc.Param9.Equals(other.Param9) {
		return false
	}

	if !cc.Param10.Equals(other.Param10) {
		return false
	}

	if !cc.Param11.Equals(other.Param11) {
		return false
	}

	return cc.Param12.Equals(other.Param12)
}

func (cc CommunityCompetition) CopyRef() types.RVTypePtr {
	copied := cc.Copy().(CommunityCompetition)
	return &copied
}

func (cc *CommunityCompetition) Deref() types.RVType {
	return *cc
}

// String returns the string representation of the CommunityCompetition
func (cc CommunityCompetition) String() string {
	return cc.FormatToString(0)
}

// FormatToString pretty-prints the CommunityCompetition using the provided indentation level
func (cc CommunityCompetition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CommunityCompetition{\n")
	b.WriteString(fmt.Sprintf("%sPersistentGathering (parent): %s,\n", indentationValues, cc.PersistentGathering.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sParam1: %s,\n", indentationValues, cc.Param1.String()))
	b.WriteString(fmt.Sprintf("%sParam2: %s,\n", indentationValues, cc.Param2.String()))
	b.WriteString(fmt.Sprintf("%sParam3: %s,\n", indentationValues, cc.Param3.String()))
	b.WriteString(fmt.Sprintf("%sParam4: %s,\n", indentationValues, cc.Param4.String()))
	b.WriteString(fmt.Sprintf("%sLanguageId: %s,\n", indentationValues, cc.Param5.String()))
	b.WriteString(fmt.Sprintf("%sCountryId: %s,\n", indentationValues, cc.CountryId.String()))
	b.WriteString(fmt.Sprintf("%sRegionId: %s,\n", indentationValues, cc.Param7.String()))
	b.WriteString(fmt.Sprintf("%sParam8: %s,\n", indentationValues, cc.Param8.String()))
	b.WriteString(fmt.Sprintf("%sParam9: %s,\n", indentationValues, cc.Param9.String()))
	b.WriteString(fmt.Sprintf("%sParam10: %s,\n", indentationValues, cc.Param10.String()))
	b.WriteString(fmt.Sprintf("%sParam11: %s,\n", indentationValues, cc.Param11.String()))
	b.WriteString(fmt.Sprintf("%sParam12: %s,\n", indentationValues, cc.Param12.String()))
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
		Param5:              types.NewUInt8(0),
		CountryId:           types.NewUInt8(0),
		Param7:              types.NewUInt8(0),
		Param8:              types.NewDateTime(0),
		Param9:              types.NewUInt32(0),
		Param10:             types.NewUInt32(0),
		Param11:             types.NewDateTime(0),
		Param12:             types.NewQBuffer(nil),
	}
}

func (cc *CommunityCompetition) SetDebugFields(packet nex.PacketInterface) {
	cc.Gathering.ID = types.NewUInt32(10000)
	cc.Gathering.OwnerPID = types.NewPID(1446831925)
	cc.Gathering.HostPID = types.NewPID(1446831925)
	cc.Gathering.MinimumParticipants = types.NewUInt16(2)
	cc.Gathering.MaximumParticipants = types.NewUInt16(30)
	cc.Gathering.ParticipationPolicy = types.NewUInt32(95)
	cc.Gathering.PolicyArgument = types.NewUInt32(0)
	cc.Gathering.Flags = types.NewUInt32(641)
	cc.Gathering.State = types.NewUInt32(0)
	cc.Gathering.Description = types.NewString("test5")

	s := "0033d01000010000ff01000a02000000ffffffffffffffffffffffffbfbffbfe00ffab7f0000000000000000bfbffbfe00ffab7f000000000000000000000000010607000a000000000000000000000000000000000000007e914214000060010102010000000000"

	data, _ := hex.DecodeString(s)

	cc.PersistentGathering.ApplicationBuffer = types.NewBuffer(data)
	// ApplicationBuffer dumped from a request:
	// 0033d01000010000ff01000a02000000ffffffffffffffffffffffffbfbffbfe00ffab7f0000000000000000bfbffbfe00ffab7f
	// 000000000000000000000000010607000a000000000000000000000000000000000000007e914214000060010102010000000000
	cc.PersistentGathering.ParticipationStartDate.FromTimestamp(time.Now().UTC().Add(-1 * time.Hour))
	cc.PersistentGathering.ParticipationEndDate.FromTimestamp(time.Now().UTC().Add(10 * time.Hour))
	cc.Param1 = types.NewUInt32(0)
	cc.Param2 = types.NewUInt8(0) // tourney format?
	cc.Param3 = types.NewUInt8(0)
	cc.Param4 = types.NewUInt16(40528)
	cc.Param5 = types.NewUInt8(1) // restrictions?
	cc.CountryId = types.NewUInt8(49)
	cc.Param7 = types.NewUInt8(2)
	cc.Param8 = types.NewDateTime(135900726528) // start date
	cc.Param9 = types.NewUInt32(0)              // victory condition
	cc.Param10 = types.NewUInt32(0)
	cc.Param11 = types.NewDateTime(0) // potentially creation time?
	cc.Param12 = types.NewQBuffer(nil)
}
