package otherFunc

import (
	"fmt"
)

// Configrules define the rules for the arguments
func Configrules() {
	configuration := `	有3种信息需要确认格式规范: 
		1) -e/--email: 邮箱地址满足 xxx@xxx.com 的规范。
				eg: sample@qq.com(right) sample@qq..com(wrong)
		2) -t/--telephone: 电话号码由11位数字组成，而且必须以1开头。
				eg: 12345678901(right) 23456789014(wrong)
		3) --startTime/--endTime: 会议起始时间和结束时间满足 YYYY-MM-DD HH:MM 的规范。
				eg: 2014-11-12 15:30(right) 2014-9-15 16:20(wrong)
	`
	fmt.Println(configuration)
}
