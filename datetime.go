package datetime

import (
	"strings"
	"time"
)

var defaultTimeFormat = "2006-01-02 15:04:05"

// Datetime 日期时间类型
type Datetime time.Time

// New 创建一个当前时间的Datetime
func New() *Datetime {
	dt := Datetime(time.Now())
	return &dt
}

// SetTimeFormat 设置字符串格式
func SetTimeFormat(timeFormat string) {
	defaultTimeFormat = timeFormat
}

// Add 增加时间
func (dt *Datetime) Add(d time.Duration) {
	dt.SetTime(dt.Time().Add(d))
}

// UnmarshalJSON 反序列化处理
func (dt *Datetime) UnmarshalJSON(data []byte) error {
	str := strings.Replace(string(data), `"`, "", -1)
	return dt.SetString(str)
}

// MarshalJSON 序列化处理
func (dt *Datetime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

// String 获取日期时间，格式：yyyy-MM-dd HH:mm:ss
func (dt *Datetime) String() string {
	return formatTimeToString(dt.Time())
}

// SetString 设置日期时间，字符串格式：yyyy-MM-dd HH:mm:ss
func (dt *Datetime) SetString(str string) error {
	tt, err := parseStringToTime(str)
	if err != nil {
		return err
	}

	dt.SetTime(tt)
	return nil
}

// Time 获取 Datetime的 time.Time 类型形式
func (dt *Datetime) Time() time.Time {
	return time.Time(*dt)
}

// SetTime 用 time.Time 类型来设置 Datetime 的值
func (dt *Datetime) SetTime(tt time.Time) {
	*dt = Datetime(tt)
}

// Unix 获取 Datetime 的 Unix 时间戮，单位：秒
func (dt *Datetime) Unix() int64 {
	return dt.Time().Unix()
}

// SetUnix 用时间戮来设置 Datetime 的值
func (dt *Datetime) SetUnix(i64 int64) {
	dt.SetTime(time.Unix(i64, 0))
}

//将 yyyy-MM-dd HH:mm:ss 转换为 time.Time 类型
func parseStringToTime(str string) (time.Time, error) {
	var (
		loc *time.Location
		tt  time.Time
		err error
	)
	if loc, err = time.LoadLocation("Local"); err == nil {
		if tt, err = time.ParseInLocation(defaultTimeFormat, str, loc); err == nil {
			return tt, nil
		}
	}
	return time.Now(), err
}

//将 time.Time 类型转换为 yyyy-MM-dd HH:mm:ss
func formatTimeToString(tt time.Time) string {
	return tt.Format(defaultTimeFormat)
}
