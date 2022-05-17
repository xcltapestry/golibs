package redisx_test

/**
 * Copyright 2022 golibs Author. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @Project golibs
 * @Description
 * @author XiongChuanLiang<br/>(xcl_168@aliyun.com)
 * @license http://www.apache.org/licenses/  Apache v2 License
 * @version 1.0
 */

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bsm/redislock"
	"github.com/xcltapestry/golibs/driver/redisx"
)

func connRedis() (*redisx.Client, error) {
	rdb := redisx.NewClient(&redisx.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	err := rdb.Ping()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

// go test -run Example
func Example() {

	rdb, err := connRedis()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rdb.Close()

	// --
	v2, err2 := rdb.Set("key3", "aaaa", 1*time.Second)
	fmt.Println("[cmd-Set] v:", v2, " err:", err2)

	v3, err3 := rdb.SetEX("key3-ex", "aaaa", 1*time.Second)
	fmt.Println("[cmd-SetEX] v:", v3, " err:", err3)

	v4, err4 := rdb.SetNX("key3-nx", "aaaa", 1*time.Second)
	fmt.Println("[cmd-SetNX] v:", v4, " err:", err4)

	// --
	ttlv5, ttlerr5 := rdb.TTL("key3-nx")
	fmt.Println("[cmd-Expire-TTL-1] ttlv5:", ttlv5, " ttlerr5:", ttlerr5)

	v5, err5 := rdb.Expire("key3-nx", 5*time.Second)
	fmt.Println("[cmd-Expire] v:", v5, " err:", err5)

	ttlv5, ttlerr5 = rdb.TTL("key3-nx")
	fmt.Println("[cmd-Expire-TTL-2] ttlv5:", ttlv5, " ttlerr5:", ttlerr5)

	// --
	v, err := rdb.Exists("key2")
	fmt.Println("[cmd-Exists3] v:", v, " err:", err)

	v, err = rdb.Exists("key2", "key3")
	fmt.Println("[cmd-Exists1] v:", v, " err:", err)

	v, err = rdb.Exists("key3")
	fmt.Println("[cmd-Exists2] v:", v, " err:", err)

	kval, kerr := rdb.Get("key3")
	fmt.Println("[cmd-Get] kval:", kval, " err:", kerr)

	ttlv, ttlerr := rdb.TTL("key3")
	fmt.Println("[cmd-TTL] ttlv:", ttlv, " ttlerr:", ttlerr)

	// --
	crv, crerr := rdb.Incr("keyCR")
	fmt.Println("[cmd-Incr] crv:", crv, " crerr:", crerr)

	crv, crerr = rdb.Decr("keyCR")
	fmt.Println("[cmd-Decr] crv:", crv, " crerr:", crerr)

	crv, crerr = rdb.Incr("keyCR")
	fmt.Println("[cmd-Incr] crv:", crv, " crerr:", crerr)

}

// go test -timeout 30s -run TestLock
func TestLock(t *testing.T) {
	client, err := connRedis()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer client.Close()

	// go-redis + redislock
	lock := client.NewLocker()
	ctx := context.Background()
	// Try to obtain lock.
	err = lock.Obtain(ctx, "my-key", 100*time.Millisecond)
	if err == redislock.ErrNotObtained {
		fmt.Println("Could not obtain lock!")
	} else if err != nil {
		log.Fatalln(err)
	}

	// Don't forget to defer Release.
	defer lock.Release(ctx)
	fmt.Println("I have a lock!")

	// Sleep and check the remaining TTL.
	time.Sleep(50 * time.Millisecond)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl > 0 {
		fmt.Println("Yay, I still have my lock!")
	}

	// Extend my lock.
	if err := lock.Refresh(ctx, 100*time.Millisecond); err != nil {
		log.Fatalln(err)
	}

	// Sleep a little longer, then check.
	time.Sleep(100 * time.Millisecond)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl == 0 {
		fmt.Println("Now, my lock has expired!")
	}

}
