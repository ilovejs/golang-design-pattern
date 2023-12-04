package alternative

import "testing"

func TestAlternativeCmdPattern(t *testing.T) {
	tv := &Tv{}

	// command includes device
	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
