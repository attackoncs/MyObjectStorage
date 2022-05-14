package objects

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method //判断请求方法时get还是post,并调用objects包中相应函数，否则返回方法不支持
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
