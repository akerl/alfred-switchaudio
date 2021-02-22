package utils

import (
	"os/exec"
	"strings"
)

// Run executes a SwitchAudio CLI command
func Run(args ...string) ([]byte, error) {
	return exec.Command("SwitchAudioSource", args...).Output()
}

// CurrentDevice returns the current device for a given direction
func CurrentDevice(direction string) (string, error) {
	out, err := Run("-c", "-t", direction)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// AllDevices returns all devices for a given direction
func AllDevices(direction string) ([]string, error) {
	out, err := Run("-a", "-t", direction)
	if err != nil {
		return []string{}, err
	}
	outstr := string(out)
	outstr = strings.ReplaceAll(outstr, " ("+direction+")", "")
	entries := strings.Split(outstr, "\n")
	return entries, nil
}

// AllOtherDevices returns all devices for a given direction except the current device
func AllOtherDevices(direction string) ([]string, error) {
	all, err := AllDevices(direction)
	if err != nil {
		return []string{}, err
	}
	current, err := CurrentDevice(direction)
	if err != nil {
		return []string{}, err
	}
	others := []string{}
	for _, device := range all {
		if device != current {
			others = append(others, device)
		}
	}
	return others, nil
}

// SetDevice sets the audio device for a direction
func SetDevice(direction, target string) error {
	_, err := Run("-t", direction, "-s", target)
	return err
}
