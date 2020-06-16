package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c       = cache.New(5*time.Minute, 10*time.Minute)
	testApp = &App{
		Key:   "439d4b804bc8187953eb36d2a8c26a02",
		Cache: *c,
		Url:   "https://samples.openweathermap.org/data/2.5/weather",
	}
	city = "London"
	r    = TestResposnse{}
	m    = InfoMessage{}
)

type TestResposnse struct {
	Result []Weather
}

func TestApiHandlerWithNoCity(t *testing.T) {

	req := httptest.NewRequest("GET", "http://127.0.0.1:5000/api", nil)
	w := httptest.NewRecorder()

	testApp.ApiHandler(w, req)

	resp := w.Result()
	json.NewDecoder(resp.Body).Decode(&m)

	if resp.StatusCode != 422 {
		t.Errorf("Expected status code is 422, but it was %d instead.", resp.StatusCode)
	}
	if m.Message != "Url Param 'city' is missing!" {
		t.Errorf("Expected value is \"Url Param 'city' is missing!\", but it was %s instead.", m.Message)
	}
}
func TestApiHandlerWithOneCity(t *testing.T) {

	req := httptest.NewRequest("GET", "http://127.0.0.1:5000/api", nil)
	w := httptest.NewRecorder()

	q := req.URL.Query()
	q.Add("city", city)
	req.URL.RawQuery = q.Encode()

	testApp.ApiHandler(w, req)

	resp := w.Result()
	json.NewDecoder(resp.Body).Decode(&r)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code is 200, but it was %d instead.", resp.StatusCode)
	}
	if r.Result[0].Name != "London" {
		t.Errorf("Expected value is London, but it was %s instead.", r.Result[0].Name)
	}
}
func TestApiHandlerWithTwoCity(t *testing.T) {

	req := httptest.NewRequest("GET", "http://127.0.0.1:5000/api", nil)
	w := httptest.NewRecorder()

	q := req.URL.Query()
	s := fmt.Sprintf("%s,%s", city, city)
	q.Add("city", s)
	req.URL.RawQuery = q.Encode()

	testApp.ApiHandler(w, req)

	resp := w.Result()
	json.NewDecoder(resp.Body).Decode(&r)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code is 200, but it was %d instead.", resp.StatusCode)
	}
	if r.Result[0].Name != "London" {
		t.Errorf("Expected value is London, but it was %s instead.", r.Result[0].Name)
	}
	if r.Result[1].Name != "London" {
		t.Errorf("Expected value is London, but it was %s instead.", r.Result[0].Name)
	}
}

func TestCacheApiHandler(t *testing.T) {

	testApp.Cache.Flush()

	req := httptest.NewRequest("GET", "http://127.0.0.1:5000/api", nil)
	w := httptest.NewRecorder()

	q := req.URL.Query()
	q.Add("city", city)
	req.URL.RawQuery = q.Encode()

	testApp.ApiHandler(w, req)

	_, found := testApp.Cache.Get(city)
	if !found {
		t.Errorf("No entry in cache, but it was expected.")
	}
}

func TestGetOWResponse(t *testing.T) {

	Weather, err := GetOWResponse(testApp.Key, city, testApp.Url)
	if err != nil {
		t.Errorf("No errors expected, but get %s", err)
	}

	if Weather.Name != city {
		t.Errorf("Expected entry was %s, but it was %s instead.", city, Weather.Name)
	}

	testType(Weather.Name, reflect.String, t)
	testType(Weather.Main.Temp, reflect.Float32, t)
	testType(Weather.Main.Pressure, reflect.Int, t)
	testType(Weather.Main.Humidity, reflect.Int, t)
	testType(Weather.Main.TempMin, reflect.Float32, t)
	testType(Weather.Main.TempMax, reflect.Float32, t)
	testType(Weather.Wind.Speed, reflect.Float32, t)
	testType(Weather.Wind.Deg, reflect.Int, t)
	testType(Weather.Clouds.All, reflect.Int, t)

}

func testType(w, v interface{}, t *testing.T) {

	if reflect.TypeOf(w).Kind() != v {
		t.Errorf("Expected entry was %v, but it was %v instead.", v, w)
	}
}
