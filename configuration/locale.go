package configuration

import (
	"context"
	"distudio.com/page"
	"encoding/json"
	"golang.org/x/text/language"
)

type Locale language.Tag

func (locale *Locale) MarshalJSON() ([]byte, error) {
	tag := language.Tag(*locale)
	return json.Marshal(tag.String())
}

func (locale *Locale) Id() string {
	return language.Tag(*locale).String()
}

func (locale *Locale) FromRepresentation(rtype page.RepresentationType, data []byte) error {
	return page.NewUnsupportedError()
}

func (locale *Locale) ToRepresentation(rtype page.RepresentationType) ([]byte, error) {
	switch rtype {
	case page.RepresentationTypeJSON:
		return json.Marshal(locale)
	}
	return nil, page.NewUnsupportedError()
}

func NewLocaleController() *page.RestController {
	man := localeManager{}
	return page.NewRestController(page.BaseRestHandler{Manager: man})
}

type localeManager struct{}

func (manager localeManager) NewResource(ctx context.Context) (page.Resource, error) {
	return nil, page.NewUnsupportedError()
}

func (manager localeManager) FromId(ctx context.Context, id string) (page.Resource, error) {
	return nil, page.NewUnsupportedError()
}

func (manager localeManager) ListOf(ctx context.Context, opts page.ListOptions) ([]page.Resource, error) {

	ws := page.Application()

	langs := ws.Options().Languages

	from := opts.Page * opts.Size
	if from > len(langs) {
		return make([]page.Resource, 0), nil
	}

	to := from + opts.Size
	if to > len(langs) {
		to = len(langs)
	}

	items := langs[from:to]
	resources := make([]page.Resource, len(items))

	for i := range items {
		locale := Locale(items[i])
		resources[i] = page.Resource(&locale)
	}

	return resources, nil
}

func (manager localeManager) ListOfProperties(ctx context.Context, opts page.ListOptions) ([]string, error) {
	return nil, page.NewUnsupportedError()
}

func (manager localeManager) Create(ctx context.Context, res page.Resource, bundle []byte) error {
	return page.NewUnsupportedError()
}

func (manager localeManager) Update(ctx context.Context, res page.Resource, bundle []byte) error {
	return page.NewUnsupportedError()
}

func (manager localeManager) Delete(ctx context.Context, res page.Resource) error {
	return page.NewUnsupportedError()
}