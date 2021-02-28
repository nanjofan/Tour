package timer

import "time"

func GetNowTime() time.Time {
	location,_ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// 在不知道d具体为什么格式的情况下，使用此函数
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// ParseDuration 从字符串d中解析出duration，d的值可以为：such as "300ms", "-1.5h" or "2h45m".
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
