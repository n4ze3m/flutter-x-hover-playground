package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/n4ze3m/hover-fluter/go/pkg"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	flutter.AddPlugin(&pkg.MyBatteryPlugin{}),
}
