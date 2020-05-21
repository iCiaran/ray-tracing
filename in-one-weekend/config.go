package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Aspect          AspectConfig `json:"aspect"`
	ImageWidth      int          `json:"imageWidth"`
	SamplesPerPixel int          `json:"samplesPerPixel"`
	MaxDepth        int          `json:"maxDepth"`
	Camera          CameraConfig `json:"camera"`
}

type AspectConfig struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type CameraConfig struct {
	From        Vec3Config `json:"from"`
	To          Vec3Config `json:"to"`
	Up          Vec3Config `json:"up"`
	DistToFocus float64    `json:"distToFocus"`
	Aperture    float64    `json:"aperture"`
	VFov        float64    `json:"vFov"`
}

type Vec3Config struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func loadConfig(filePath string) (*Config, error) {
	configFile, err := os.Open("config.json")

	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	var config Config

	err = json.Unmarshal([]byte(byteValue), &config)

	if err != nil {
		return nil, err
	}

	return &config, nil

}
