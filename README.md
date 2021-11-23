# golib

this is a go lib.

### jlog

```go

//example1
```go

timeUtil:=BuildTimeUtils()


```

### jhttp

```go
//example1
client := jhttp.CreateGet(url, timeout)
if client == nil {
	fmt.Println("ERR:Client NIL")
	return
}
client.SetHeader("User-Agent", strNowUA)
client.AddCookie("name", value, "/", domain)
body, err := client.Do()
if err != nil {
	return ""
}

......

//example2

client := jhttp.CreateGet("https://v.sogou.com/?forceredirect=2&ie=utf8", timeout)

if client == nil {
	fmt.Println("ERR:Client null")
	return
}

client.SetHeader("User-Agent", ua)

resp, err := client.DoWithResp()

if err != nil {
	fmt.Println("ERR:" + err.Error())
	return
}

for _, cookie := range resp.Response.Cookies() {
    ......
}

```

### jbase62

```go

en:=NewEncoding()
s:=en.ToBase62(79876545352)
i:=en.FromBase62("uyiouj8")


//example
	t := jtime.New()

	en := NewEncoding()

	ok := 0
	err := 0

	fmt.Println(t.GetDateTime())

	for i := 10000000; i > 0; i-- {

		s := en.ToBase62(int64(i))

		i1 := en.FromBase62(s)

		if i1 == int64(i) {
			ok++
		} else {
			err++
		}
	}

	fmt.Println(t.GetDateTime())

	fmt.Printf("OK:%d,ERR:%d\n", ok, err)

Test Result:

PS F:\yjh201960613\yjhgo\golib\test> go run main.go
2021-06-30 15:38:45
2021-06-30 15:38:51
OK:10000000,ERR:0

```
