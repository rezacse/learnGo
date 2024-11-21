package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	//"math"
	"io"
	"time"
)


func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	download(url)

	fileName := "./fromString.txt"
	writeFile(fileName)
	defer readFile(fileName)

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)


	states := make(map[string]string)
	states["WA"] = "Washington"
	states["NY"] = "New York"
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println(states)

	delete(states, "WA")
	fmt.Println(states)

	n := time.Date(2024, 11, 10, 03, 0,0,0, time.UTC);
	fmt.Println(n.Format(time.ANSIC))

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter text:")
	input, _ := reader.ReadString('\n');
	fmt.Println("You entered:" , input);
	
	fmt.Println("Enter number:")
	numInput, _ := reader.ReadString('\n');
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:" , aFloat);
	}
}

// func add(v1 int, v2 int) int {
// 	return v1 + v2
// }

// func adds(vs ...int) int {
// 	 r := 0
// 	for _, v := range vs {
// 		r += v
// 	}
// 	return r
// }

func download(url string) {
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	fmt.Println("content", content)

	tours := tours2FromJson(content)
	for _, tour := range tours {
		fmt.Println(tour)
	}
	
	tours = toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}

	return tours
}
func tours2FromJson(content string) []Tour {
	var tours []Tour

	err := json.Unmarshal([]byte(content), &tours)
	checkError(err)

	return tours
}

func writeFile(fileName string) {
	content := "Hello Go"
	file, err := os.Create(fileName)
	checkError(err)
	length, _err := io.WriteString(file, content)
	checkError(_err)
	fmt.Printf("File content len %v\n", length)

	defer file.Close()
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println(data)
}

func checkError(err error) {
 if err != nil {
	fmt.Println(err)
	panic(err)
 }
}

type Dog struct {
	Bread string
	Weight int
}

func (d Dog) Speak() {
	fmt.Println(d.Bread)
}

type Tour struct {
	Name, Price string
}