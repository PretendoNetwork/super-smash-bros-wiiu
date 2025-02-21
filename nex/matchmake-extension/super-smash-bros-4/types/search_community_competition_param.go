// Package types implements all the types used by the MatchmakeExtensionSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SearchCommunityCompetitionParam is a type within the MatchmakeExtensionSuperSmashBros.4 protocol
type SearchCommunityCompetitionParam struct {
	types.Structure
	Param1 types.UInt8 // 0 = recommended, 1 = for fun, 2 = for glory
	Param2 types.UInt8 // very likely country id
	Param3 types.UInt8 // very likely competition availability? (2 is public 0 is same country)
}

// WriteTo writes the SearchCommunityCompetitionParam to the given writable
func (sscp SearchCommunityCompetitionParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sscp.Param1.WriteTo(contentWritable)
	sscp.Param2.WriteTo(contentWritable)
	sscp.Param3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sscp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SearchCommunityCompetitionParam from the given readable
func (sscp *SearchCommunityCompetitionParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sscp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SearchCommunityCompetitionParam header. %s", err.Error())
	}

	err = sscp.Param1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SearchCommunityCompetitionParam.Param1. %s", err.Error())
	}

	err = sscp.Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SearchCommunityCompetitionParam.Param2. %s", err.Error())
	}

	err = sscp.Param3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SearchCommunityCompetitionParam.Param3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SearchCommunityCompetitionParam
func (sscp SearchCommunityCompetitionParam) Copy() types.RVType {
	copied := NewSearchCommunityCompetitionParam()

	copied.StructureVersion = sscp.StructureVersion
	copied.Param1 = sscp.Param1.Copy().(types.UInt8)
	copied.Param2 = sscp.Param2.Copy().(types.UInt8)
	copied.Param3 = sscp.Param3.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given SearchCommunityCompetitionParam contains the same data as the current SearchCommunityCompetitionParam
func (sscp SearchCommunityCompetitionParam) Equals(o types.RVType) bool {
	if _, ok := o.(SearchCommunityCompetitionParam); !ok {
		return false
	}

	other := o.(SearchCommunityCompetitionParam)

	if sscp.StructureVersion != other.StructureVersion {
		return false
	}

	if sscp.Param1 != other.Param1 {
		return false
	}

	if sscp.Param2 != other.Param2 {
		return false
	}

	return sscp.Param3.Equals(other.Param3)
}

func (sscp SearchCommunityCompetitionParam) CopyRef() types.RVTypePtr {
	copied := sscp.Copy().(SearchCommunityCompetitionParam)
	return &copied
}

func (sscp *SearchCommunityCompetitionParam) Deref() types.RVType {
	return *sscp
}

// String returns the string representation of the SearchCommunityCompetitionParam
func (sscp SearchCommunityCompetitionParam) String() string {
	return sscp.FormatToString(0)
}

// FormatToString pretty-prints the SearchCommunityCompetitionParam using the provided indentation level
func (sscp SearchCommunityCompetitionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SearchCommunityCompetitionParam{\n")
	b.WriteString(fmt.Sprintf("%sParam1: %s,\n", indentationValues, sscp.Param1))
	b.WriteString(fmt.Sprintf("%sParam2: %s,\n", indentationValues, sscp.Param2))
	b.WriteString(fmt.Sprintf("%sParam3: %s,\n", indentationValues, sscp.Param3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSearchCommunityCompetitionParam returns a new SearchCommunityCompetitionParam
func NewSearchCommunityCompetitionParam() SearchCommunityCompetitionParam {
	return SearchCommunityCompetitionParam{
		Param1: types.NewUInt8(0),
		Param2: types.NewUInt8(0),
		Param3: types.NewUInt8(0),
	}
}
