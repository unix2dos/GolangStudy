#### flag

```
var test = flag.String("name", "default", "only test")

func main() {
	flag.Parse() //一定要Parse()
	fmt.Println(*test)
}
```


#### json

Unmarshal

```
	type Test struct {
		Key1 string
		Key2 string
	}

	byte1 := []byte(`{"key1":"value1", "key2":"value2"}`)
	var test Test
	json.Unmarshal(byte1, &test)
	fmt.Println(test.Key1)
	fmt.Println(test.Key2)
```

marshal

```
	type Test struct {
		Key1 string
		Key2 string
	}

	test := &Test{Key1: "a", Key2: "b"}
	byte2, _ := json.Marshal(test)
	fmt.Println(string(byte2))
```
#### convert
int -> string :

```
    strconv.Itoa(value)
```
