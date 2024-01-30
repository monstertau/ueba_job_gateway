package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"job-gateway/importer"
	"job-gateway/kdb"
	"job-gateway/view"
	"os"
)

type (
	Services struct {
		Profiling *ProfilingJobService
		FlinkSQL  *FlinkSQLJobService
	}
	ProfilingJobService struct {
	}
	FlinkSQLJobService struct {
	}
)

func NewProfilingJobService() *ProfilingJobService {
	return &ProfilingJobService{}
}

func (j *ProfilingJobService) GetMockJobs() ([]*view.ProfilingJobConfig, error) {
	// Read the JSON file
	data, err := os.ReadFile("./mock_data/profiling_job.json")
	if err != nil {
		return nil, err
	}
	var result []*view.ProfilingJobConfig

	// Unmarshal the JSON data into the result slice
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (j *ProfilingJobService) GetJobs() ([]*view.ProfilingJobConfig, error) {
	var result []*view.ProfilingJobConfig
	cp := importer.GlobalImporter.GetContent()
	for _, p := range cp.Profiles {
		_, ok := cp.Behaviors[p.BehaviorID]
		if !ok {
			return nil, errors.Errorf("cant find behavior with id: %v", p.BehaviorID)
		}
		cfg := &view.ProfilingJobConfig{
			ID: p.ID,
			ProfileConfig: &view.ProfileConfig{
				Name:           p.Name,
				Status:         p.Status,
				ProfileType:    p.ProfileType,
				Entities:       p.Entities,
				Attributes:     p.Attributes,
				ProfileTime:    p.ProfileTime,
				SavingDuration: p.SavingDurationMinute,
				Threshold:      p.Threshold,
			},
			LogSourceBuilderConfig: &view.KafkaConfig{
				BootstrapServer: "localhost:9092",
				Topic:           fmt.Sprintf("profiling_sink_%v", p.ID),
			},
			LogSourcePredictorConfig: &view.KafkaConfig{
				BootstrapServer: "localhost:9092",
				Topic:           fmt.Sprintf("behavior_sink_%v", p.BehaviorID),
			},
			OutputConfig: &view.OutputConfig{
				Type: "kafka",
				Config: map[string]interface{}{
					"bootstrap.servers": "localhost:9092",
					"topic":             fmt.Sprintf("profiling_predictor_%v", p.ID),
				},
			},
		}
		result = append(result, cfg)
	}
	return result, nil
}

func NewFlinkSQLJobService() *FlinkSQLJobService {
	return &FlinkSQLJobService{}
}

func (j *FlinkSQLJobService) GetMockBehaviorJobs() ([]*view.BehaviorJobConfig, error) {
	// Read the JSON file
	data, err := os.ReadFile("./mock_data/behavior_job.json")
	if err != nil {
		return nil, err
	}
	var result []*view.BehaviorJobConfig

	// Unmarshal the JSON data into the result slice
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (j *FlinkSQLJobService) GetBehaviorJobs() ([]*view.BehaviorJobConfig, error) {
	var result []*view.BehaviorJobConfig
	cp := importer.GlobalImporter.GetContent()
	for _, bhv := range cp.Behaviors {
		var profile *kdb.ProfileKDB
		for _, p := range cp.Profiles {
			if p.BehaviorID == bhv.ID {
				profile = p
				break
			}
		}
		src, ok := cp.Sources[bhv.SourceID]
		if !ok {
			return nil, errors.Errorf("cant find source with id: %v", bhv.SourceID)
		}
		cfg := &view.BehaviorJobConfig{
			ID: bhv.ID,
			LogSourceConfig: &view.KafkaConfig{
				BootstrapServer: src.Config["bootstrap.servers"],
				Topic:           src.Config["topic"],
				Schema:          bhv.Schema,
				TimestampField:  bhv.TimestampField,
			},
			ProfileConfig: &view.ProfileConfig{
				ID:             profile.ID,
				Name:           profile.Name,
				Status:         profile.Status,
				ProfileType:    profile.ProfileType,
				Entities:       profile.Entities,
				Attributes:     profile.Attributes,
				ProfileTime:    profile.ProfileTime,
				SavingDuration: profile.SavingDurationMinute,
				Threshold:      profile.Threshold,
			},
			ProfileOutput: &view.KafkaConfig{
				BootstrapServer: "helk-kafka-broker:9092",
				Topic:           fmt.Sprintf("profiling_sink_%v", profile.ID),
			},
			BehaviorOutput: &view.KafkaConfig{
				BootstrapServer: "helk-kafka-broker:9092",
				Topic:           fmt.Sprintf("behavior_sink_%v", bhv.ID),
			},
			BehaviorFilter: bhv.Filter,
		}
		result = append(result, cfg)
	}

	return result, nil
}

func (j *FlinkSQLJobService) GetRuleJobs() ([]*view.RuleJobConfig, error) {
	var result []*view.RuleJobConfig
	cp := importer.GlobalImporter.GetContent()
	for _, rule := range cp.Rules {
		p, ok := cp.Profiles[rule.ProfileID]
		if !ok {
			return nil, errors.Errorf("cant find profile with id: %v", rule.ProfileID)
		}
		bhv, ok := cp.Behaviors[p.BehaviorID]
		if !ok {
			return nil, errors.Errorf("cant find behavior with id: %v", p.BehaviorID)
		}
		cfg := &view.RuleJobConfig{
			ID:        rule.ID,
			Name:      rule.Name,
			Filter:    rule.Filter,
			Object:    rule.Object,
			Technique: rule.Technique,
			Severity:  rule.Severity,
			RiskScore: rule.RiskScore,
			ProfilePredictorOutput: &view.KafkaConfig{
				BootstrapServer: "helk-kafka-broker:9092",
				Topic:           fmt.Sprintf("profiling_predictor_%v", p.ID),
				Schema:          customizeRuleSchema(bhv.Schema, bhv.TimestampField),
				TimestampField:  bhv.TimestampField,
			},
			RuleOutput: &view.KafkaConfig{
				BootstrapServer: "helk-kafka-broker:9092",
				Topic:           "alert_sink",
			},
		}
		result = append(result, cfg)
	}
	return result, nil
}

func customizeRuleSchema(orig map[string]string, timestampField string) map[string]string {
	schema := make(map[string]string)
	for k, v := range orig {
		schema[k] = v
	}
	schema[timestampField] = "STRING"
	schema["profile_predictor_entities"] = "STRING"
	schema["profile_predictor_attributes"] = "STRING"
	schema["profile_predictor_result"] = "FLOAT"
	schema["profile_predictor_threshold"] = "FLOAT"
	return schema
}

func (j *FlinkSQLJobService) GetMockRuleJobs() ([]*view.RuleJobConfig, error) {
	// Read the JSON file
	data, err := os.ReadFile("./mock_data/rule_job.json")
	if err != nil {
		return nil, err
	}
	var result []*view.RuleJobConfig

	// Unmarshal the JSON data into the result slice
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
