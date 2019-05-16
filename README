 # Golang server

  The project is to demonstrate basic knowledge and skills of GOlang/HTTP requests usage


##Launch and installation stage


 ####**Make sure you have GO distrib installed**
	       
 Otherwise, proceed to [Golang.org](https://golang.org/doc/install) and download the **distribution package**.

		        
Then type ```go run server.go``` in the project folder.

##Usage
####  All the requests could be either run in the browser, or directly from the console using ```curl``` command
``` curl "http://localhost:8080/api/now" ``` - get server's UTC time.


``` curl -X POST -d "{\"array\": [], \"uniq\":}" http://localhost:8080/api/sort ``` - returns an array, sorted in ascending order.
You could fill the **"array"** parameter with integers/floats(**there 0 < length <= 100**) and set the **"uniq"** parameter to ```true``` to remove duplicates **or** not specify it at all.



``` curl "http://localhost:8080/api/weather?city=NAME" ``` - get the temperature**(in Fahrenheits**) of the chosen city from the [OpenWeatherMap API]( https://openweathermap.org/api). The city's name is specified with the ```NAME``` parameter.

____
Coded by Andrey Shibaev (https://github.com/DWhistle)
