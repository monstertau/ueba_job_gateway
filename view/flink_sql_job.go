package view

type (
	BehaviorJobConfig struct {
		ID              string         `json:"id"`
		LogSourceConfig *KafkaConfig   `json:"source_config" binding:"required"`
		ProfileConfig   *ProfileConfig `json:"profile_config" binding:"required"`
		ProfileOutput   *KafkaConfig   `json:"profile_output_config" binding:"required"`
		BehaviorOutput  *KafkaConfig   `json:"behavior_output_config" binding:"required"`
		BehaviorFilter  string         `json:"filter" binding:"required"`
	}
	RuleJobConfig struct {
		ID                     string       `json:"id"`
		Name                   string       `json:"name"`
		Filter                 string       `json:"filter"`
		Object                 string       `json:"object"`
		Technique              string       `json:"technique"`
		Severity               string       `json:"severity"`
		RiskScore              int          `json:"risk_score"`
		ProfilePredictorOutput *KafkaConfig `json:"profile_predictor_config" binding:"required"`
		RuleOutput             *KafkaConfig `json:"rule_output_config" binding:"required"`
	}
)
