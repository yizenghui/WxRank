// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import "testing"
import "fmt"

func Test_Post(t *testing.T) {
	Post("https://mp.weixin.qq.com/s/e_0BhJ0vOvIqt610MjgKLA")
	Post("https://mp.weixin.qq.com/s/lu9yU_Wvd8yQs6alZ7OBqg")
	Post("https://mp.weixin.qq.com/s/widHAGN7zOymivIUihkuTw")
	Post("https://mp.weixin.qq.com/s/wXOhpZHdzlHlf5jgC54AaA")
}

func Test_Hot(t *testing.T) {
	articles, _ := Hot()
	fmt.Println(articles)
}

func Test_New(t *testing.T) {
	articles, _ := New()
	fmt.Println(articles)
}
