package main

import (
	"github.com/patrickmn/go-cache"
)

type Weather struct {
	Name   string   `json:"name"`
	Main   MainInfo `json:"main"`
	Wind   Wind     `json:"wind"`
	Clouds Clouds   `json:"clouds"`
}

type MainInfo struct {
	Temp     float32 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float32 `json:"temp_min"`
	TempMax  float32 `json:"temp_max"`
}

type Wind struct {
	Speed float32 `json:"speed"`
	Deg   int     `json:"Deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type App struct {
	Key   string
	Cache cache.Cache
	Url   string
}
