package lovense

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
)

type Toy struct {
	ID     string
	Host   string
	Name   string
	Status Status
}

type Status int

const (
	Disconnected Status = 0
	Connected    Status = 1
)

var client *http.Client

func init() {
	// Remote API use self-signed certificate, so we skip verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{
		Transport: transport,
	}
}

type Vibrator int
type RotateDirection int

const (
	AllVibrator Vibrator = 0
	Vibrator1   Vibrator = 1
	Vibrator2   Vibrator = 2

	Normal        RotateDirection = 1
	Clockwise     RotateDirection = 2
	AntiClockwise RotateDirection = 3
)

func (t *Toy) sendCommand(method string) error {
	resp, err := client.Get(t.Host + method)

	if err != nil {
		return fmt.Errorf("failed to send command: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("command return bad status code: %d %s", resp.StatusCode, resp.Status)
	}

	return nil
}

// Vibrate change speed vibration of Lush, Hush, Ambi, Edge, Domi and Osci sex toys
// Only Edge can specify vibrator 1 or 2
func (t *Toy) Vibrate(vibrator Vibrator, speed int) error {
	if speed < 0 || speed > 20 {
		return errors.New("speed value not allowed (range is 0 to 20)")
	}

	method := "Vibrate"

	if vibrator != AllVibrator {
		method += string(vibrator)
	}

	return t.sendCommand(fmt.Sprintf("/%s?v=%d&t=%s", method, speed, t.ID))
}

// RotateDirection change speed rotation of Nora sex toy
func (t *Toy) Rotate(rotate RotateDirection, speed int) error {
	if speed < 0 || speed > 20 {
		return errors.New("speed value not allowed (range is 0 to 20)")
	}

	method := "Rotate"

	switch rotate {
	case Clockwise:
		method += "Clockwise"
	case AntiClockwise:
		method += "AntiClockwise"
	}

	return t.sendCommand(fmt.Sprintf("/%s?v=%d&t=%s", method, speed, t.ID))
}

// RotateChange change the rotation direction of Nora sex toy
func (t *Toy) RotateChange() error {
	return t.sendCommand(fmt.Sprintf("/RotateChange?t=%s", t.ID))
}

// AirAuto start contraction of Max sex toy
func (t *Toy) AirAuto(speed int) error {
	if speed < 0 || speed > 3 {
		return errors.New("speed value not allowed (range is 0 to 3)")
	}

	return t.sendCommand(fmt.Sprintf("/AirAuto?v=%d&t=%s", speed, t.ID))
}

// AirIn pump in the air of Max sex toy
func (t *Toy) AirIn() error {
	return t.sendCommand(fmt.Sprintf("/AirIn?t=%s", t.ID))
}

// AirOut release the air of Max sex toy
func (t *Toy) AirOut() error {
	return t.sendCommand(fmt.Sprintf("/AirOut?t=%s", t.ID))
}

// Preset vibrate the toy by predefined patterns of Lush, Hush, Ambi, Edge, Domi and Osci
func (t *Toy) Preset(pattern int) error {
	if pattern < 0 || pattern > 3 {
		return errors.New("pattern value not allowed (range is 0 to 3)")
	}

	return t.sendCommand(fmt.Sprintf("/Preset?v=%d&t=%s", pattern, t.ID))
}
