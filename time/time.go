package time

import "time"

var location = time.UTC

// SetLocation sets the location in the process. After calling this, generated Time is in the timezone of loc.
func SetLocation(loc *time.Location) {
	location = loc
}

type Time interface {
	Format(layout string) string
	Time() time.Time
}

type iTime struct {
	time time.Time
}

func (it *iTime) Format(layout string) string {
	return it.time.Format(layout)
}

func (it *iTime) Time() time.Time {
	return it.time
}

func NewTime(t time.Time) Time {
	return &iTime{time: t.In(location)}
}

var Now = func() Time {
	return &iTime{
		time: time.Now().In(location),
	}
}
