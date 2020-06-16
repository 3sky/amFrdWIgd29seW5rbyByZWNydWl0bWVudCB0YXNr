package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/patrickmn/go-cache"
)

var (
	URL string = "https://api.openweathermap.org/data/2.5/weather"
)

func main() {

	var appid string

	appid, err := getAppid()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	c := cache.New(5*time.Minute, 10*time.Minute)

	app := &App{
		Key:   appid,
		Cache: *c,
		Url:   URL,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	log.Println("APP started succesfully!")
	r.Get("/api", app.ApiHandler)
	http.ListenAndServe(":"+getPort(), r)
}

func (a *App) ApiHandler(w http.ResponseWriter, r *http.Request) {

	var WeatherList []interface{}

	cities, ok := r.URL.Query()["city"]

	if !ok || len(cities[0]) < 1 {
		w.WriteHeader(422)
		render.JSON(w, r, &InfoMessage{Message: "Url Param 'city' is missing!"})
		return
	}

	citiesList := strings.Split(cities[0], ",")

	for _, city := range citiesList {
		value, found := a.Cache.Get(city)
		if found {
			WeatherList = append(WeatherList, value)
		} else {
			resp, err := GetOWResponse(a.Key, city, a.Url)
			if err != nil {
				log.Printf("%+v\n", err)
			}
			a.Cache.Set(city, resp, cache.DefaultExpiration)
			WeatherList = append(WeatherList, resp)
		}
	}
	render.JSON(w, r, &Resposnse{Result: WeatherList})
}

func getPort() string {
	var p string
	if len(os.Getenv("APP_PORT")) <= 0 {
		p = "3000"
	} else {
		p = os.Getenv("APP_PORT")
	}

	return p
}

func getAppid() (string, error) {
	var a string
	if len(os.Getenv("APPID")) <= 0 {
		return "", errors.New("No APPID key provided")
	} else {
		a = os.Getenv("APPID")
	}

	return a, nil
}

func GetOWResponse(appid, city, url string) (Weather, error) {

	var client http.Client
	data := Weather{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}
	q := req.URL.Query()
	q.Add("q", city)
	q.Add("appid", appid)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	if resp.StatusCode != 200 {
		return data, errors.New("Cannot connect to OpenWather")
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
