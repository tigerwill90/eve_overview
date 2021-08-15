package overviewsdk

import (
	"gihtub.com/evetools/overviewsdk/config"
	"gihtub.com/evetools/overviewsdk/config/appearance"
	"gihtub.com/evetools/overviewsdk/config/preset"
)

type Overview struct {
	Backgrounds []appearance.BackgroundSet
	Flags       []appearance.FlagSet
	Columns     []config.ColumnSet
	Presets     map[preset.Name]preset.Item
}
