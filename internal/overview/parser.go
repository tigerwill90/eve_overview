package overview

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"strconv"
	"strings"
)

type RawOverview struct {
	BackgroundOrder     []uint8         `yaml:"backgroundOrder"`
	BackgroundStates    []uint8         `yaml:"backgroundStates"`
	ColumnOrder         []string        `yaml:"columnOrder"`
	FlagOrder           []uint8         `yaml:"flagOrder"`
	FlagStates          []uint8         `yaml:"flagStates"`
	OverviewColumns     []string        `yaml:"overviewColumns"`
	StateBlinks         [][]interface{} `yaml:"stateBlinks"`
	StateColorsNameList [][]string      `yaml:"stateColorsNameList"`
	Presets             [][]interface{} `yaml:"presets"`
}

func RawUnmarshal(r io.Reader) (*RawOverview, error) {
	row := new(RawOverview)
	if err := yaml.NewDecoder(r).Decode(row); err != nil {
		return nil, err
	}
	return row, nil
}

var (
	ErrNoColor       = errors.New("no color found")
	ErrInvalidColor  = errors.New("invalid color")
	ErrNoBlinks      = errors.New("no blink state found")
	ErrInvalidBlinks = errors.New("invalid blink state")
)

type ColorType uint8

const (
	Background ColorType = iota
	Flag
)

func (c ColorType) String() string {
	return [...]string{"background", "flag"}[c]
}

const (
	backgroundKey = "background_"
	flagKey       = "flag_"
)

func (o *RawOverview) ParseBlink() (map[ColorType]map[uint8]bool, error) {
	if len(o.StateBlinks) == 0 {
		return nil, ErrNoBlinks
	}
	blinks := map[ColorType]map[uint8]bool{
		Background: make(map[uint8]bool),
		Flag:       make(map[uint8]bool),
	}

	for _, blink := range o.StateBlinks {
		if len(blink) != 2 {
			return nil, ErrInvalidBlinks
		}
		blinkString, ok := blink[0].(string)
		if !ok {
			return nil, ErrInvalidBlinks
		}
		blinkState, ok := blink[1].(bool)
		if !ok {
			return nil, ErrInvalidBlinks
		}
		if strings.Contains(blinkString, backgroundKey) {
			n, err := parseBlinks(blinkString, backgroundKey)
			if err != nil {
				return nil, err
			}
			blinks[Background][n] = blinkState
		} else if strings.Contains(blinkString, flagKey) {
			n, err := parseBlinks(blinkString, flagKey)
			if err != nil {
				return nil, err
			}
			blinks[Flag][n] = blinkState
		} else {
			return nil, ErrInvalidBlinks
		}
	}
	return blinks, nil
}

func parseBlinks(s, replaceKey string) (uint8, error) {
	b := strings.Replace(s, replaceKey, "", 1)
	n, err := strconv.ParseUint(b, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", err, ErrInvalidBlinks)
	}
	return uint8(n), nil
}

func (o *RawOverview) ParseColors() (map[ColorType]map[uint8]string, error) {
	if len(o.StateColorsNameList) == 0 {
		return nil, ErrNoColor
	}
	colors := map[ColorType]map[uint8]string{
		Background: make(map[uint8]string),
		Flag:       make(map[uint8]string),
	}
	for _, color := range o.StateColorsNameList {
		if len(color) != 2 {
			return nil, ErrInvalidColor
		}
		if strings.Contains(color[0], backgroundKey) {
			n, c, err := parseColor(color, backgroundKey)
			if err != nil {
				return nil, err
			}
			colors[Background][n] = c
		} else if strings.Contains(color[0], flagKey) {
			n, c, err := parseColor(color, flagKey)
			if err != nil {
				return nil, err
			}
			colors[Flag][n] = c
		} else {
			return nil, ErrInvalidColor
		}
	}
	return colors, nil
}

func parseColor(color []string, replaceKey string) (uint8, string, error) {
	c := strings.Replace(color[0], replaceKey, "", 1)
	n, err := strconv.ParseUint(c, 10, 8)
	if err != nil {
		return 0, "", fmt.Errorf("%s: %w", err, ErrInvalidColor)
	}
	return uint8(n), color[1], nil
}
