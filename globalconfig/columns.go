package globalconfig

type Column uint8

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
