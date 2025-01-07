// Package types implements all the types used by the MatchmakeExtensionSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Tournament is a type within the MatchmakeExtensionSuperSmashBros.4 protocol
type Tournament struct {
	types.Structure
	Param1  *types.PrimitiveU8
	Param2  *types.PrimitiveU32
	Param3  *types.PrimitiveU8
	Param4  *types.PrimitiveU8
	Param5  *types.PrimitiveU8
	Param6  *types.List[*types.PrimitiveU8]
	Param7  *types.PrimitiveU8
	Param8  *types.List[*types.PrimitiveU32]
	Param9  *types.List[*types.PrimitiveU16]
	Param10 *types.List[*types.PrimitiveU8]
	Param11 *types.List[*types.PrimitiveU8]
	Param12 *types.PrimitiveU64
	Param13 *types.PrimitiveU8
	Param14 *types.List[*types.PrimitiveU8]
}

// WriteTo writes the Tournament to the given writable
func (dsppp *Tournament) WriteTo(writable types.Writable) {
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

	content := contentWritable.Bytes()

	dsppp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Tournament from the given readable
func (dsppp *Tournament) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsppp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Tournament header. %s", err.Error())
	}

	/*err = dsppp.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Tournament.Profile. %s", err.Error())
	}*/

	return nil
}

// Copy returns a new copied instance of Tournament
func (dsppp *Tournament) Copy() types.RVType {
	copied := NewTournament()

	copied.StructureVersion = dsppp.StructureVersion
	//copied.Profile = dsppp.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given Tournament contains the same data as the current Tournament
func (dsppp *Tournament) Equals(o types.RVType) bool {
	if _, ok := o.(*Tournament); !ok {
		return false
	}

	other := o.(*Tournament)

	if dsppp.StructureVersion != other.StructureVersion {
		return false
	}

	return true //dsppp.Profile.Equals(other.Profile)
}

// String returns the string representation of the Tournament
func (dsppp *Tournament) String() string {
	return dsppp.FormatToString(0)
}

// FormatToString pretty-prints the Tournament using the provided indentation level
func (dsppp *Tournament) FormatToString(indentationLevel int) string {
	//indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Tournament{\n")
	//b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, dsppp.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewTournament returns a new Tournament
func NewTournament() *Tournament {
	dsppp := &Tournament{
		Param1:  types.NewPrimitiveU8(0),
		Param2:  types.NewPrimitiveU32(0),
		Param3:  types.NewPrimitiveU8(0),
		Param4:  types.NewPrimitiveU8(0),
		Param5:  types.NewPrimitiveU8(0),
		Param6:  types.NewList[*types.PrimitiveU8](),
		Param7:  types.NewPrimitiveU8(0),
		Param8:  types.NewList[*types.PrimitiveU32](),
		Param9:  types.NewList[*types.PrimitiveU16](),
		Param10: types.NewList[*types.PrimitiveU8](),
		Param11: types.NewList[*types.PrimitiveU8](),
		Param12: types.NewPrimitiveU64(0),
		Param13: types.NewPrimitiveU8(0),
		Param14: types.NewList[*types.PrimitiveU8](),
	}

	return dsppp
}
