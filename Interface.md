## 云中台服务接口

### 用户注册

`url`  

/cp-ua/v1/user/register

`method`  

post

`formdata`

- phone 用户手机号, 必须
- name 用户姓名，与gogs的用户名一致， 必须
- email 用户邮箱，与gogs的用户邮箱一致， 必须 
- password 用户密码，与gogs的密码一致，必须

`response`

```
{
  "code": 0,
  "info": "Success"
}
```