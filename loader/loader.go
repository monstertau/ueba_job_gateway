package loader

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"job-gateway/kdb"
	"os"
	"path"
	"strings"
)

const (
	ProfileContentDir  = "behavior_profile"
	BehaviorContentDir = "behavior"
	SourceContentDir   = "source"
	UseCaseContentDir  = "rule"
)

type ContentPackLoader struct {
	path string
	cp   *kdb.ContentPack
}

func New(path string) *ContentPackLoader {
	return &ContentPackLoader{path: path, cp: &kdb.ContentPack{
		Sources:   make(map[string]*kdb.SourceKDB),
		Behaviors: make(map[string]*kdb.BehaviorKDB),
		Rules:     make(map[string]*kdb.RuleKDB),
		Profiles:  make(map[string]*kdb.ProfileKDB),
	}}
}

func (l *ContentPackLoader) Load() error {
	entries, err := os.ReadDir(l.path)
	if err != nil {
		logrus.Errorf("error in read content pack: %v", err)
		return err
	}
	for _, dirEnt := range entries {
		switch dirEnt.Name() {
		case SourceContentDir:
			err = l.loadMultiSourceContent(l.path, dirEnt)
		case BehaviorContentDir:
			err = l.loadMultiBehaviorContent(l.path, dirEnt)
		case UseCaseContentDir:
			err = l.loadMultiUseCaseContent(l.path, dirEnt)
		case ProfileContentDir:
			err = l.loadMultiProfileContent(l.path, dirEnt)
		default:
			err = fmt.Errorf("suspicious file in content pack: %v", dirEnt.Name())
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *ContentPackLoader) loadMultiSourceContent(dirPath string, entry os.DirEntry) error {
	folderPath := path.Join(dirPath, entry.Name())
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}
	for _, dirEnt := range entries {
		if !strings.Contains(dirEnt.Name(), ".yml") {
			return fmt.Errorf("suspicious file in content pack: %v", dirEnt.Name())
		}
		src, err := loadSourceContent(path.Join(folderPath, dirEnt.Name()))
		if err != nil {
			return err
		}
		l.cp.Sources[src.ID] = src
	}
	return nil
}

func (l *ContentPackLoader) loadMultiBehaviorContent(dirPath string, entry os.DirEntry) error {
	folderPath := path.Join(dirPath, entry.Name())
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}
	for _, dirEnt := range entries {
		if !strings.Contains(dirEnt.Name(), ".yml") {
			return fmt.Errorf("suspicious file in content pack: %v", dirEnt.Name())
		}
		bhv, err := loadBehaviorContent(path.Join(folderPath, dirEnt.Name()))
		if err != nil {
			return err
		}
		l.cp.Behaviors[bhv.ID] = bhv
	}
	return nil
}

func (l *ContentPackLoader) loadMultiUseCaseContent(dirPath string, entry os.DirEntry) error {
	folderPath := path.Join(dirPath, entry.Name())
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}
	for _, dirEnt := range entries {
		if !strings.Contains(dirEnt.Name(), ".yml") {
			return fmt.Errorf("suspicious file in content pack: %v", dirEnt.Name())
		}
		uc, err := loadUseCaseContent(path.Join(folderPath, dirEnt.Name()))
		if err != nil {
			return err
		}
		l.cp.Rules[uc.ID] = uc
	}
	return nil
}

func (l *ContentPackLoader) GetContentPack() *kdb.ContentPack {
	return l.cp
}

func loadSourceContent(fp string) (*kdb.SourceKDB, error) {
	var source *kdb.SourceKDB
	yamlFile, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &source)
	if err != nil {
		return nil, err
	}
	return source, nil
}

func loadBehaviorContent(fp string) (*kdb.BehaviorKDB, error) {
	var bhv *kdb.BehaviorKDB
	yamlFile, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &bhv)
	if err != nil {
		return nil, err
	}
	return bhv, nil
}
func loadUseCaseContent(fp string) (*kdb.RuleKDB, error) {
	var uc *kdb.RuleKDB
	yamlFile, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &uc)
	if err != nil {
		return nil, err
	}
	return uc, nil
}

func (l *ContentPackLoader) loadMultiProfileContent(dirPath string, entry os.DirEntry) error {
	folderPath := path.Join(dirPath, entry.Name())
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}
	for _, dirEnt := range entries {
		if !strings.Contains(dirEnt.Name(), ".yml") {
			return fmt.Errorf("suspicious file in content pack: %v", dirEnt.Name())
		}
		profile, err := loadProfileContent(path.Join(folderPath, dirEnt.Name()))
		if err != nil {
			return err
		}
		l.cp.Profiles[profile.ID] = profile
	}
	return nil
}

func loadProfileContent(fp string) (*kdb.ProfileKDB, error) {
	var bhv *kdb.ProfileKDB
	yamlFile, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &bhv)
	if err != nil {
		return nil, err
	}
	return bhv, nil
}
