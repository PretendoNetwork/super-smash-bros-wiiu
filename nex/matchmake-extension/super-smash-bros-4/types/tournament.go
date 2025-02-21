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
	Param1  types.UInt8
	Param2  types.UInt32
	Param3  types.UInt8
	Param4  types.UInt8
	Param5  types.UInt8
	Param6  types.List[types.UInt8]
	Param7  types.UInt8
	Param8  types.List[types.UInt32]
	Param9  types.List[types.UInt16]
	Param10 types.List[types.UInt8]
	Param11 types.List[types.UInt8]
	Param12 types.UInt64
	Param13 types.UInt8
	Param14 types.List[types.UInt8]
}

// WriteTo writes the Tournament to the given writable
func (t Tournament) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	t.Param1.WriteTo(contentWritable)
	t.Param2.WriteTo(contentWritable)
	t.Param3.WriteTo(contentWritable)
	t.Param4.WriteTo(contentWritable)
	t.Param5.WriteTo(contentWritable)
	t.Param6.WriteTo(contentWritable)
	t.Param7.WriteTo(contentWritable)
	t.Param8.WriteTo(contentWritable)
	t.Param9.WriteTo(contentWritable)
	t.Param10.WriteTo(contentWritable)
	t.Param11.WriteTo(contentWritable)
	t.Param12.WriteTo(contentWritable)
	t.Param13.WriteTo(contentWritable)
	t.Param14.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	t.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Tournament from the given readable
func (t *Tournament) ExtractFrom(readable types.Readable) error {
	var err error

	err = t.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Tournament header. %s", err.Error())
	}

	/*err = t.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Tournament.Profile. %s", err.Error())
	}*/

	return nil
}

// Copy returns a new copied instance of Tournament
func (t Tournament) Copy() types.RVType {
	copied := NewTournament()

	copied.StructureVersion = t.StructureVersion
	//copied.Profile = t.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given Tournament contains the same data as the current Tournament
func (t Tournament) Equals(o types.RVType) bool {
	if _, ok := o.(Tournament); !ok {
		return false
	}

	other := o.(Tournament)

	if t.StructureVersion != other.StructureVersion {
		return false
	}

	return true //t.Profile.Equals(other.Profile)
}

func (t Tournament) CopyRef() types.RVTypePtr {
	copied := t.Copy().(Tournament)
	return &copied
}

func (t *Tournament) Deref() types.RVType {
	return *t
}

// String returns the string representation of the Tournament
func (t Tournament) String() string {
	return t.FormatToString(0)
}

// FormatToString pretty-prints the Tournament using the provided indentation level
func (t Tournament) FormatToString(indentationLevel int) string {
	//indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Tournament{\n")
	//b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, t.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewTournament returns a new Tournament
func NewTournament() Tournament {
	return Tournament{
		Param1:  types.NewUInt8(0),
		Param2:  types.NewUInt32(0),
		Param3:  types.NewUInt8(0),
		Param4:  types.NewUInt8(0),
		Param5:  types.NewUInt8(0),
		Param6:  types.NewList[types.UInt8](),
		Param7:  types.NewUInt8(0),
		Param8:  types.NewList[types.UInt32](),
		Param9:  types.NewList[types.UInt16](),
		Param10: types.NewList[types.UInt8](),
		Param11: types.NewList[types.UInt8](),
		Param12: types.NewUInt64(0),
		Param13: types.NewUInt8(0),
		Param14: types.NewList[types.UInt8](),
	}
}
