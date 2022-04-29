package json

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	redistruct "github.com/pedidopago/redis-struct"
	"github.com/stretchr/testify/require"
)

type TestData struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func TestJSONClient(t *testing.T) {
	dbint, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	cl := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		DB:       dbint,
		Password: os.Getenv("REDIS_PASSWORD"),
		Username: os.Getenv("REDIS_USERNAME"),
	})
	defer cl.Close()

	titem := TestData{
		Name:  "John Doe",
		Age:   30,
		Score: 48.5,
	}

	jcl := New(cl)
	require.NoError(t, jcl.Set(context.Background(), "rj_test_k1", titem, time.Second*3))
	var titem2 TestData
	require.NoError(t, jcl.Get(context.Background(), "rj_test_k1", &titem2))
	require.Equal(t, titem.Name, titem2.Name)
	require.Equal(t, titem.Age, titem2.Age)
	require.Equal(t, titem.Score, titem2.Score)
	time.Sleep(time.Second * 4)
	require.Equal(t, redistruct.ErrNotFound, jcl.Get(context.Background(), "rj_test_k1", &titem2))
}
