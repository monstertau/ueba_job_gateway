package service

import (
	"encoding/json"
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

func (j *ProfilingJobService) GetJobs() ([]*view.ProfilingJobConfig, error) {
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

func NewFlinkSQLJobService() *FlinkSQLJobService {
	return &FlinkSQLJobService{}
}

func (j *FlinkSQLJobService) GetBehaviorJobs() ([]*view.BehaviorJobConfig, error) {
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

func (j *FlinkSQLJobService) GetRuleJobs() ([]*view.RuleJobConfig, error) {
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
