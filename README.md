# redis-struct

```go
package main

import (
    "context"
    "fmt"
    "time"

    rj "github.com/pedidopago/redis-struct/json"
    redistruct "github.com/pedidopago/redis-struct"
    "github.com/go-redis/redis/v8"
)

type MyData struct {
    Name string `json:"name"`
}

func main() {

    rediscl := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
	})

    cl := rj.New(rediscl)

    var dst MyData
    if err := cl.Get(context.TODO(), "teste", &dst); err == redistruct.ErrNotFound {
        fmt.Println("not found (expected)")
    }
    dst.Name = "Hello"
    _ = cl.Set(context.TODO(), "teste", dst, time.Second*10)
    time.Sleep(time.Second * 2)
    cl.Del(context.TODO(), "teste")
}
```