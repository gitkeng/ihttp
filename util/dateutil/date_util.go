package dateutil

import (
	"fmt"
	"time"
)

func DefaultCurrentDateString() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func DefaultCurrentDateTimeString() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetCurrentEpochTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func countDigit(number int64) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func DateTime2DefaultString(time *time.Time) string {
	return fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}

func DateTime2Epoch(t *time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func Time2Epoch(t time.Time, digit int) (int64, error) {
	switch digit {
	case 13:
		return t.UnixNano() / int64(time.Millisecond), nil
	case 10:
		return t.Unix(), nil
	default:
		return -1, fmt.Errorf("support only 10, 13 digit epoch time")
	}
}

func Epoch2Time(epoch int64) (time.Time, error) {
	digit := countDigit(epoch)
	switch digit {
	case 13:
		return time.UnixMilli(epoch), nil
	case 10:
		return time.Unix(epoch, 0), nil
	default:
		return time.Time{}, fmt.Errorf("support only 10, 13 digit epoch time")
	}
}

func DefaultEpoch2DateTimeString(epoch int64) (string, error) {
	epochTime := time.Unix(epoch/int64(time.Microsecond), 0)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		epochTime.Year(),
		epochTime.Month(),
		epochTime.Day(),
		epochTime.Hour(),
		epochTime.Minute(),
		epochTime.Second()), nil
}

func DefaultEpoch2DateTimeStringInLocation(epoch int64, loc *time.Location) (string, error) {
	epochTime := time.Unix(epoch/int64(time.Microsecond), 0)
	epochTime = epochTime.In(loc)

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		epochTime.Year(),
		epochTime.Month(),
		epochTime.Day(),
		epochTime.Hour(),
		epochTime.Minute(),
		epochTime.Second()), nil
}

func DefaultDateTimeString2Epoch(dateTime string) (int64, error) {
	t, err := time.Parse("2006-01-02 15:04:05", dateTime)
	if err != nil {
		return -1, err
	}
	return t.UnixNano() / int64(time.Millisecond), nil
}

func LoadThaiLocation() (*time.Location, error) {
	dtLoc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, err
	}
	return dtLoc, nil
}

func DateTimeString2Epoch(layout string, dateTime string) (int64, error) {
	t, err := time.Parse(layout, dateTime)
	if err != nil {
		return -1, err
	}
	return t.UnixNano() / int64(time.Millisecond), nil
}

func DateTimeString2EpochInLocation(layout string, dateTime string, loc *time.Location) (int64, error) {
	t, err := time.ParseInLocation(layout, dateTime, loc)
	if err != nil {
		return -1, err
	}
	return t.UnixNano() / int64(time.Millisecond), nil
}
