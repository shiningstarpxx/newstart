/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:07 AM  24/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package main

import (
	"expvar"
	"net/http"
	"time"
)

func main() {
	lastAccess := expvar.NewString("lastAccess")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
		lastAccess.Set(time.Now().String())
	})

	http.ListenAndServe(":8080", nil)
}

// http :8080/debug/vars | grep lastAccess
// 在 expvar.go 中有init，注册了debug/vars，会发布所有Publish的变量
/*
func init() {
        http.HandleFunc("/debug/vars", expvarHandler)
        Publish("cmdline", Func(cmdline))
        Publish("memstats", Func(memstats))
}
 */