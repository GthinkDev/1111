package flag

import (
	"fmt"
	"gin-blog/global"
)

func Version() {
	v := global.Config.SiteInfo.Version
	fmt.Printf("当前版本号: %s\n", v)
}
