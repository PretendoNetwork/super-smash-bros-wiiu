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
	*matchmaking_types.PersistentGathering
	Param1  *types.PrimitiveU32
	Param2  *types.PrimitiveU8
	Param3  *types.PrimitiveU8
	Param4  *types.PrimitiveU16
	Param5  *types.PrimitiveU8
	Param6  *types.PrimitiveU8
	Param7  *types.PrimitiveU8
	Param8  *types.PrimitiveU64
	Param9  *types.PrimitiveU32
	Param10 *types.PrimitiveU32
	Param11 *types.PrimitiveU64
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
	var err error

	err = ms.PersistentGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.PersistentGathering. %s", err.Error())
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

	err = ms.Param5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param5. %s", err.Error())
	}

	err = ms.Param6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param6. %s", err.Error())
	}

	err = ms.Param7.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CommunityCompetition.Param7. %s", err.Error())
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
func (ms *CommunityCompetition) Copy() types.RVType {
	copied := NewCommunityCompetition()

	copied.PersistentGathering = ms.PersistentGathering.Copy().(*matchmaking_types.PersistentGathering)
	copied.Param1 = ms.Param1.Copy().(*types.PrimitiveU32)
	copied.Param2 = ms.Param2.Copy().(*types.PrimitiveU8)
	copied.Param3 = ms.Param3.Copy().(*types.PrimitiveU8)
	copied.Param4 = ms.Param4.Copy().(*types.PrimitiveU16)
	copied.Param5 = ms.Param5.Copy().(*types.PrimitiveU8)
	copied.Param6 = ms.Param6.Copy().(*types.PrimitiveU8)
	copied.Param7 = ms.Param7.Copy().(*types.PrimitiveU8)
	copied.Param8 = ms.Param8.Copy().(*types.PrimitiveU64)
	copied.Param9 = ms.Param9.Copy().(*types.PrimitiveU32)
	copied.Param10 = ms.Param10.Copy().(*types.PrimitiveU32)
	copied.Param11 = ms.Param11.Copy().(*types.PrimitiveU64)
	copied.Param12 = ms.Param12.Copy().(*types.QBuffer)

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

	if !ms.Param5.Equals(other.Param5) {
		return false
	}

	if !ms.Param6.Equals(other.Param6) {
		return false
	}

	if !ms.Param7.Equals(other.Param7) {
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
	b.WriteString(fmt.Sprintf("%sParam1: %s,\n", indentationValues, ms.Param1.String()))
	b.WriteString(fmt.Sprintf("%sParam2: %s,\n", indentationValues, ms.Param2.String()))
	b.WriteString(fmt.Sprintf("%sParam3: %s,\n", indentationValues, ms.Param3.String()))
	b.WriteString(fmt.Sprintf("%sParam4: %s,\n", indentationValues, ms.Param4.String()))
	b.WriteString(fmt.Sprintf("%sParam5: %s,\n", indentationValues, ms.Param5.String()))
	b.WriteString(fmt.Sprintf("%sParam6: %s,\n", indentationValues, ms.Param6.String()))
	b.WriteString(fmt.Sprintf("%sParam7: %s,\n", indentationValues, ms.Param7.String()))
	b.WriteString(fmt.Sprintf("%sParam8: %s,\n", indentationValues, ms.Param8.String()))
	b.WriteString(fmt.Sprintf("%sParam9: %s,\n", indentationValues, ms.Param9.String()))
	b.WriteString(fmt.Sprintf("%sParam10: %s,\n", indentationValues, ms.Param10.String()))
	b.WriteString(fmt.Sprintf("%sParam11: %s,\n", indentationValues, ms.Param11.String()))
	b.WriteString(fmt.Sprintf("%sParam12: %s,\n", indentationValues, ms.Param12.String()))
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
		Param8:              types.NewPrimitiveU64(0),
		Param9:              types.NewPrimitiveU32(0),
		Param10:             types.NewPrimitiveU32(0),
		Param11:             types.NewPrimitiveU64(0),
		Param12:             types.NewQBuffer(nil),
	}

	return ms
}

func (ms *CommunityCompetition) SetDebugFields(packet nex.PacketInterface) {
	ms.Gathering.ID = types.NewPrimitiveU32(0xFFFFFFFF)
	ms.Gathering.OwnerPID = packet.Sender().PID()
	ms.Gathering.HostPID = packet.Sender().PID()
	ms.Gathering.MinimumParticipants = types.NewPrimitiveU16(2)
	ms.Gathering.MaximumParticipants = types.NewPrimitiveU16(30)
	ms.Gathering.ParticipationPolicy = types.NewPrimitiveU32(95)
	ms.Gathering.PolicyArgument = types.NewPrimitiveU32(0)
	ms.Gathering.Flags = types.NewPrimitiveU32(641)
	ms.Gathering.State = types.NewPrimitiveU32(0)
	ms.Gathering.Description = types.NewString("test5")

	s := "0033D01000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	data, _ := hex.DecodeString(s)

	ms.PersistentGathering.ApplicationBuffer = types.NewBuffer(data)
	// ApplicationBuffer dumped from a request:
	// 0033d01000010000ff01000a02000000ffffffffffffffffffffffffbfbffbfe00ffab7f0000000000000000bfbffbfe00ffab7f
	// 000000000000000000000000010607000a000000000000000000000000000000000000007e914214000060010102010000000000
	ms.PersistentGathering.ParticipationStartDate = types.NewDateTime(0).FromTimestamp(time.Now().UTC().Add(-1 * time.Hour))
	ms.PersistentGathering.ParticipationEndDate = types.NewDateTime(0).FromTimestamp(time.Now().UTC().Add(10 * time.Hour))
	ms.Param1 = types.NewPrimitiveU32(9472)
	ms.Param2 = types.NewPrimitiveU8(0)
	ms.Param3 = types.NewPrimitiveU8(0)
	ms.Param4 = types.NewPrimitiveU16(0)
	ms.Param5 = types.NewPrimitiveU8(0)
	ms.Param6 = types.NewPrimitiveU8(0)
	ms.Param7 = types.NewPrimitiveU8(0)
	ms.Param8 = types.NewPrimitiveU64(5802043704348846040)
	ms.Param9 = types.NewPrimitiveU32(8100)
	ms.Param10 = types.NewPrimitiveU32(0)
	ms.Param11 = types.NewPrimitiveU64(0)
	ms.Param12 = types.NewQBuffer(nil)
}
