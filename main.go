package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"syscall/js"
)

func getValue(fieldID js.Value) int {
	value := js.Global().Get("document").Call("getElementById", fieldID.String()).Get("value").String()

	intVal, _ := strconv.Atoi(value)

	return intVal
}

func setValue(value int, fieldID js.Value) {
	js.Global().Get("document").Call("getElementById", fieldID.String()).Set("value", value)
}

func add(i []js.Value) {
	value1 := getValue(i[0])
	value2 := getValue(i[1])

	setValue(value1+value2, i[2])
}

func subtract(i []js.Value) {
	value1 := getValue(i[0])
	value2 := getValue(i[1])

	setValue(value1-value2, i[2])
}

func getKeys(urls []js.Value) {
	resp, _ := http.Get(urls[0].String())
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
	js.Global().Set("getKeys", js.NewCallback(getKeys))
}

func main() {
	ch := make(chan (struct{}))

	fmt.Println("Go WebAssembly Initialized 2")
	registerCallbacks()

	<-ch
}
