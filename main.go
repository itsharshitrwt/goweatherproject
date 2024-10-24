// package main

// import(
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )


// type apiConfigData struct{
// 	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
// }


// type weatherData struct{
// 	name string `jaon:"name"`
// 	Main struct{
// 		kelvin float64 `jaon:"temp"`

// 	}`json:"main"`
// }

// func loadApiConfig(filename string) (apiConfigData, error){
// 	bytes, err := ioutil.ReadFile((filename))

// 	if err != nil{
// 		return apiConfigData{}, err
// 	}

// 	var c apiConfigData

// 	err = json.Unmarshal(bytes, &c)
// 	if err != nil {
// 		return apiConfigData{}, err
// 	}
// 	return c, nil
// }

// func hello(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("hello from go!"))
// }
// func query(city string)(weatherData, error){
// 	apiconfig, err = loadApiConfig(".apiconfig")
// 	if err!= nil{
// 		return weatherData{}, err
// 	}

// 	resp, err := http.Get("http://api.openweather.org/data/2.5/weather?APPID="+ apiConfig.OpenWeatherMapApiKey + "&q"+city)
// 	if err!= nil{
// 		return weatherData{}, err
// 	}

// 	defer resp.Body.close()

// 	var d weatherData
// 	if err := json.NewDecoder(resp.Body).Decode(&d); err!= nil{
//           return weatherData{}, err
// 	}
// 	return d, nil
// }
// func main(){
// 	http.HandleFunc("/hello", hello)

// 	http.HandleFunc("/weather/",
// 	func(w http.Responsewriter, r *http.Request){
// 		city := strings.SplitN(e.URL.Path, "/",3)[2]
// 		data, err := query(city)
// 		if err != nil{
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("content-Type", "application/json; charset = utf-8")
//         json.NewEncoder(w).Encode(data)
// 	}
// )

//     http.ListenAndServe(":8080", nil)
// }

// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// // apiConfigData holds the API configuration data
// type apiConfigData struct {
// 	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
// }

// // weatherData holds the weather data
// type weatherData struct {
// 	Name string `json:"name"`
// 	Main struct {
// 		Temp float64 `json:"temp"`
// 	} `json:"main"`
// }

// // loadApiConfig loads the API configuration from a file
// func loadApiConfig(filename string) (apiConfigData, error) {
// 	bytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return apiConfigData{}, err
// 	}

// 	var c apiConfigData
// 	err = json.Unmarshal(bytes, &c)
// 	if err != nil {
// 		return apiConfigData{}, err
// 	}
// 	return c, nil
// }

// // hello handles the /hello endpoint
// func hello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello from go!"))
// }

// // query queries the OpenWeatherMap API for the weather data
// func query(city string) (weatherData, error) {
// 	apiConfig, err := loadApiConfig(".apiconfig")
// 	if err != nil {
// 		return weatherData{}, err
// 	}

// 	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
// 	if err != nil {
// 		return weatherData{}, err
// 	}

// 	defer resp.Body.Close()

// 	var d weatherData
// 	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
// 		return weatherData{}, err
// 	}
// 	return d, nil
// }

// func main() {
// 	http.HandleFunc("/hello", hello)

// 	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
// 		city := strings.SplitN(r.URL.Path, "/", 3)[2]
// 		data, err := query(city)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		json.NewEncoder(w).Encode(data)
// 	})

// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// apiConfigData holds the API configuration data
type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// weatherData holds the weather data
type weatherData struct {
	Name    string  `json:"name"`
	Main    struct {
		TempKelvin float64 `json:"temp"` // Temperature in Kelvin
		TempCelsius float64 `json:"temp_celsius"` // Temperature in Celsius
	} `json:"main"`
}

// loadApiConfig loads the API configuration from a file
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

// hello handles the /hello endpoint
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!"))
}

// query queries the OpenWeatherMap API for the weather data
func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiconfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	// Convert temperature from Kelvin to Celsius
	d.Main.TempCelsius = d.Main.TempKelvin - 273.15

	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}