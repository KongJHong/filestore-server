/*
 * @Descripttion: http请求拦截器
 * @version: 11
 * @Author: KongJHong
 * @Date: 2019-08-03 21:24:20
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-03 21:26:24
 */
 
package handler

import (
	"net/http"
)

//HTTPInterceptor http请求拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")
		
			if len(username) < 3 || !IsTokenValid(username, token){
				w.WriteHeader(http.StatusForbidden)
				return
			}
			h(w,r)
		})
}