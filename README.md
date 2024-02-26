# oneid-jwt-auth-cli
oneid jwt auth command line tool

usage: oneid-jwt-auth auth --config ./config.json

1. 下载[工具](bin/oneid-jwt-auth): mac版本
2. 启动一个terminal, 设置可执行权限: chmod a+x ./oneid-jwt-auth
3. 下载[配置文件模版](./config.json): 并把自己的配置更新至该文件, 注: 该模版支持通过
    两种方式来指定用户信息, 即"user" 或者 "claim", 优先使用claim
    - claim方式: 用户指定一个key/value形式的map, 生成的免登token用用户指定的key/value, 比如:
    ```
    "claim": {
        "id": "f99530d4-8317-4900-bd02-0127bb8c44de",
        "name": "张三",
        "email": "zhangsan@example.com"
    }
    ```
    - user方式: 遵循[OpenID中id_token标准claim](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)来生成免登token, 比如:
    ```
    "user": {
        "id": "f99530d4-8317-4900-bd02-0127bb8c44de",
        "username": "zhangsan",
        "name": "张三",
        "email": "zhangsan@example.com",
        "mobile": "+86 13411112222",
        "extension": {
        "picture": "https://www.example.com/avatar1.png"
        }
    }
    ```

   

