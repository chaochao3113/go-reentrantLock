/**
* @Author: Chao
* @Date: 2022/5/3 18:13
* @Version: 1.0
 */

package reentrantlock

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func TestGetGId(t *testing.T) {
	buf := [64]byte{}
	s := buf[:runtime.Stack(buf[:], false)]
	s = s[len("goroutine "):]
	s = s[:bytes.IndexAny(s, " ")]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	fmt.Println(gid)
}
