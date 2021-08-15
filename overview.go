package overviewsdk

import (
	"gihtub.com/evetools/overviewsdk/globalconfig/appearance"
	"github.com/lucasb-eyer/go-colorful"
	"strings"
)

type Overview struct {
	Backgrounds []appearance.Background
	Flags       []appearance.Flag
	Presets     map[*PresetName]Preset
}

type PresetName struct {
	s strings.Builder
}

func NewPresetName() *PresetName {
	return &PresetName{}
}

func (p *PresetName) Write(s string) error {
	p.s.WriteString(s)
	return nil
}

func (p *PresetName) MustWrite(s string) *PresetName {
	p.s.WriteString(s)
	return p
}

func (p PresetName) String() string {
	return p.s.String()
}

type token struct {
	s      string
	Color  colorful.HexColor
	Bold   bool
	Italic bool
}

type Preset struct {
	AlwaysShownStates []uint8 `yaml:"alwaysShownStates"`
	FilteredStates    []uint8 `yaml:"filteredStates"`
	Groups            []uint8 `yaml:"groups"`
}
