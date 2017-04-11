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

	input := "files/config_read_ifmgr.json"
	output := "files/config_read_ifmgr.json_formatted.json"

	b, err := ioutil.ReadFile(input) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//Add indentation
	b, _ = prettyprint(b)
	//Remove indentation
	//b, _ = unprettyprint(b)

	err = ioutil.WriteFile(output, b, 0644)
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
