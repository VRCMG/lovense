# Lovense

[![GoDoc](https://godoc.org/github.com/sextech/lovense?status.svg)](https://godoc.org/github.com/sextech/lovense)

Non-official Go Lovense API **UNDER DEVELOPMENT**

### Getting started

You need to download and setup Lovense Connect application on you computer or mobile.  
Follow this link for more information: https://www.lovense.com/cam-model/download

#### Start remote control

You can discover Lovense sex toys available on your network (LAN) by using `Discover()` method.

```go
remote := lovense.NewRemote()
toys, err := remote.Discover()

if err != nil {
    log.Fatal(err)
}
```

#### Vibrate

Lush, Hush, Ambi, Edge, Domi and Osci

```go
// Start vibration (speed 0 to 20)
toy.Vibrate(lovense.AllVibrator, 10)

// Stop vibration
toy.Vibrate(lovense.AllVibrator, 0)
```

On Edge you can select Vibrator

```go
toy.Vibrate(lovense.Vibrator1, 10)
toy.Vibrate(lovense.Vibrator1, 0)
```

#### Rotate

Nora

```go
// Start rotation (speed 0 to 20)
toy.Rotate(lovense.Normal, 10)
toy.Rotate(lovense.Clockwise, 10)
toy.Rotate(lovense.AntiClockwise, 10)

// Change rotation direction
toy.RotateChange()
```

#### Air

Max

```go
// Start contraction (speed 0 to 3)
toy.AirAuto(2)

// Pump in the air
toy.AirIn()

// Release air
toy.AirOut()

```

#### Preset

Lush, Hush, Ambi, Edge, Domi and Osci

```go
// Start preset (patter 0 to 3)
toy.Preset(2)
```