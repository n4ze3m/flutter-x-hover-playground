package pkg

import (
	"github.com/distatus/battery"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "com.myapp/battery"

type MyBatteryPlugin struct{}

var _ flutter.Plugin = &MyBatteryPlugin{}

func (p *MyBatteryPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getBatteryLevel", handleGetBatteryLevel)
	return nil // no error
}

func handleGetBatteryLevel(arguments interface{}) (reply interface{}, err error) {
	var data = map[string]interface{}{}
	batteries, err := battery.GetAll()
	if err != nil {
		return nil, err
	}
	for _, b := range batteries {
		data["percentage"] = b.Current
		data["state"] = b.State.String()
		data["volts"] = b.Voltage
		data["isCharging"] = b.State == battery.Charging
	}
	return data, nil
}
