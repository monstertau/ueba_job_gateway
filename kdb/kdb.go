package kdb

import "job-gateway/view"

type (
	ContentPack struct {
		Sources   map[string]*SourceKDB
		Behaviors map[string]*BehaviorKDB
		Profiles  map[string]*ProfileKDB
		Rules     map[string]*RuleKDB
	}
	SourceKDB struct {
		ID     string            `yaml:"id"`
		Name   string            `yaml:"name"`
		Config map[string]string `yaml:"config"`
	}
	BehaviorKDB struct {
		ID                string            `yaml:"id"`
		Name              string            `yaml:"name"`
		MitreDataSourceID string            `yaml:"mitre_data_source_id"`
		SourceID          string            `yaml:"source_id"`
		Filter            string            `yaml:"filter"`
		Schema            map[string]string `yaml:"schema"`
		TimestampField    string            `yaml:"timestamp_field"`
	}
	ProfileKDB struct {
		ID                   string         `yaml:"id"`
		Name                 string         `yaml:"name"`
		BehaviorID           string         `yaml:"behavior_id"`
		Status               int64          `yaml:"status"`
		ProfileType          string         `yaml:"profile_type"`
		Entities             []*view.Object `yaml:"entities"`
		Attributes           []*view.Object `yaml:"attributes"`
		ProfileTime          string         `yaml:"profile_time"`
		SavingDurationMinute int            `yaml:"saving_duration_minute"`
		Threshold            float64        `yaml:"threshold"`
	}
	RuleKDB struct {
		ID        string `yaml:"id"`
		Name      string `yaml:"name"`
		Technique string `yaml:"technique"`
		ProfileID string `yaml:"profile_id"`
		Filter    string `yaml:"filter"`
		Severity  string `yaml:"severity"`
		RiskScore int    `yaml:"risk_score"`
		Object    string `yaml:"object"`
	}
)
