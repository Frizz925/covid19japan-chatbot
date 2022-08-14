package covid19goid

import (
	"context"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/id/covid19goid"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	idFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/id"
)

const FIXTURE_FILE_UPDATE = "update.json"

type FixtureFetcher struct {
	fetcher.FixtureFetcher
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{
		FixtureFetcher: fetcher.FixtureFetcher{
			Directory:  dir,
			Country:    country.ID,
			SourceName: idFetcher.DATA_SOURCE_COVID19_GO_ID,
		},
	}
}

func (f *FixtureFetcher) Source() string {
	return f.GetPath("")
}

func (f *FixtureFetcher) Update(ctx context.Context) (*covid19goid.UpdateResponse, error) {
	u, err := f.GetFullURL(FIXTURE_FILE_UPDATE)
	if err != nil {
		return nil, err
	}
	rc, err := f.ReadFile(u.Path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return covid19goid.ParseUpdate(rc, data.Source{
		URL:     SOURCE_URL,
		DataURL: u.String(),
		Comment: SOURCE_COMMENT + " " + fetcher.SOURCE_COMMENT_FIXTURE_SUFFIX,
	})
}

func (f *FixtureFetcher) DailySummary(ctx context.Context) (*data.DailySummary, error) {
	ur, err := f.Update(ctx)
	if err != nil {
		return nil, err
	}
	return ur.Normalize()
}
