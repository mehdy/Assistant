package feed

import (
	"context"
	"fmt"
	"time"

	"github.com/mehdy/Assistant/pkg/apis"
	"github.com/mmcdole/gofeed"
)

type Watcher struct {
	el          apis.EventLoop
	url         string
	fp          *gofeed.Parser
	lastChecked time.Time
}

func NewWatcher(el apis.EventLoop, url string) *Watcher {
	return &Watcher{el: el, url: url, fp: gofeed.NewParser()}
}

func (w *Watcher) update() {
	feed, err := w.fp.ParseURLWithContext(w.url, context.Background())
	if err != nil {
		fmt.Printf("Failed to fetch and parse %s: %s\n", w.url, err)

		return
	}

	for _, item := range feed.Items {
		// FIXME: This doesn't work correctly necessarily
		if item.PublishedParsed != nil && item.PublishedParsed.After(w.lastChecked) {
			w.el.Emit(apis.NewEvent("Feed::Watcher::Update", item))
		}
	}

	w.lastChecked = time.Now()
}
func (w *Watcher) Run() {
	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {
		w.update()
	}
}
