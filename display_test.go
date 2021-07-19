package carbon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_ToTimestamp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimestamp(), test.output)
	}
}

func TestCarbon_ToTimestampWithSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimestampWithSecond(), test.output)
	}
}

func TestCarbon_ToTimestampWithMillisecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655000},
		{"2020-01-31 13:14:15", 1580447655000},
		{"2020-02-01 13:14:15", 1580534055000},
		{"2020-02-28 13:14:15", 1582866855000},
		{"2020-02-29 13:14:15", 1582953255000},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimestampWithMillisecond(), test.output)
	}
}

func TestCarbon_ToTimestampWithMicrosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655000000},
		{"2020-01-31 13:14:15", 1580447655000000},
		{"2020-02-01 13:14:15", 1580534055000000},
		{"2020-02-28 13:14:15", 1582866855000000},
		{"2020-02-29 13:14:15", 1582953255000000},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimestampWithMicrosecond(), test.output)
	}
}

func TestCarbon_ToTimestampWithNanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655000000000},
		{"2020-01-31 13:14:15", 1580447655000000000},
		{"2020-02-01 13:14:15", 1580534055000000000},
		{"2020-02-28 13:14:15", 1582866855000000000},
		{"2020-02-29 13:14:15", 1582953255000000000},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimestampWithNanosecond(), test.output)
	}
}

func TestCarbon_String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(fmt.Sprintf("%s", Parse(test.input)), test.output)
	}
}

func TestCarbon_ToString(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToString(), test.output)
	}
}

func TestCarbon_ToUtcString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 05:14:15 +0000 UTC"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToUtcString(), test.output)
	}
}

func TestCarbon_ToMonthString1(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "January"},
		{"2020-02-05", "February"},
		{"2020-03-05", "March"},
		{"2020-04-05", "April"},
		{"2020-05-05", "May"},
		{"2020-06-05", "June"},
		{"2020-07-05", "July"},
		{"2020-08-05", "August"},
		{"2020-09-05", "September"},
		{"2020-10-05", "October"},
		{"2020-11-05", "November"},
		{"2020-12-05", "December"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToMonthString(), test.output)
	}
}

func TestCarbon_ToMonthString2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0", "en", ""},
		{"0000-00-00", "en", ""},
		{"00:00:00", "en", ""},
		{"0000-00-00 00:00:00", "en", ""},

		{"2020-08-05", "en", "August"},
		{"2020-08-05", "zh-CN", "八月"},
		{"2020-08-05", "zh-Tw", "八月"},
		{"2020-08-05", "jp", "はちがつ"},
		{"2020-08-05", "kr", "팔월"},
		{"2020-08-05", "xx", ""},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.SetLocale(test.param).ToMonthString(), test.output)
	}
}

func TestCarbon_ToShortMonthString1(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "Jan"},
		{"2020-02-05", "Feb"},
		{"2020-03-05", "Mar"},
		{"2020-04-05", "Apr"},
		{"2020-05-05", "May"},
		{"2020-06-05", "Jun"},
		{"2020-07-05", "Jul"},
		{"2020-08-05", "Aug"},
		{"2020-09-05", "Sep"},
		{"2020-10-05", "Oct"},
		{"2020-11-05", "Nov"},
		{"2020-12-05", "Dec"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToShortMonthString(), test.output)
	}
}

func TestCarbon_ToShortMonthString2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0", "en", ""},
		{"0000-00-00", "en", ""},
		{"00:00:00", "en", ""},
		{"0000-00-00 00:00:00", "en", ""},

		{"2020-08-05", "en", "Aug"},
		{"2020-08-05", "zh-CN", "8月"},
		{"2020-08-05", "zh-Tw", "8月"},
		{"2020-08-05", "jp", "8がつ"},
		{"2020-08-05", "kr", "8월"},
		{"2020-08-05", "xx", ""},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.SetLocale(test.param).ToShortMonthString(), test.output)
	}
}

func TestCarbon_ToWeekString1(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "Saturday"},
		{"2020-08-02", "Sunday"},
		{"2020-08-03", "Monday"},
		{"2020-08-04", "Tuesday"},
		{"2020-08-05", "Wednesday"},
		{"2020-08-06", "Thursday"},
		{"2020-08-07", "Friday"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToWeekString(), test.output)
	}
}

func TestCarbon_ToWeekString2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0", "en", ""},
		{"0000-00-00", "en", ""},
		{"00:00:00", "en", ""},
		{"0000-00-00 00:00:00", "en", ""},

		{"2020-08-05", "en", "Wednesday"},
		{"2020-08-05", "zh-CN", "星期三"},
		{"2020-08-05", "zh-Tw", "星期三"},
		{"2020-08-05", "jp", "もくようび"},
		{"2020-08-05", "kr", "수요일"},
		{"2020-08-05", "xx", ""},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.SetLocale(test.param).ToWeekString(), test.output)
	}
}

func TestCarbon_ToShortWeekString1(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "Sat"},
		{"2020-08-02", "Sun"},
		{"2020-08-03", "Mon"},
		{"2020-08-04", "Tue"},
		{"2020-08-05", "Wed"},
		{"2020-08-06", "Thu"},
		{"2020-08-07", "Fri"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToShortWeekString(), test.output)
	}
}

func TestCarbon_ToShortWeekString2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0", "en", ""},
		{"0000-00-00", "en", ""},
		{"00:00:00", "en", ""},
		{"0000-00-00 00:00:00", "en", ""},

		{"2020-08-05", "en", "Wed"},
		{"2020-08-05", "zh-CN", "周三"},
		{"2020-08-05", "zh-Tw", "週三"},
		{"2020-08-05", "jp", "週三"},
		{"2020-08-05", "kr", "수요일"},
		{"2020-08-05", "xx", ""},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.SetLocale(test.param).ToShortWeekString(), test.output)
	}
}

func TestCarbon_ToFormatString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "Y年m月d日", ""},
		{"0", "Y年m月d日", ""},
		{"0000-00-00", "Y年m月d日", ""},
		{"00:00:00", "Y年m月d日", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", ""},

		{"2020-08-05 13:14:15", "Y年m月d日", "2020年08月05日"},
		{"2020-08-05 01:14:15", "j", "5"},
		{"2020-08-05 01:14:15", "W", "32"},
		{"2020-08-05 01:14:15", "M", "Aug"},
		{"2020-08-05 01:14:15", "F", "August"},
		{"2020-08-05 01:14:15", "N", "3"},
		{"2020-08-05 01:14:15", "L", "1"},
		{"2020-08-05 01:14:15", "L", "1"},
		{"2021-08-05 01:14:15", "L", "0"},
		{"2020-08-05 01:14:15", "G", "1"},
		{"2020-08-05 13:14:15", "U", "1596604455"},
		{"2020-08-05 13:14:15.999", "u", "999"},
		{"2020-08-05 13:14:15", "w", "2"},
		{"2020-08-05 13:14:15", "t", "31"},
		{"2020-08-05 13:14:15", "z", "217"},
		{"2020-08-05 13:14:15", "e", "Local"},
		{"2020-08-05 13:14:15", "jS", "5th"},
		{"2020-08-22 13:14:15", "jS", "22nd"},
		{"2020-08-23 13:14:15", "jS", "23rd"},
		{"2020-08-31 13:14:15", "jS", "31st"},
		{"2020-08-31 13:14:15", "I\\t \\i\\s Y-m-d H:i:s", "It is 2020-08-31 13:14:15"},
		{"2020-08-05 13:14:15", "l jS of F Y h:i:s A", "Wednesday 5th of August 2020 01:14:15 PM"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToFormatString(test.param), test.output)
	}
}

func TestCarbon_Format(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "Y年m月d日", ""},
		{"0", "Y年m月d日", ""},
		{"0000-00-00", "Y年m月d日", ""},
		{"00:00:00", "Y年m月d日", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", ""},

		{"2020-08-05 13:14:15", "Y年m月d日", "2020年08月05日"},
		{"2020-08-05 01:14:15", "j", "5"},
		{"2020-08-05 01:14:15", "W", "32"},
		{"2020-08-05 01:14:15", "M", "Aug"},
		{"2020-08-05 01:14:15", "F", "August"},
		{"2020-08-05 01:14:15", "N", "3"},
		{"2020-08-05 01:14:15", "L", "1"},
		{"2020-08-05 01:14:15", "L", "1"},
		{"2021-08-05 01:14:15", "L", "0"},
		{"2020-08-05 01:14:15", "G", "1"},
		{"2020-08-05 13:14:15", "U", "1596604455"},
		{"2020-08-05 13:14:15.999", "u", "999"},
		{"2020-08-05 13:14:15", "w", "2"},
		{"2020-08-05 13:14:15", "t", "31"},
		{"2020-08-05 13:14:15", "z", "217"},
		{"2020-08-05 13:14:15", "e", "Local"},
		{"2020-08-05 13:14:15", "jS", "5th"},
		{"2020-08-22 13:14:15", "jS", "22nd"},
		{"2020-08-23 13:14:15", "jS", "23rd"},
		{"2020-08-31 13:14:15", "jS", "31st"},
		{"2020-08-31 13:14:15", "I\\t \\i\\s Y-m-d H:i:s", "It is 2020-08-31 13:14:15"},
		{"2020-08-05 13:14:15", "l jS of F Y h:i:s A", "Wednesday 5th of August 2020 01:14:15 PM"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Format(test.param), test.output)
	}
}

func TestCarbon_ToDayDateTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, Aug 5, 2020 1:14 PM"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDayDateTimeString(), test.output)
	}
}

func TestCarbon_ToDayDateTimeStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, Aug 5, 2020 5:14 AM"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDayDateTimeStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToDateTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output)
	}
}

func TestCarbon_ToDateTimeStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 05:14:15"},
		{"2020-08-05", "2020-08-04 16:00:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05"},
		{"2020-08-05", "2020-08-05"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateString(), test.output)
	}
}

func TestCarbon_ToDateStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05"},
		{"2020-08-05", "2020-08-04"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "13:14:15"},
		{"2020-08-05", "00:00:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimeString(), test.output)
	}
}

func TestCarbon_ToTimeStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05:14:15"},
		{"2020-08-05", "16:00:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToTimeStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToAtomString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToAtomString(), test.output)
	}
}

func TestCarbon_ToAtomStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T05:14:15Z"},
		{"2020-08-05", "2020-08-04T16:00:00Z"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToAtomStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToAnsicString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 2020"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToAnsicString(), test.output)
	}
}

func TestCarbon_ToAnsicStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug  5 05:14:15 2020"},
		{"2020-08-05", "Tue Aug  4 16:00:00 2020"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToAnsicStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToCookieString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-2020 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-2020 00:00:00 CST"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToCookieString(), test.output)
	}
}

func TestCarbon_ToCookieStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-2020 05:14:15 UTC"},
		{"2020-08-05", "Tuesday, 04-Aug-2020 16:00:00 UTC"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToCookieStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRssString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRssString(), test.output)
	}
}

func TestCarbon_ToRssStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 05:14:15 +0000"},
		{"2020-08-05", "Tue, 04 Aug 2020 16:00:00 +0000"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRssStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToW3cString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToW3cString(), test.output)
	}
}

func TestCarbon_ToW3cStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T05:14:15Z"},
		{"2020-08-05", "2020-08-04T16:00:00Z"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToW3cStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToUnixDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 CST 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 CST 2020"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToUnixDateString(), test.output)
	}
}

func TestCarbon_ToUnixDateStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug  5 05:14:15 UTC 2020"},
		{"2020-08-05", "Tue Aug  4 16:00:00 UTC 2020"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToUnixDateStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRubyDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug 05 13:14:15 +0800 2020"},
		{"2020-08-05", "Wed Aug 05 00:00:00 +0800 2020"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRubyDateString(), test.output)
	}
}

func TestCarbon_ToRubyDateStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug 05 05:14:15 +0000 2020"},
		{"2020-08-05", "Tue Aug 04 16:00:00 +0000 2020"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRubyDateStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToKitchenString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "1:14PM"},
		{"2020-08-05", "12:00AM"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToKitchenString(), test.output)
	}
}

func TestCarbon_ToKitchenStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "5:14AM"},
		{"2020-08-05", "4:00PM"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToKitchenStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc822String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05 Aug 20 13:14 CST"},
		{"2020-08-05", "05 Aug 20 00:00 CST"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc822String(), test.output)
	}
}

func TestCarbon_ToRfc822StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05 Aug 20 05:14 UTC"},
		{"2020-08-05", "04 Aug 20 16:00 UTC"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc822StringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc822zString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05 Aug 20 13:14 +0800"},
		{"2020-08-05", "05 Aug 20 00:00 +0800"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc822zString(), test.output)
	}
}

func TestCarbon_ToRfc822zStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05 Aug 20 05:14 +0000"},
		{"2020-08-05", "04 Aug 20 16:00 +0000"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc822zStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc850String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-20 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-20 00:00:00 CST"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc850String(), test.output)
	}
}

func TestCarbon_ToRfc850StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-20 05:14:15 UTC"},
		{"2020-08-05", "Tuesday, 04-Aug-20 16:00:00 UTC"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc850StringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc1036String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 20 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 20 00:00:00 +0800"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc1036String(), test.output)
	}
}

func TestCarbon_ToRfc1036StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 20 05:14:15 +0000"},
		{"2020-08-05", "Tue, 04 Aug 20 16:00:00 +0000"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc1036StringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc1123String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 CST"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 CST"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc1123String(), test.output)
	}
}

func TestCarbon_ToRfc1123StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 05:14:15 UTC"},
		{"2020-08-05", "Tue, 04 Aug 2020 16:00:00 UTC"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc1123StringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc1123ZString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc1123ZString(), test.output)
	}
}

func TestCarbon_ToRfc1123ZStringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 05:14:15 +0000"},
		{"2020-08-05", "Tue, 04 Aug 2020 16:00:00 +0000"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc1123ZStringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc2822String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc2822String(), test.output)
	}
}

func TestCarbon_ToRfc2822StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 05:14:15 +0000"},
		{"2020-08-05", "Tue, 04 Aug 2020 16:00:00 +0000"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc2822StringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc3339String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc3339String(), test.output)
	}
}

func TestCarbon_ToRfc3339StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T05:14:15Z"},
		{"2020-08-05", "2020-08-04T16:00:00Z"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc3339StringWithTimezone(UTC), test.output)
	}
}

func TestCarbon_ToRfc7231String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 GMT"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 GMT"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc7231String(), test.output)
	}
}

func TestCarbon_ToRfc7231StringWithTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 05:14:15 GMT"},
		{"2020-08-05", "Tue, 04 Aug 2020 16:00:00 GMT"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToRfc7231StringWithTimezone(UTC), test.output)
	}
}
