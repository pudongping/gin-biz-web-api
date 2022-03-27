# 接口请求示例

## 请求用户列表

### 接口

```shell
curl --location --request GET '0.0.0.0:3000/api/auth/user?page=3&per_page=2&order_by=id,desc|created_at,asc' 
```

### 返回值示例

```json

{
    "code": 0,
    "data": {
        "paginate": {
            "page": 3,
            "per_page": 2,
            "total_page": 5,
            "total_count": 10,
            "prev_page_url": "0.0.0.0:3000/api/auth/user?per_page=2&order_by=id,desc|created_at,asc&page=2",
            "next_page_url": "0.0.0.0:3000/api/auth/user?per_page=2&order_by=id,desc|created_at,asc&page=4"
        },
        "users": [
            {
                "id": 12,
                "account": "41234535",
                "email": "411113@12.com",
                "phone": "",
                "password": "$2a$14$5zOYI.wV8Wi7CIh7w0fdpOR4vAQ6KeKsz6TXbhF2MPfQ8NGM8CX4m",
                "province": "",
                "city": "",
                "country": "",
                "nickname": "",
                "introduction": "",
                "avatar": "",
                "created_at": 1647271760,
                "updated_at": 1647271760
            },
            {
                "id": 11,
                "account": "31234535",
                "email": "311113@12.com",
                "phone": "",
                "password": "$2a$14$5zOYI.wV8Wi7CIh7w0fdpOR4vAQ6KeKsz6TXbhF2MPfQ8NGM8CX4m",
                "province": "",
                "city": "",
                "country": "",
                "nickname": "",
                "introduction": "",
                "avatar": "",
                "created_at": 1647271760,
                "updated_at": 1647271760
            }
        ]
    },
    "msg": "请求成功"
}

```