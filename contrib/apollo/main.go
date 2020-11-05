package apollo

import (
	"fmt"
	"time"

	// "errors"
	"context"

	"github.com/etcd-io/etcd/clientv3"

	jsoniter "github.com/json-iterator/go"
)

var (
	conn           *clientv3.Client
	cjson          = jsoniter.ConfigCompatibleWithStandardLibrary
	dialTimeout    = 5 * time.Second
	requestTimeout = 8 * time.Second
)

func New(endpoints []string) {

	var err error
	conn, err = clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   endpoints,
	})

	if err != nil {
		panic(err)
	}
}

func Close() {
	conn.Close()
}

func Parse(key string, v interface{}) error {

	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	kv := clientv3.NewKV(conn)
	gr, _ := kv.Get(ctx, key)
	if gr == nil || len(gr.Kvs) == 0 {
		return fmt.Errorf("No more " + key)
	}

	return cjson.Unmarshal(gr.Kvs[0].Value, v)
}
