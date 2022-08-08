package providers

import "github.com/gookit/event"

func BootstrapEvent() {
	event.On("update-exercise", event.ListenerFunc(func(e event.Event) error {
		return nil
	}), event.Normal)
}
