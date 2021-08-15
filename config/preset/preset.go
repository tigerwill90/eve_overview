package preset

import (
	"errors"
	"fmt"
	"gihtub.com/evetools/overviewsdk/internal/overview"
	"golang.org/x/net/html"
	"image/color"
	"io"
	"strings"
)

type Color uint8

const (
	White Color = iota
	Silver
	Gray
	Black
	Red
	Maroon
	Yellow
	Olive
	Lime
	Green
	Aqua
	Teal
	Blue
	Navy
	Fuchsia
	Purple
)

func (c Color) String() string {
	return [16]string{"white", "silver", "gray", "black", "red", "maroon", "yellow", "olive", "lime", "green", "aqua", "teal", "blue", "navy", "fuchsia", "purple"}[c]
}

var validHtmlColors = map[string]color.RGBA{
	"black":   {0, 0, 0, 255},
	"white":   {255, 255, 255, 255},
	"gray":    {128, 128, 128, 255},
	"silver":  {192, 192, 192, 255},
	"maroon":  {128, 0, 0, 255},
	"red":     {255, 0, 0, 255},
	"purple":  {128, 0, 128, 255},
	"fushsia": {255, 0, 255, 255},
	"green":   {0, 128, 0, 255},
	"lime":    {0, 255, 0, 255},
	"olive":   {128, 128, 0, 255},
	"yellow":  {255, 255, 0, 255},
	"navy":    {0, 0, 128, 255},
	"blue":    {0, 0, 255, 255},
	"teal":    {0, 128, 128, 255},
	"aqua":    {0, 255, 255, 255},
}

type Name struct {
	text     string
	raw      string
	color    color.RGBA
	hasColor bool
	prefix   string
	indent   int
}

func (n Name) HexColor() string {
	if !n.hasColor {
		return ""
	}
	return fmt.Sprintf("0xff%02x%02x%02x", n.color.R, n.color.G, n.color.B)
}

type Item struct {
	AlwaysShownStates []int `yaml:"alwaysShownStates"`
	FilteredStates    []int `yaml:"filteredStates"`
	Groups            []int `yaml:"groups"`
}

var (
	ErrInvalidColorTag = errors.New("invalid color tag")
	ErrInvalidColor    = errors.New("invalid color")
)

func ParseName(str string) (Name, error) {
	name := Name{
		raw: str,
	}
	z := html.NewTokenizer(strings.NewReader(str))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				break
			}
			return Name{}, z.Err()
		}
		tag, _ := z.TagName()
		if strings.Contains(string(tag), "color=") {
			split := strings.SplitN(string(tag), "=", 2)
			if len(split) != 2 {
				return Name{}, ErrInvalidColorTag
			}
			c, err := ParseHexColor(split[1])
			if err != nil {
				return Name{}, err
			}
			name.color = c
			name.hasColor = true
		}
		text := string(z.Text())
		if text != "" {
			name.text = text
			name.indent = countLeadingSpaces(text)
		}
	}
	return name, nil
}

func countLeadingSpaces(line string) int {
	count := 0
	for _, v := range line {
		if v == ' ' {
			count++
		} else {
			break
		}
	}

	return count
}

var ErrInvalidHexFormat = errors.New("invalid hex format")

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	// Is a base html color
	if rgba, valid := validHtmlColors[s]; valid {
		return rgba, nil
	}

	// Parse hex format
	if strings.HasPrefix(s, "0xff") {
		s = strings.Replace(s, "0xff", "#", 1)
	} else if s[0] != '#' {
		return c, ErrInvalidHexFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = ErrInvalidHexFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = ErrInvalidHexFormat
	}
	return
}

func Make(rawOverview *overview.RawOverview) ([]Item, error) {
	presets, err := rawOverview.ParsePresets()
	if err != nil {
		return nil, err
	}

	for name := range presets {
		n, err := ParseName(name)
		if err != nil {
			return nil, err
		}
		fmt.Println(n.raw)
		fmt.Println(n.color, "=>", n.HexColor())
	}
	return nil, nil
}
