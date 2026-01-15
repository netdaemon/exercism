package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(h, m int) Clock {
	totalMinutesPerHour := 60
	totalMinutesPerDay := 1440

	totalMinutes := h*totalMinutesPerHour + m

	totalMinutes %= totalMinutesPerDay

	if totalMinutes < 0 {
		totalMinutes += totalMinutesPerDay
	}

	return Clock{
		Hour:   totalMinutes / totalMinutesPerHour,
		Minute: totalMinutes % totalMinutesPerHour,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hour, c.Minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
