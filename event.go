package cagliostro

import (
	"time"

	"github.com/KuroiKitsu/go-gbf"
)

const (
	EventCacheTTL = 30 * time.Minute
)

// CachedEvent wraps an event, its details, and its expiration date.
type CachedEvent struct {
	Event     *gbf.Event
	Details   *gbf.EventDetails
	ExpiresAt time.Time
}

// CurrentEvents returns the ongoing events.
func (c *Cagliostro) CurrentEvents() (events []*CachedEvent, err error) {
	if len(c.currentEvents) > 0 {
		alive := true

		for _, event := range c.currentEvents {
			if time.Now().After(event.ExpiresAt) {
				alive = false
				break
			}
		}

		if alive {
			events = c.currentEvents

			return
		}
	}

	var (
		currentEvents []*gbf.Event
	)

	currentEvents, err = gbf.CurrentEvents()
	if err != nil {
		return
	}

	events = make([]*CachedEvent, len(currentEvents))
	for i, event := range currentEvents {
		cachedEvent := &CachedEvent{
			Event:     event,
			ExpiresAt: time.Now().Add(EventCacheTTL),
		}

		cachedEvent.Details, err = gbf.EventDetailsURL(event.URL)
		if err != nil {
			return
		}

		events[i] = cachedEvent
	}

	c.currentEvents = events

	return
}

// UpcomingEvents returns the upcoming events.
func (c *Cagliostro) UpcomingEvents() (events []*CachedEvent, err error) {
	if len(c.upcomingEvents) > 0 {
		alive := true

		for _, event := range c.upcomingEvents {
			if time.Now().After(event.ExpiresAt) {
				alive = false
				break
			}
		}

		if alive {
			events = c.upcomingEvents

			return
		}
	}

	var (
		upcomingEvents []*gbf.Event
	)

	upcomingEvents, err = gbf.UpcomingEvents()
	if err != nil {
		return
	}

	events = make([]*CachedEvent, len(upcomingEvents))
	for i, event := range upcomingEvents {
		cachedEvent := &CachedEvent{
			Event:     event,
			ExpiresAt: time.Now().Add(EventCacheTTL),
		}

		if event.URL != "" {
			cachedEvent.Details, err = gbf.EventDetailsURL(event.URL)
			if err != nil {
				return
			}
		}

		events[i] = cachedEvent
	}

	c.upcomingEvents = events

	return
}
