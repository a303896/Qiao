package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	rdb  *redis.Client
	once sync.Once
)

func Instance() *redis.Client {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // 没有密码，默认值
			DB:       0,  // 默认DB 0
		})
	})
	return rdb
}

func TransactionTest() {
	rdb := Instance()
	ctx := context.Background()
	rdb.Set(ctx, "k1", 10, time.Second*120)

	wg := sync.WaitGroup{}
	wg.Add(15)
	for i := 0; i < 15; i++ {
		go func() {
			defer wg.Done()
			stock, _ := rdb.Get(ctx, "k1").Int()
			if stock > 0 {
				rdb.Decr(ctx, "k1")
			}
			fmt.Printf("stock: %d\n", stock)
		}()
	}
	wg.Wait()
}

// 事务版本
func TransactionTest2() {
	rdb := Instance()
	ctx := context.Background()
	rdb.Set(ctx, "k1", 10, time.Second*120)

	wg := sync.WaitGroup{}
	wg.Add(15)
	for i := 0; i < 15; i++ {
		go func() {
			defer wg.Done()
			stock, _ := rdb.Get(ctx, "k1").Int()
			if stock > 0 {
				err := rdb.Watch(ctx, func(tx *redis.Tx) error {
					_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
						pipe.Decr(ctx, "k1")
						return nil
					})
					return err
				}, "k1")
				if err != nil {
					fmt.Printf("transaction fail:%v\n", err)
				}
			}
			fmt.Printf("stock: %d\n", stock)
		}()
	}
	wg.Wait()
}
