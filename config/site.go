package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at"`
	ICP         string `yaml:"icp" json:"icp"`
	Title       string `yaml:"title" json:"title"`
	QQImage     string `yaml:"qq_image" json:"qq_image"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image"`
	Version     string `yaml:"version" json:"version"`
	Email       string `yaml:"email" json:"email"`
	Phone       string `yaml:"phone" json:"phone"`
	Name        string `yaml:"name" json:"name"`
	Job         string `yaml:"job" json:"job"`
	Addr        string `yaml:"addr" json:"addr"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	SloganEn    string `yaml:"slogan_en" json:"slogan_en"`
	Web         string `yaml:"web" json:"web"`
	GithubUrl   string `yaml:"github_url" json:"github_url"`
}
