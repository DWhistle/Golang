package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)
type toSort struct{
	values[] float64 	`json:"array"`
	flag bool		`json:"uniq"`
}


func removeDuplicates(initial []float64) []float64 {

	duplicates := map[float64]bool{}
	result := []float64{}

	for v := range initial {
		if duplicates[initial[v]] == true {
		} else
		{
			duplicates[initial[v]] = true
			result = append(result, initial[v])
		}
	}
	return result
}


func doArrayOps(data map[string]interface{}, writer http.ResponseWriter) (toSort, int){

	var toSort toSort

	arr := data["array"].([]interface{})
	for i, element := range arr {
		toSort.values = append(toSort.values, element.(float64))
		if i >= 100 {
			http.Error(writer, "400 bad request.", http.StatusBadRequest)
			return toSort, 1
		}
	}
	if len(toSort.values) == 0 {
		http.Error(writer, "400 bad request.", http.StatusBadRequest)
		return toSort, 1
	}
	if _, ok := data["uniq"]; ok {
		toSort.flag = data["uniq"].(bool)
	}
	log.Println(len(toSort.values))
	if toSort.flag == true {
		toSort.values = removeDuplicates(toSort.values)
	}
	sort.Float64s(toSort.values)
	return toSort, 0
}

func getTemp(city string)(float64, error){

	var data map[string]interface{}

	weather, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=21cfc4b6bc395060673995f604a198e6")
	if err != nil || weather.StatusCode == 404{
		return -2, err
	}
	log.Println(weather)
	log.Println(data)
	values, err := ioutil.ReadAll(weather.Body)
	defer weather.Body.Close()
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(values, &data)
	if err != nil {
		return -1, err
	}
	main := data["main"].(map[string]interface{})
	log.Println(data)
	log.Println(data["main"].(map[string]interface{}))
	temp := main["temp"].(float64)
	log.Println(temp)
	return temp, err
}

func sortArray(writer http.ResponseWriter, request *http.Request) {

	var data map[string]interface{}

	if request.Method == "POST"{
		values, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		log.Println(err)
		err = json.Unmarshal(values, &data)
		log.Println(err)
		toSort , empty := doArrayOps(data, writer)
		if empty == 1{
			return
		}
		retvals, err := json.Marshal(toSort.values)
		log.Println(err)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(retvals)
	}
}
func getTime(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, time.Now().UTC().String())

}



func getWeather(writer http.ResponseWriter, request *http.Request){

	city, _ := request.URL.Query()["city"]

	if len(city) == 0{
		http.Error(writer, "400 bad request.", http.StatusBadRequest)
		return
	}
	temp, _ := getTemp(city[0])
	if temp == -1 {
		http.Error(writer, "400 bad request.", http.StatusBadRequest)
	} else if temp == -2{
		http.Error(writer, "404 not found.", http.StatusNotFound)
	} else{
		fmt.Fprintln(writer, temp)
	}

}



func main() {
	http.Handle("/api/now", http.HandlerFunc(getTime))
	http.Handle("/api/sort", http.HandlerFunc(sortArray))
	http.Handle("/api/weather", http.HandlerFunc(getWeather))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
