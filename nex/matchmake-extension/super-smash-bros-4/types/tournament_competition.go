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
	TournamentId    *types.PrimitiveU32
	Param2          *types.PrimitiveU8
	Param3          *types.PrimitiveU8
	TournamentName  *types.String
	Param5          *types.PrimitiveU8
	StartTime       *types.DateTime
	EndTime         *types.DateTime
	ClosingTime     *types.PrimitiveU32 // closing time in seconds
	Param9          *types.PrimitiveU32
	Param10         *types.PrimitiveU8
	Param11         *types.PrimitiveU8
	MaxParticipants *types.PrimitiveU16
	Param13         *types.PrimitiveU16
	Param14         *types.PrimitiveU8
	Param15         *types.PrimitiveU8
	RoundTime       *types.PrimitiveU16 // round time in seconds
	StockCount      *types.PrimitiveU8
}

// WriteTo writes the TournamentCompetition to the given writable
func (dsppp *TournamentCompetition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsppp.TournamentId.WriteTo(contentWritable)
	dsppp.Param2.WriteTo(contentWritable)
	dsppp.Param3.WriteTo(contentWritable)
	dsppp.TournamentName.WriteTo(contentWritable)
	dsppp.Param5.WriteTo(contentWritable)
	dsppp.StartTime.WriteTo(contentWritable)
	dsppp.EndTime.WriteTo(contentWritable)
	dsppp.ClosingTime.WriteTo(contentWritable)
	dsppp.Param9.WriteTo(contentWritable)
	dsppp.Param10.WriteTo(contentWritable)
	dsppp.Param11.WriteTo(contentWritable)
	dsppp.MaxParticipants.WriteTo(contentWritable)
	dsppp.Param13.WriteTo(contentWritable)
	dsppp.Param14.WriteTo(contentWritable)
	dsppp.Param15.WriteTo(contentWritable)
	dsppp.RoundTime.WriteTo(contentWritable)
	dsppp.StockCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsppp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TournamentCompetition from the given readable
func (dsppp *TournamentCompetition) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsppp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition header. %s", err.Error())
	}

	/*err = dsppp.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TournamentCompetition.Profile. %s", err.Error())
	}*/

	return nil
}

// Copy returns a new copied instance of TournamentCompetition
func (dsppp *TournamentCompetition) Copy() types.RVType {
	copied := NewTournamentCompetition()

	copied.StructureVersion = dsppp.StructureVersion
	//copied.Profile = dsppp.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given TournamentCompetition contains the same data as the current TournamentCompetition
func (dsppp *TournamentCompetition) Equals(o types.RVType) bool {
	if _, ok := o.(*TournamentCompetition); !ok {
		return false
	}

	other := o.(*TournamentCompetition)

	if dsppp.StructureVersion != other.StructureVersion {
		return false
	}

	return true //dsppp.Profile.Equals(other.Profile)
}

// String returns the string representation of the TournamentCompetition
func (dsppp *TournamentCompetition) String() string {
	return dsppp.FormatToString(0)
}

// FormatToString pretty-prints the TournamentCompetition using the provided indentation level
func (dsppp *TournamentCompetition) FormatToString(indentationLevel int) string {
	//indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TournamentCompetition{\n")
	//b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, dsppp.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewTournamentCompetition returns a new TournamentCompetition
func NewTournamentCompetition() *TournamentCompetition {
	dsppp := &TournamentCompetition{
		TournamentId:    types.NewPrimitiveU32(0),
		Param2:          types.NewPrimitiveU8(0),
		Param3:          types.NewPrimitiveU8(0),
		TournamentName:  types.NewString(""),
		Param5:          types.NewPrimitiveU8(0),
		StartTime:       types.NewDateTime(0).Now(),
		EndTime:         types.NewDateTime(0).Now(),
		ClosingTime:     types.NewPrimitiveU32(0),
		Param9:          types.NewPrimitiveU32(0),
		Param10:         types.NewPrimitiveU8(0),
		Param11:         types.NewPrimitiveU8(0),
		MaxParticipants: types.NewPrimitiveU16(0),
		Param13:         types.NewPrimitiveU16(0),
		Param14:         types.NewPrimitiveU8(0),
		Param15:         types.NewPrimitiveU8(0),
		RoundTime:       types.NewPrimitiveU16(0),
		StockCount:      types.NewPrimitiveU8(0),
	}

	return dsppp
}
