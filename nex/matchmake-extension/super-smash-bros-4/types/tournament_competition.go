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
	Param1  *types.PrimitiveU32
	Param2  *types.PrimitiveU8
	Param3  *types.PrimitiveU8
	Param4  *types.String
	Param5  *types.PrimitiveU8
	Param6  *types.DateTime
	Param7  *types.DateTime
	Param8  *types.PrimitiveU32
	Param9  *types.PrimitiveU32
	Param10 *types.PrimitiveU8
	Param11 *types.PrimitiveU8
	Param12 *types.PrimitiveU16
	Param13 *types.PrimitiveU16
	Param14 *types.PrimitiveU8
	Param15 *types.PrimitiveU8
	Param16 *types.PrimitiveU16
	Param17 *types.PrimitiveU8
}

// WriteTo writes the TournamentCompetition to the given writable
func (dsppp *TournamentCompetition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsppp.Param1.WriteTo(contentWritable)
	dsppp.Param2.WriteTo(contentWritable)
	dsppp.Param3.WriteTo(contentWritable)
	dsppp.Param4.WriteTo(contentWritable)
	dsppp.Param5.WriteTo(contentWritable)
	dsppp.Param6.WriteTo(contentWritable)
	dsppp.Param7.WriteTo(contentWritable)
	dsppp.Param8.WriteTo(contentWritable)
	dsppp.Param9.WriteTo(contentWritable)
	dsppp.Param10.WriteTo(contentWritable)
	dsppp.Param11.WriteTo(contentWritable)
	dsppp.Param12.WriteTo(contentWritable)
	dsppp.Param13.WriteTo(contentWritable)
	dsppp.Param14.WriteTo(contentWritable)
	dsppp.Param15.WriteTo(contentWritable)
	dsppp.Param16.WriteTo(contentWritable)
	dsppp.Param17.WriteTo(contentWritable)

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
		Param1:  types.NewPrimitiveU32(0),
		Param2:  types.NewPrimitiveU8(0),
		Param3:  types.NewPrimitiveU8(0),
		Param4:  types.NewString(""),
		Param5:  types.NewPrimitiveU8(0),
		Param6:  types.NewDateTime(0).Now(),
		Param7:  types.NewDateTime(0).Now(),
		Param8:  types.NewPrimitiveU32(0),
		Param9:  types.NewPrimitiveU32(0),
		Param10: types.NewPrimitiveU8(0),
		Param11: types.NewPrimitiveU8(0),
		Param12: types.NewPrimitiveU16(0),
		Param13: types.NewPrimitiveU16(0),
		Param14: types.NewPrimitiveU8(0),
		Param15: types.NewPrimitiveU8(0),
		Param16: types.NewPrimitiveU16(0),
		Param17: types.NewPrimitiveU8(0),
	}

	return dsppp
}
