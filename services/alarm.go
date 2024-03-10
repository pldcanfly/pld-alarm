package services

import (
	"fmt"
	"time"
)

type Alarm struct {
	Active      bool
	Hour        int
	Minute      int
	NextRing    time.Time
	Snoozing    bool
	AlarmSound  string
	MediaState  *MediaState
	Ringing     bool
	SnoozeCount int
}

func NewAlarm(h int, m int, s string, ms *MediaState) *Alarm {
	a := &Alarm{
		Active:      false,
		Hour:        h,
		Minute:      m,
		Snoozing:    false,
		AlarmSound:  s,
		MediaState:  ms,
		Ringing:     false,
		SnoozeCount: 0,
	}

	go func() {
		for {

			if a.Active {
				now := time.Now()

				if now.Compare(a.NextRing) >= 1 {
					a.Trigger()
				}

			}

			time.Sleep(1 * time.Second)
		}
	}()

	return a
}

func (a *Alarm) SetActive(active bool) {
	now := time.Now()
	d := now.Day()
	if a.Hour <= now.Hour() {
		d++
	}
	a.Active = active
	a.NextRing = time.Date(now.Year(), now.Month(), d, a.Hour, a.Minute, 0, 0, now.Location())
}

func (a *Alarm) Snooze() {
	if !a.Active {
		return
	}

	a.SnoozeCount++
	if a.SnoozeCount > 5 {
		tmrw := time.Now().Add(24 * time.Hour)
		a.NextRing = time.Date(tmrw.Year(), tmrw.Month(), tmrw.Day(), a.Hour, a.Minute, 0, 0, tmrw.Location())
		a.Snoozing = false
	} else {

		a.Snoozing = true
		a.NextRing = time.Now().Add(10 * time.Minute)
	}

	a.Ringing = false
	// a.NextRing = time.Now().Add(5 * time.Minute)
	a.MediaState.StopAudio()

	fmt.Println("snooze", a.NextRing)
}

func (a *Alarm) Trigger() error {
	a.Snoozing = false
	if !a.MediaState.Playing {
		a.Ringing = true
		a.MediaState.PlayAudio(a.AlarmSound)
	}
	return nil
}
