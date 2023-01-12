package internal

type Config struct {
	BackupPath        string `yaml:"backupPath"`
	UpdatePluginsPath string `yaml:"updatePluginsPath"`
	OldPluginsPath    string `yaml:"oldPluginsPath"`
}
