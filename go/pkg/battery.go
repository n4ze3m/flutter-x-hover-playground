package pkg

import (
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
	batteryLevel := int32(55) // Your platform-specific API
	return batteryLevel, nil
}
