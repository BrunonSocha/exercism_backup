package clock

import (
	"fmt"
)

type Clock struct {
	Hour int
	Minute int
}

func New(h, m int) Clock {
	if h < 0 {
		h = 24 + ((24 + h) % 24)
	}
	if m >= 60 {
		h += m / 60
		m = m % 60
	}
	if m < 0 {
		if m >= -60 {
			h--
		} else {
			h -= m / -60
			h--
		}
		m = 60 + (m % -60)
	}
	h = h % 24
	m = m % 60
	if h < 0 {
		h = 24+h
	}
	return Clock{Hour: h, Minute: m}
}

func (c Clock) Add(m int) Clock {
	var ptrC *Clock = &c
	if m + ptrC.Minute >= 60 {
		ptrC.Hour += ((m + ptrC.Minute) / 60)
		ptrC.Minute = (m + ptrC.Minute) % 60
	} else {
		ptrC.Minute += m
	}
	ptrC.Hour = ptrC.Hour % 24
	return c
}

func (c Clock) Subtract(m int) Clock {
	var ptrC *Clock = &c
	if m > ptrC.Minute {
		ptrC.Hour -= m / 60
		ptrC.Minute -= m % 60
		if ptrC.Minute < 0 {
			ptrC.Hour--
			ptrC.Minute = 60 + ptrC.Minute
		}
		if ptrC.Hour < 0 {
			ptrC.Hour = 24 + ptrC.Hour
		}
	} else {
		ptrC.Minute -= m
	}
	ptrC.Hour = ptrC.Hour % 24
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute) 
}
