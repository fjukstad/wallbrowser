package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/fjukstad/gowebsocket"
)

var config map[string]string

var indexTemplate = template.Must(template.ParseFiles("templates/index.html"))
var controllerTemplate = template.Must(template.ParseFiles("templates/controller.html"))

func renderTemplate(t *template.Template, w http.ResponseWriter, d interface{}) {

	// Cross domain requests in browser pow pow!
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Applies template to dataobject d
	err := t.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(indexTemplate, w, nil)
}

func ControllerHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(controllerTemplate, w, nil)
}

// Get x,y for hostname
func coordinates(hostname string) (int, int) {
	a := strings.Split(hostname, "-")
	x, _ := strconv.Atoi(a[1])
	y, _ := strconv.Atoi(a[2])
	return x, y
}

// translate hostname to real offset pixel x y.
// Remember that 0,0 is top left on the wall, while tile-0-0 is bottom left
func wallcoordinates(hostname string) (int, int) {

	x, y := coordinates(hostname)

	projectorX := 1024
	projectorY := 768

	wallY := projectorY * 3

	return x * projectorX, wallY - (y * projectorY)

}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received location request from ", r.RemoteAddr)

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Panic(err)
	}

	hostname := config[ip]
	if hostname == "" {
		hostname = "tile-0-2"
		log.Println(ip, "is not a part of the display wall!")
		//fmt.Fprintf(w, "You're not a part of the display wall!")
		//fmt.Fprintf(w, "200,300")
		//return
	}
	x, y := wallcoordinates(hostname)

	fmt.Fprintf(w, "%d,%d", x, y)
}

// Reads config file and returns map ip -> hostname
func ReadConfig(filename string) map[string]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	config := make(map[string]string)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}

		hostname := record[0]
		ip := record[1]

		config[ip] = hostname
	}

	return config

}
func main() {

	// read config
	config = ReadConfig("config.csv")

	// set up websocket hub for messages etc
	ip := "10.1.1.60"
	wsport := ":9192"
	server := gowebsocket.New(ip, wsport)
	server.Start()

	// http handlers
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/location", LocationHandler)
	http.HandleFunc("/controller", ControllerHandler)

	port := ":9191"

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Panic(err)
	}

}
