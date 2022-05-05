/**
* @Author: Chao
* @Date: 2022/5/3 18:33
* @Version: 1.0
 */

package main

import (
	"fmt"
	"reentrantLock/reentrantlock"
)

func main() {
	lock := reentrantlock.NewReentrantLock()

	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock()
			fmt.Println("lock", i)
		}

		for i := 0; i < 10; i++ {
			lock.Unlock()
			fmt.Println("unlock", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock()
			fmt.Println("lock", i)
		}

		for i := 0; i < 10; i++ {
			lock.Unlock()
			fmt.Println("unlock", i)
		}
	}()

	select {}
}
