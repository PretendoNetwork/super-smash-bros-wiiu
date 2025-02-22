// Package types implements all the types used by the MatchmakeExtensionSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TournamentCompetition is a type within the MatchmakeExtensionSuperSmashBros.4 protocol
type TournamentCompetition struct {
	types.Structure
	TournamentId    types.UInt32
	Param2          types.UInt8
	Param3          types.UInt8
	TournamentName  types.String
	Param5          types.UInt8
	StartTime       types.DateTime
	EndTime         types.DateTime
	ClosingTime     types.UInt32 // closing time in seconds
	Param9          types.UInt32
	Param10         types.UInt8
	Param11         types.UInt8
	MaxParticipants types.UInt16
	Param13         types.UInt16
	Param14         types.UInt8
	Param15         types.UInt8
	RoundTime       types.UInt16 // round time in seconds
	StockCount      types.UInt8
}

// WriteTo writes the TournamentCompetition to the given writable
func (tc TournamentCompetition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	tc.TournamentId.WriteTo(contentWritable)
	tc.Param2.WriteTo(contentWritable)
	tc.Param3.WriteTo(contentWritable)
	tc.TournamentName.WriteTo(contentWritable)
	tc.Param5.WriteTo(contentWritable)
	tc.StartTime.WriteTo(contentWritable)
	tc.EndTime.WriteTo(contentWritable)
	tc.ClosingTime.WriteTo(contentWritable)
	tc.Param9.WriteTo(contentWritable)
	tc.Param10.WriteTo(contentWritable)
	tc.Param11.WriteTo(contentWritable)
	tc.MaxParticipants.WriteTo(contentWritable)
	tc.Param13.WriteTo(contentWritable)
	tc.Param14.WriteTo(contentWritable)
	tc.Param15.WriteTo(contentWritable)
	tc.RoundTime.WriteTo(contentWritable)
	tc.StockCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	tc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TournamentCompetition from the given readable
func (tc *TournamentCompetition) ExtractFrom(readable types.Readable) error {
	var err error

	err = tc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition header. %s", err.Error())
	}

	err = tc.TournamentId.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.TournamentId. %s", err.Error())
	}

	err = tc.Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param2. %s", err.Error())
	}

	err = tc.Param3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param3. %s", err.Error())
	}

	err = tc.TournamentName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.TournamentName. %s", err.Error())
	}

	err = tc.Param5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param5. %s", err.Error())
	}

	err = tc.StartTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.StartTime. %s", err.Error())
	}

	err = tc.EndTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.EndTime. %s", err.Error())
	}

	err = tc.ClosingTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.ClosingTime. %s", err.Error())
	}

	err = tc.Param9.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param9. %s", err.Error())
	}

	err = tc.Param10.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param10. %s", err.Error())
	}

	err = tc.Param11.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param11. %s", err.Error())
	}

	err = tc.MaxParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.MaxParticipants. %s", err.Error())
	}

	err = tc.Param13.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param13. %s", err.Error())
	}

	err = tc.Param14.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param14. %s", err.Error())
	}

	err = tc.Param15.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Param15. %s", err.Error())
	}

	err = tc.RoundTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.RoundTime. %s", err.Error())
	}

	err = tc.StockCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.StockCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TournamentCompetition
func (tc TournamentCompetition) Copy() types.RVType {
	copied := NewTournamentCompetition()

	copied.StructureVersion = tc.StructureVersion
	copied.TournamentId = tc.TournamentId.Copy().(types.UInt32)
	copied.Param2 = tc.Param2.Copy().(types.UInt8)
	copied.Param3 = tc.Param3.Copy().(types.UInt8)
	copied.TournamentName = tc.TournamentName.Copy().(types.String)
	copied.Param5 = tc.Param5.Copy().(types.UInt8)
	copied.StartTime = tc.StartTime.Copy().(types.DateTime)
	copied.EndTime = tc.EndTime.Copy().(types.DateTime)
	copied.ClosingTime = tc.ClosingTime.Copy().(types.UInt32)
	copied.Param9 = tc.Param9.Copy().(types.UInt32)
	copied.Param10 = tc.Param10.Copy().(types.UInt8)
	copied.Param11 = tc.Param11.Copy().(types.UInt8)
	copied.MaxParticipants = tc.MaxParticipants.Copy().(types.UInt16)
	copied.Param13 = tc.Param13.Copy().(types.UInt16)
	copied.Param14 = tc.Param14.Copy().(types.UInt8)
	copied.Param15 = tc.Param15.Copy().(types.UInt8)
	copied.RoundTime = tc.RoundTime.Copy().(types.UInt16)
	copied.StockCount = tc.StockCount.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given TournamentCompetition contains the same data as the current TournamentCompetition
func (tc TournamentCompetition) Equals(o types.RVType) bool {
	if _, ok := o.(TournamentCompetition); !ok {
		return false
	}

	other := o.(TournamentCompetition)

	if tc.StructureVersion != other.StructureVersion {
		return false
	}

	if !tc.TournamentId.Equals(other.TournamentId) {
		return false
	}

	if !tc.Param2.Equals(other.Param2) {
		return false
	}

	if !tc.Param3.Equals(other.Param3) {
		return false
	}

	if !tc.TournamentName.Equals(other.TournamentName) {
		return false
	}

	if !tc.Param5.Equals(other.Param5) {
		return false
	}

	if !tc.StartTime.Equals(other.StartTime) {
		return false
	}

	if !tc.EndTime.Equals(other.EndTime) {
		return false
	}

	if !tc.ClosingTime.Equals(other.ClosingTime) {
		return false
	}

	if !tc.Param9.Equals(other.Param9) {
		return false
	}

	if !tc.Param10.Equals(other.Param10) {
		return false
	}

	if !tc.Param11.Equals(other.Param11) {
		return false
	}

	if !tc.MaxParticipants.Equals(other.MaxParticipants) {
		return false
	}

	if !tc.Param13.Equals(other.Param13) {
		return false
	}

	if !tc.Param14.Equals(other.Param14) {
		return false
	}

	if !tc.Param15.Equals(other.Param15) {
		return false
	}

	if !tc.RoundTime.Equals(other.RoundTime) {
		return false
	}

	return tc.StockCount.Equals(other.StockCount)
}

func (tc TournamentCompetition) CopyRef() types.RVTypePtr {
	copied := tc.Copy().(TournamentCompetition)
	return &copied
}

func (tc *TournamentCompetition) Deref() types.RVType {
	return *tc
}

// String returns the string representation of the TournamentCompetition
func (tc TournamentCompetition) String() string {
	return tc.FormatToString(0)
}

// FormatToString pretty-prints the TournamentCompetition using the provided indentation level
func (tc TournamentCompetition) FormatToString(indentationLevel int) string {
	//indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TournamentCompetition{\n")
	//b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, dsppp.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewTournamentCompetition returns a new TournamentCompetition
func NewTournamentCompetition() TournamentCompetition {
	return TournamentCompetition{
		TournamentId:    types.NewUInt32(0),
		Param2:          types.NewUInt8(0),
		Param3:          types.NewUInt8(0),
		TournamentName:  types.NewString(""),
		Param5:          types.NewUInt8(0),
		StartTime:       types.NewDateTime(0).Now(),
		EndTime:         types.NewDateTime(0).Now(),
		ClosingTime:     types.NewUInt32(0),
		Param9:          types.NewUInt32(0),
		Param10:         types.NewUInt8(0),
		Param11:         types.NewUInt8(0),
		MaxParticipants: types.NewUInt16(0),
		Param13:         types.NewUInt16(0),
		Param14:         types.NewUInt8(0),
		Param15:         types.NewUInt8(0),
		RoundTime:       types.NewUInt16(0),
		StockCount:      types.NewUInt8(0),
	}
}
