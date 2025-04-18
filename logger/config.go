package logger

type Config struct {
	Level      string `yaml:"level" json:"level"`
	Format     string `yaml:"format" json:"format"`
	Output     string `yaml:"output" json:"output"`
	FilePath   string `yaml:"file_path" json:"file_path"`
	MaxSizeMB  int    `yaml:"max_size_mb" json:"max_size"`
	MaxBackups int    `yaml:"max_backups" json:"max_backups"`
	MaxAgeDays int    `yaml:"max_age_days" json:"max_age"`
	Compress   bool   `yaml:"compress" json:"compress"`
}
