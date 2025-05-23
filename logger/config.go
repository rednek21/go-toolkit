package logger

type Config struct {
	Level      string
	Format     string
	Output     string
	FilePath   string
	MaxSizeMB  int
	MaxBackups int
	MaxAgeDays int
	Compress   bool
}
