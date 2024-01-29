package view

type (
	BehaviorJobConfig struct {
		ID              string         `json:"id"`
		LogSourceConfig *kafkaConfig   `json:"source_config" binding:"required"`
		ProfileConfig   *ProfileConfig `json:"profile_config" binding:"required"`
		ProfileOutput   *kafkaConfig   `json:"profile_output_config" binding:"required"`
		BehaviorOutput  *kafkaConfig   `json:"behavior_output_config" binding:"required"`
		BehaviorFilter  string         `json:"filter" binding:"required"`
	}
	RuleJobConfig struct {
		ID                     string       `json:"id"`
		Filter                 string       `json:"filter"`
		ProfilePredictorOutput *kafkaConfig `json:"profile_predictor_config" binding:"required"`
		RuleOutput             *kafkaConfig `json:"rule_output_config" binding:"required"`
	}
)
