# oneid-jwt-auth-cli
oneid jwt auth command line tool

usage: oneid-jwt-auth auth --config ./config.json

1. 下载[工具](bin/oneid-jwt-auth): mac版本
2. 启动一个terminal, 设置可执行权限: chmod a+x ./oneid-jwt-auth
3. 下载[配置文件模版](./config.json): 并把自己的配置更新至该文件
    - 通过SDK中Userinfo来生成免登token, 比如:
    ```
    "user": {
        "id": "f99530d4-8317-4900-bd02-0127bb8c44de",
        "name": "张三",
        "username": "zhangsan",
        "email": "zhangsan@example.com",
        "mobile": "+86 13411112222"
    }
    ```

   

