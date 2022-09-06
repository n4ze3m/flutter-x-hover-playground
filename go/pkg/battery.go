package pkg

import (
	"fmt"

	"github.com/distatus/battery"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "com.myapp/battery"

type MyBatteryPlugin struct{}

var _ flutter.Plugin = &MyBatteryPlugin{}

func (p *MyBatteryPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getBattery", handleGetBatteryLevel)
	return nil // no error
}

func handleGetBatteryLevel(arguments interface{}) (reply interface{}, err error) {
	var result string
	batteries, err := battery.GetAll()
	if err != nil {
		return nil, err
	}

	for _, b := range batteries {
		result += fmt.Sprintf("Current State: %s\n", b.State.String())
		result += fmt.Sprintf("Current Capacity: %f mWh\n", b.Current)
		result += fmt.Sprintf("Charge rate: %f mW,\n", b.ChargeRate)
		result += fmt.Sprintf("Voltage: %f mW,\n", b.Voltage)
	}

	return result, nil
}
