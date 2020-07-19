# API文档
## Basic Response Body

序号  |参数|类型|规则|
---- | -----| -----|-----|
1|code|number|正常响应200|
2|message|string|{...}|
3|data|object|实例或null|

最好用postman

## 1.注册

**1.获取邮箱验证码**

- URL：/getcode

- method:GET

- Request querystring

    ```
    email:lsh1987231487@126.com
    ```

- Response Body

        {
           "mes": "验证码发送成功"
        }

​     作用:向你的邮箱发送验证码

**2.注册**

- URL：/register

- method:GET

- Request querystring

  ```
  name:123456     //用户名
  password:123456  //密码
  code:844999   //上面的验证码 
  level:0  //权限等级
  ```

- Response Body

      {
          "code": 200,
          "data": {},
          "msg": "ok"
      }

# 2.登录

- URL：/login

- method:GET

- Request querystring

  ```
  username:123456
  password:123456
  ```

- Response Body

  ```
  {
      "code": 200,
      "data": {
       "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMzQ1NiIsImxldmVsIjowLCJleHAiOjE1OTUxNjgxNTcsImlzcyI6ImJsb2cifQ.cfVH62iOp48U3zhdmKvlPIDMUX4ig7j8GPUd94e9lKI"
      },
      "msg": "ok"
  }{
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#   3.文章的增删改查

# 注:下面的接口除第一个都默认传了token

#### 1.获取全部文章

- URL：/level/0/articles

- method:GET

- Request querystring

  ```
  token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMzQ1NiIsImxldmVsIjowLCJleHAiOjE1OTUxNjgxNTcsImlzcyI6ImJsb2cifQ.cfVH62iOp48U3zhdmKvlPIDMUX4ig7j8GPUd94e9lKI
  ```

- Response Body

  ```
  {
      "code": 200,
      "data": {
          "lists": [
              {
                  "ID": 1,
                  "CreatedAt": "2020-07-19T07:09:15Z",
                  "UpdatedAt": "2020-07-19T07:12:05Z",
                  "DeletedAt": null,
                  "user_ld": 1,
                  "title": "567dawfsvsv",
                  "text": "l  w ad wad w",
                  "state": 0,
                  "create_by": "1987231487",
                  "update_by": "1987231487"
              },
              {
                  "ID": 2,
                  "CreatedAt": "2020-07-19T07:10:00Z",
                  "UpdatedAt": "2020-07-19T09:09:37Z",
                  "DeletedAt": null,
                  "user_ld": 1,
                  "title": "567",
                  "text": "hahhah",
                  "state": 0,
                  "create_by": "1987231487",
                  "update_by": ""
              }
          ],
          "total": 2
      },
      "msg": "ok"
  }
  ```

#### 2.获取指定文章



- URL：/level/0/articles/:id

- method:GET

- Request param

  ```
  例:/level/0/articles/1
  ```

- Response Body

  ```
  {
      "code": 200,
      "data": {
          "ID": 1,
          "CreatedAt": "2020-07-19T07:09:15Z",
          "UpdatedAt": "2020-07-19T07:12:05Z",
          "DeletedAt": null,
          "user_ld": 1,
          "title": "567dawfsvsv",
          "text": "l  w ad wad w",
          "state": 0,
          "create_by": "1987231487",
          "update_by": "1987231487"
      },
      "msg": "ok"
  }
  ```

#### 3.新建文章 //level等级为1可用(level在token中获取)





- URL：/level/1/articles

- method:POST

- Request querystring

  ```
  user_id:1
  title:hhh
  text:123
  state:0
  create_by:1987231487
  ```

- Response Body

  ```
  {
      "code": 10001,
      "data": {},
      "msg": "用户权限不足"
  }
  ```

或

- 

  ```
  {
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#### 4.更新指定文章//level等级为1可用(level在token中获取)

- URL：/level/1/articles/:id

- method:PUT

- Request param

  ```
  id:1
  ```

- Request querystring

  ```
  user_id:1
  title:567dawfsvsv
  text:l  w ad wad w
  state:1
  updated_by:1987231487
  token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjE5ODcyMzE0ODciLCJsZXZlbCI6MSwiZXhwIjoxNTk1MTY1MzI2LCJpc3MiOiJibG9nIn0.HekfcJ-ST0O5nRVnNHFUFv5bCTOuQDS2zkHKNGBSTuM
  ```

  

- Response Body

  ```
  {
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#### 

#### 5.删除指定文章//level等级为1可用(level在token中获取)

- URL：/level/1/articles/:id

- method:DELETE

- Request param

  ```
  id:1
  ```

  

- Response Body

  ```
  {
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#### 

#### 6.设置文章是否可被评论//level等级为1可用(level在token中获取)

- URL：/level/1/articles/:id

- method:POST

- Request param

  ```
  id:1
  ```

  Request querystring

  ```
  state:1
  ```

  //state设置为1,文章不可评论

- Response Body

  ```
  {
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#### 

# 4.评论增删改查

#### 1.获取文章所有评论

- URL：/level/0/comment/:id

- method:GET

- Request param

  ```
  id:1
  ```

- Response Body

  ```
  {
      "code": 200,
      "data": {
          "comments": [
              {
                  "ID": 4,
                  "CreatedAt": "2020-07-19T08:53:02Z",
                  "UpdatedAt": "2020-07-19T08:53:02Z",
                  "DeletedAt": null,
                  "user_id": 1,
                  "article_id": 1,
                  "text": "nmsl",
                  "create_by": "1987231487"
              }
          ]
      },
      "msg": "ok"
  }
  ```

#### 

#### 2.给文章添加评论

- URL：/level/0/comment/:id

- method:POST

- Request param

  ```
  id:1
  ```

  Request querystring

  ```
  user_id:1
  text:你们上来啊啊啊啊啊啊
  created_by:1987231487
  ```

  - Response Body

  ```
  {
      "code": 10004,
      "data": {},
      "msg": "文章不可评论"
  }
  ```

或

- 

  ```
  {
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#### 

#### 3.删除评论//level等级为1可用(level在token中获取)

- URL：/level/1/comment/:id

- method:DELETE

- Request param

  ```
  id:1
  ```

- Response Body

  ```
  {
      "code": 200,
      "data": {},
      "msg": "ok"
  }
  ```

#### 

#### 