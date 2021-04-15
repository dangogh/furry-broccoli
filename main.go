package main

import "net/http"

const MTBA_API = "https://api-v3.mbta.com"

func get(path string, attrs map[string]string) {
	http.Get(MTBA_API + "/" + path)
}
func main() {

}
