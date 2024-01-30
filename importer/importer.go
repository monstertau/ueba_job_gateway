package importer

import (
	"github.com/sirupsen/logrus"
	"job-gateway/kdb"
	"job-gateway/loader"
	"time"
)

type Importer struct {
	path   string
	cp     *kdb.ContentPack
	logger *logrus.Entry
}

var GlobalImporter *Importer

func NewImporter(path string) (*Importer, error) {
	i := &Importer{
		path:   path,
		logger: logrus.WithField("service", "importer"),
	}
	err := i.load()
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *Importer) Sync() {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			err := i.load()
			if err != nil {
				i.logger.Errorf("cant load content pack: %v", err)
			}
		}
	}
}

func (i *Importer) GetContent() *kdb.ContentPack {
	return i.cp
}

func (i *Importer) load() error {
	load := loader.New(i.path)
	err := load.Load()
	if err != nil {
		return err
	}
	i.cp = load.GetContentPack()
	return nil
}
