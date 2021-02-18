package carbon

// DaysInYear 获取本年的总天数
func (c Carbon) DaysInYear() int {
	if c.IsZero() {
		return 0
	}
	days := DaysPerNormalYear
	if c.IsLeapYear() {
		days = DaysPerLeapYear
	}
	return days
}

// DaysInMonth 获取本月的总天数
func (c Carbon) DaysInMonth() int {
	if c.IsZero() {
		return 0
	}
	return c.EndOfMonth().Time.In(c.Loc).Day()
}

// MonthOfYear 获取本年的第几月
func (c Carbon) MonthOfYear() int {
	if c.IsZero() {
		return 0
	}
	return int(c.Time.In(c.Loc).Month())
}

// DayOfYear 获取本年的第几天
func (c Carbon) DayOfYear() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).YearDay()
}

// DayOfMonth 获取本月的第几天
func (c Carbon) DayOfMonth() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Day()
}

// DayOfWeek 获取本周的第几天
func (c Carbon) DayOfWeek() int {
	if c.IsZero() {
		return 0
	}
	return int(c.Time.In(c.Loc).Weekday())
}

// WeekOfYear 获取本年的第几周
func (c Carbon) WeekOfYear() int {
	if c.IsZero() {
		return 0
	}
	_, week := c.Time.In(c.Loc).ISOWeek()
	return week
}

// WeekOfMonth 获取本月的第几周
func (c Carbon) WeekOfMonth() int {
	if c.IsZero() {
		return 0
	}
	day := c.Time.In(c.Loc).Day()
	if day < DaysPerWeek {
		return 1
	}
	return day%DaysPerWeek + 1
}

// Timezone 获取时区
func (c Carbon) Timezone() string {
	return c.Loc.String()
}

// Locale 获取语言区域
func (c Carbon) Locale() string {
	return c.Lang.locale
}

// Age 获取年龄
func (c Carbon) Age() int {
	if c.IsZero() {
		return 0
	}
	now := Now()
	if c.ToTimestamp() > now.ToTimestamp() {
		return 0
	}
	return int(c.DiffInYears(now))
}

// 获取当前年
func (c Carbon) Year() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Year()
}

// 获取当前季度
func (c Carbon) Quarter() int {
	if c.IsZero() {
		return 0
	}
	month := c.Month()
	switch {
	case month >= 10:
		return 4
	case month >= 7:
		return 3
	case month >= 4:
		return 2
	case month >= 1:
		return 1
	}
	return 0
}

// 获取当前月
func (c Carbon) Month() int {
	if c.IsZero() {
		return 0
	}
	return c.MonthOfYear()
}

// 获取当前日
func (c Carbon) Day() int {
	if c.IsZero() {
		return 0
	}
	return c.DayOfMonth()
}

// 获取当前小时
func (c Carbon) Hour() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Hour()
}

// 获取当前分钟数
func (c Carbon) Minute() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Minute()
}

// 获取当前秒数
func (c Carbon) Second() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Second()
}

// 获取当前毫秒数
func (c Carbon) Millisecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e6
}

// 获取当前微秒数
func (c Carbon) Microsecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e9
}

// 获取当前纳秒数
func (c Carbon) Nanosecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond()
}