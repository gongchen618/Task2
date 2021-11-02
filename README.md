#### github 的登录过程分析 

初始在 /login `Code:200`

POST /session `Code:302`
> 这是提交用户名和密码\
> 1, 新 set 了很多 cookie，比如
> user_session=blabla, logged_in=yes,
> dotcom_user=username\
> 2, 验证通过回复 location: github.com/login

GET /login `Code:302`
> 和初始的地址是一样的，区别在于 \
> 1, 有 location，故会立即跳转，状态码也不同 \
> 2, cookie 中 logged_in 的字段从 no 换成 yes\
> cookie user_session=blabla; 
>               dotcom_user=username;\
> 3, location: github.com

GET / `Code:200`

POST /logout
> 退出登录时重新 set 了很多 cookie

其他：
- location 和 referer 是有联系的
- `_gs_sess` 字段是一个很长很奇怪的 cookie 字符串 
，它每次 response 都被重新 set，但 request
中返回的又不是上一次被 set 的字符串
