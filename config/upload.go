package config

type Upload struct {
	Size     int64  `json:"size" yaml:"size"`
	SavePath string `yaml:"save_path" json:"save_path"`
	//AllowExt []string `json:"allow_ext" yaml:"allow_ext"`
}
