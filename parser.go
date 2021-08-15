package overviewsdk

import (
	"gihtub.com/evetools/overviewsdk/globalconfig/appearance"
	"gihtub.com/evetools/overviewsdk/internal/overview"
	"io"
)

func Unmarshal(r io.Reader) (*Overview, error) {
	rawOverview, err := overview.RawUnmarshal(r)
	if err != nil {
		return nil, err
	}
	backgrounds, flags, err := appearance.Make(rawOverview)
	if err != nil {
		return nil, err
	}

	return &Overview{
		Backgrounds: backgrounds,
		Flags:       flags,
	}, nil
}
