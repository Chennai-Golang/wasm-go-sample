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

func add(this js.Value, i []js.Value) interface{} {
	value1 := getValue(i[0])
	value2 := getValue(i[1])

	fmt.Println("Adding values: ", value1, " + ", value2)

	setValue(value1+value2, i[2])

	return nil
}

func subtract(this js.Value, i []js.Value) interface{} {
	value1 := getValue(i[0])
	value2 := getValue(i[1])

	fmt.Println("Subtracting values: ", value1, " + ", value2)

	setValue(value1-value2, i[2])

	return nil
}

func getKeys(this js.Value, urls []js.Value) interface{} {
	resp, _ := http.Get(urls[0].String())
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("getKeys", js.FuncOf(getKeys))
}

func main() {
	ch := make(chan (struct{}))

	fmt.Println("Go WebAssembly Initialized 2")
	registerCallbacks()

	<-ch
}
