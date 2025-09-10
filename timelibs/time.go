package timelibs

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

type timeImpl struct {
	time time.Time
}

func (ti *timeImpl) Format(layout string) string {
	return ti.time.Format(layout)
}

func (ti *timeImpl) Time() time.Time {
	return ti.time
}

func NewTime(t time.Time) Time {
	return &timeImpl{time: t.In(location)}
}

var Now = func() Time {
	return &timeImpl{
		time: time.Now().In(location),
	}
}
