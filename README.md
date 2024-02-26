# oneid-jwt-auth-cli
oneid jwt auth command line tool

usage: oneid-jwt-auth auth --config ./config.json

1. 下载[工具](bin/oneid-jwt-auth): mac版本
2. 启动一个terminal, 设置可执行权限: chmod a+x ./oneid-jwt-auth
3. 下载[配置文件模版](./config.json): 并把自己的配置更新至该文件
    - 遵循[OpenID中id_token标准claim](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)来生成免登token, 比如:
    ```
    "user": {
        "sub": "f99530d4-8317-4900-bd02-0127bb8c44de",
        "preferred_username": "zhangsan",
        "name": "张三",
        "email": "zhangsan@example.com",
        "phone_number": "+86 13411112222"
    }
    ```

   

