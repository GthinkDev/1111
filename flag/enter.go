package flag

import SysFlag "flag"

type Option struct {
	Version bool
	DB      bool
}

// Parse 全局命令行参数注册器
func Parse() Option {
	version := SysFlag.Bool("version", true, "版本号")
	db := SysFlag.Bool("db", true, "初始化数据库")
	// 解析命令行参数写入注册的 flag 里
	SysFlag.Parse()
	return Option{
		Version: *version,
		DB:      *db,
	}
}

// IsWebStop 判断是否需要执行web服务
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	if option.Version {
		return true
	}
	return false
}

// SwitchOption 判断是否需要执行web服务
func SwitchOption(option Option) {
	if option.DB {
		MakeMigration()
	}
	if option.Version {
		Version()
	}
}
