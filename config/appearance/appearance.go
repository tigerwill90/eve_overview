package appearance

import (
	"errors"
	"fmt"
	"gihtub.com/evetools/overviewsdk/internal/overview"
	"strconv"
)

type BackgroundSet struct {
	Type   Type
	Color  Color
	Blink  bool
	Enable bool
}

type FlagSet struct {
	Type   Type
	Color  Color
	Blink  bool
	Enable bool
}

var validType = map[uint8]Type{
	9:  LowSecurityStatus,
	10: Pirate,
	11: Fleet,
	12: PlayerCorporation,
	13: CorporationWar,
	14: Alliance,
	15: ExcellentStanding,
	16: GoodStanding,
	17: NeutralStanding,
	18: BadStanding,
	19: TerribleStanding,
	21: Agent,
	45: Militia,
	44: MilitiaWar,
	48: NoStanding,
	49: AllyWar,
	50: Suspect,
	51: Criminal,
	52: LimitedEngagement,
	53: KillRight,
	66: NpcCorporation,
}

const (
	LowSecurityStatus Type = iota
	Pirate
	Fleet
	PlayerCorporation
	CorporationWar
	Alliance
	ExcellentStanding
	GoodStanding
	NeutralStanding
	BadStanding
	TerribleStanding
	Agent
	Militia
	MilitiaWar
	NoStanding
	AllyWar
	Suspect
	Criminal
	LimitedEngagement
	KillRight
	NpcCorporation
)

func (t Type) Uint8() uint8 {
	return [21]uint8{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 45, 44, 48, 49, 50, 51, 52, 53, 66}[t]
}

func (t Type) String() string {
	return strconv.FormatUint(uint64(t.Uint8()), 10)
}

const unknownType uint8 = 20

type Type uint8

var ErrInvalidType = errors.New("invalid appearance type")

// TODO move this somewhere else
func Make(rawOverview *overview.RawOverview) ([]BackgroundSet, []FlagSet, error) {
	colors, err := rawOverview.ParseColors()
	if err != nil && !errors.Is(err, overview.ErrNoColor) {
		return nil, nil, err
	}

	blinks, err := rawOverview.ParseBlinks()
	if err != nil && !errors.Is(err, overview.ErrNoBlinks) {
		return nil, nil, err
	}

	backgrounds := make([]BackgroundSet, 0, len(rawOverview.BackgroundOrder))
	for _, backgroundCode := range rawOverview.BackgroundOrder {
		if backgroundCode == unknownType {
			continue
		}

		bType, valid := validType[backgroundCode]
		if !valid {
			return nil, nil, fmt.Errorf("%d is not a valid background code: %w", backgroundCode, ErrInvalidType)
		}

		background := BackgroundSet{
			Type: bType,
		}

		colorCode, ok := colors[overview.Background][backgroundCode]
		if !ok {
			background.Color = Default
		} else {
			cType, valid := validColor[colorCode]
			if !valid {
				return nil, nil, fmt.Errorf("%s is not a valid color: %w", colorCode, ErrInvalidType)
			}
			background.Color = cType
		}

		background.Blink = blinks[overview.Background][backgroundCode]
		background.Enable = isActive(bType, rawOverview.BackgroundStates)
		backgrounds = append(backgrounds, background)
	}

	flags := make([]FlagSet, 0, len(rawOverview.FlagOrder))
	for _, flagCode := range rawOverview.FlagOrder {
		if flagCode == unknownType {
			continue
		}

		fType, valid := validType[flagCode]
		if !valid {
			return nil, nil, fmt.Errorf("%d is not a valid flag code: %w", flagCode, ErrInvalidType)
		}

		flag := FlagSet{
			Type: fType,
		}

		colorCode, ok := colors[overview.Flag][flagCode]
		if !ok {
			flag.Color = Default
		} else {
			cType, valid := validColor[colorCode]
			if !valid {
				return nil, nil, fmt.Errorf("%s is not a valid color: %w", colorCode, ErrInvalidType)
			}
			flag.Color = cType
		}

		flag.Blink = blinks[overview.Flag][flagCode]
		flag.Enable = isActive(fType, rawOverview.FlagStates)
		flags = append(flags, flag)
	}

	return backgrounds, flags, nil
}

func isActive(target Type, types []uint8) bool {
	for _, t := range types {
		if target.Uint8() == t {
			return true
		}
	}
	return false
}

type Color uint8

const (
	White Color = iota
	Purple
	Orange
	Blue
	Turquoise
	DarkBlue
	Green
	Yellow
	DarkTurquoise
	Red
	Indigo
	Default
)

var validColor = map[string]Color{
	"white":         White,
	"purple":        Purple,
	"orange":        Orange,
	"blue":          Blue,
	"turquoise":     Turquoise,
	"darkBlue":      DarkBlue,
	"green":         Green,
	"yellow":        Yellow,
	"darkTurquoise": DarkTurquoise,
	"red":           Red,
	"indigo":        Indigo,
	"default":       Default,
}

func (c Color) IsDefault() bool {
	return c == Default
}

func (c Color) String() string {
	return [12]string{
		"white",
		"purple",
		"orange",
		"blue",
		"turquoise",
		"darkBlue",
		"green",
		"yellow",
		"darkTurquoise",
		"red",
		"indigo",
		"default",
	}[c]
}
