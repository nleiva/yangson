# yangson

## Synopsis

Add and removes indentation to **YANG** config files **JSON** encoded. 

## Code Example

`prettyprint` takes a compact JSON config and pretty prints it.
- input: [compact JSON](files/bgp.json)
- output: [pretty printed JSON](files/bgp_formatted.json)

```go
	b, err := ioutil.ReadFile("files/bgp.json")
	if err != nil {
		fmt.Print(err)
	}

	b, _ = prettyprint(b)
```

`unprettyprint` takes an indented JSON config and pretty prints it.
- input: [pretty printed JSON](files/bgp_formatted.json)
- output: [compact JSON](files/bgp.json)

```go
	b, err := ioutil.ReadFile("files/bgp_formatted.json")
	if err != nil {
		fmt.Print(err)
	}

	b, _ = unprettyprint(b)
```

## Where are we taking these JSON config files from?

- Check out this great example: [YDK bgp_xr_json](https://github.com/CiscoDevNet/ydk-gen/tree/master/sdk/cpp/samples/bgp_xr_json)
- We will include further YDK samples here: [ydk-cpp-samples](ydk-cpp-samples)

## To Do list

- Incorporate commit/push configs on ([github.go](github.go))
- Add suport to clone or Wget/cURL a config ([github.go](github.go))


## See it in action

- [Go Playground](https://play.golang.org/p/TUdaIF_NBp)