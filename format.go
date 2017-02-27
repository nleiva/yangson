package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	//Option to read from static string
	//b := []byte(`{"Cisco-IOS-XR-ipv4-bgp-cfg:bgp":"empty"}`)

	//Option to read from file
	b, err := ioutil.ReadFile("files/bgp.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//Add indentation
	b, _ = prettyprint(b)
	//Remove indentation
	//b, _ = unprettyprint(b)

	//Print one of the two options for debugging
	//fmt.Printf("%s", b)

	//
	err = ioutil.WriteFile("files/bgp_formatted.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func unprettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Compact(&out, b)
	return out.Bytes(), err
}
