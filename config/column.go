package config

import (
	"errors"
	"fmt"
	"gihtub.com/evetools/overviewsdk/internal/overview"
)

type Column uint8

type ColumnSet struct {
	Type   Column
	Enable bool
}

const (
	Icon Column = iota
	Distance
	Name
	Type
	Tag
	Corporation
	Alliance
	Faction
	Militia
	Size
	Velocity
	RadialVelocity
	TransversalVelocity
	AngularVelocity
)

var validColumn = map[string]Column{
	"ICON":                Icon,
	"DISTANCE":            Distance,
	"NAME":                Name,
	"TYPE":                Type,
	"TAG":                 Tag,
	"CORPORATION":         Corporation,
	"ALLIANCE":            Alliance,
	"FACTION":             Faction,
	"MILITIA":             Militia,
	"SIZE":                Size,
	"VELOCITY":            Velocity,
	"RADIALVELOCITY":      RadialVelocity,
	"TRANSVERSALVELOCITY": TransversalVelocity,
	"ANGULARVELOCITY":     AngularVelocity,
}

func (c Column) String() string {
	return [14]string{
		"ICON",
		"DISTANCE",
		"NAME",
		"TYPE",
		"TAG",
		"CORPORATION",
		"ALLIANCE",
		"FACTION",
		"MILITIA",
		"SIZE",
		"VELOCITY",
		"RADIALVELOCITY",
		"TRANSVERSALVELOCITY",
		"ANGULARVELOCITY",
	}[c]
}

var ErrInvalidType = errors.New("invalid column type")

func MakeColumn(rawOverview *overview.RawOverview) ([]ColumnSet, error) {
	columns := make([]ColumnSet, 0, len(rawOverview.ColumnOrder))
	for _, columnCode := range rawOverview.ColumnOrder {
		cType, valid := validColumn[columnCode]
		if !valid {
			return nil, fmt.Errorf("%s is not a valid column code: %w", columnCode, ErrInvalidType)
		}
		column := ColumnSet{
			Type:   cType,
			Enable: isActive(cType, rawOverview.OverviewColumns),
		}
		columns = append(columns, column)
	}
	return columns, nil
}

func isActive(col Column, columns []string) bool {
	for _, c := range columns {
		if col.String() == c {
			return true
		}
	}
	return false
}
