/*
 * @Descripttion: main函数调用的用户处理API，包括注册，登录
 * @version: 10
 * @Author: KongJHong
 * @Date: 2019-08-03 14:36:07
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-03 22:07:42
 */
package handler

import(
	"net/http"
	"io/ioutil"
	"filestore-server/util"
	dblayer "filestore-server/db"
	"fmt"
	// "io"
	"time"
	//"strconv"

)

const (
	pwdSalt = "*#890"
	tokenSalt = "_tokensalt"
	tokenExpire int64 = 604800	//token超时
)

//SignupHandler 处理用户注册请求
func SignupHandler(w http.ResponseWriter,r *http.Request){

	//运行两个逻辑
	//1.如果是get方法，则返回一个注册页面
	//2.如果是post方法，则处理注册事件

	if r.Method == http.MethodGet{
		data,err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	//2.校验参数的有效性
	if len(username) < 3 || len(passwd) < 5{
		w.Write([]byte("Invalid parameter"))
		return
	}

	//3. 密码加盐-加密
	encPasswd := util.Sha1([]byte(passwd+pwdSalt))
	
	//4. 将用户信息注册到用户表中
	suc := dblayer.UserSignup(username, encPasswd)
	if suc{
		w.Write([]byte("SUCCESS"))
	}else{
		w.Write([]byte("FAILED"))
	}
}

//SignInHandler 登录接口，没有设置路由，而是通过静态地址访问
func SignInHandler(w http.ResponseWriter,r *http.Request){

	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	encPasswd := util.Sha1([]byte(password+pwdSalt))
	
	//1.校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked{
		w.Write([]byte("FAILED"))
		return
	}

	//2.生成访问凭证(token)---生成一个40位字符的token
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username,token)
	if !upRes{
		w.Write([]byte("FAILED"))
		return
	}
	//3.登录成功后，重定向到首页
	//w.Write([]byte("http://"+r.Host+"/static/view/home.html"))

	//除了登录和注册，访问其他页面都是需要token的
	// resp := util.NewRespMsg(0, "OK", struct{
	// 	Location string
	// 	Username string
	// 	Token  string
	// }{
	// 	Location : "http://"+r.Host+"/static/view/home.html",
	// 	Username : username,
	// 	Token : token,
	// })

	resp := util.RespMsg{
		Code : 0,
		Msg :"OK",
		Data : struct{
			Location string
			Username string
			Token 	 string
		}{
			Location: "http://"+r.Host+"/static/view/home.html",
			Username: username,
			Token:	  token,
		},
	}
	
	w.Write(resp.JSONBytes())
	fmt.Println("传输完毕")
}

//UserInfoHandler 返回用户信息
func UserInfoHandler(w http.ResponseWriter,r *http.Request){
	//1.解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	//token := r.Form.Get("token")

	//2.验证Token是否有效(加了拦截器，这里可以注释掉)
	// isValidToken := IsTokenValid(username,token)
	// if !isValidToken{
	// 	w.WriteHeader(http.StatusForbidden)	//403
	// 	return
	// }
	
	//3.查询用户信息
	userInfo,err := dblayer.GetUserInfo(username)
	if err != nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//4.组装并且响应用户数据
	resp := util.RespMsg{
		Code : 0,
		Msg :"OK",
		Data :userInfo,
	}

	w.Write(resp.JSONBytes())
}


//GenToken 生成一个40位字符的token
func GenToken(username string)string{
	//40位字符:md5(username+timestamp+tokenSalt) + timestamp[:8]
	ts := fmt.Sprintf("%x",time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username+ts+tokenSalt))
	return tokenPrefix + ts[:8]
}

//IsTokenValid 判断token是否失效
func IsTokenValid(username,token string) bool {
	//fmt.Println(token)
	if len(token) != 40{
		fmt.Println("token过长")
		return false
	}

	//TODO: 判断token的时效性，是否过期 可以取出token的后8位判断日期是否大于多少天来判断是否过期
	
	

	// nowTs := time.Now().Unix()
	// duration := nowTs - int64(ts)
	// if  duration > tokenExpire{
	// 	fmt.Println("过期")
	// 	return false
	// }

	//TODO: 从数据库表tbl_user_token查询username对应的token信息
	//TODO: 对比两个token是否一致
	// userToken := dblayer.GetUserToken(username)
	// if userToken != token{
	// 	fmt.Println("token不相等")
	// 	return false
	// }
	
	return true
}