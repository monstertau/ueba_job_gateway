package view

type (
	ProfileConfig struct {
		Name           string    `json:"name"`
		Status         int64     `json:"status"`
		ProfileType    string    `json:"profile_type"`
		Entities       []*Object `json:"entity"`
		Attributes     []*Object `json:"attribute"`
		ProfileTime    string    `json:"profile_time"`
		SavingDuration int       `json:"saving_duration_minute"`
		Threshold      float64   `json:"threshold"`
	}
	Object struct {
		Name      string                 `json:"field_name" example:"server_id"`
		Type      string                 `json:"type" enums:"original,mapping,reference" example:"original"`
		ExtraData map[string]interface{} `json:"extra_data,omitempty"`
	}
	kafkaConfig struct {
		BootstrapServer string            `json:"bootstrap.servers" binding:"required"`
		Topic           string            `json:"topic,omitempty"`
		AuthenType      string            `json:"authen_type,omitempty"`
		Keytab          string            `json:"keytab,omitempty"`
		Principal       string            `json:"principal,omitempty"`
		Schema          map[string]string `json:"schema,omitempty"`
		TimestampField  string            `json:"timestamp_field,omitempty"`
	}
	LogSourceConfig struct {
		Config kafkaConfig `json:"config" binding:"required"`
	}
	ProfilingJobConfig struct {
		ID                       string         `json:"id"`
		ProfileConfig            *ProfileConfig `json:"profile_config" binding:"required"`
		LogSourceBuilderConfig   *kafkaConfig   `json:"builder_source_config" binding:"required"`
		LogSourcePredictorConfig *kafkaConfig   `json:"predictor_source_config" binding:"required"`
		OutputConfig             *OutputConfig  `json:"output_config" binding:"required"`
	}
	OutputConfig struct {
		ExpirationTime string                 `mapstructure:"expiration_time" json:"expiration_time"`
		Type           string                 `mapstructure:"type" json:"type"`
		Config         map[string]interface{} `mapstructure:"config" json:"config"`
		Workers        int                    `mapstructure:"workers" json:"workers"`
	}
)
