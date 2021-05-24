/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:07 AM  24/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package mic_test

import (
	"expvar"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", expvar.Handler())
}