package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_FromStdTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt := time.Now().In(loc)
	expected := tt.Format(DateTimeLayout)
	actual := FromStdTime(tt).ToDateTimeString()
	assert.Equal(t, expected, actual)
}

func TestCarbon_Time2Carbon(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt := time.Now().In(loc)
	expected := tt.Format(DateTimeLayout)
	actual := Time2Carbon(tt).ToDateTimeString()
	assert.Equal(t, expected, actual)
}

func TestCarbon_Carbon2Time(t *testing.T) {
	expected := time.Now().Format(DateTimeLayout)
	actual := Now().Carbon2Time().Format(DateTimeLayout)
	assert.Equal(t, expected, actual)
}
