# FamilyMoneyRecord

[toc]

# MySql 数据库设计

## 用户表（users）

| id              | username       | password            | name        | receipt_sum | disbursement_sum | advance_consumption |
| --------------- | -------------- | ------------------- | ----------- | ----------- | ---------------- | ------------------- |
| bigint unsigned | varchar(100)   | char(32)            | varchar(20) | double      | double           | double              |
| 自增主键        | 用户名（账号） | 密码（MD5加密存储） | 姓名        | 收入        | 支出             | 预消费金额          |

## 账单表（bills）

| id              | user_id         | receipt | disbursement | type        | note         | time     |
| --------------- | --------------- | ------- | ------------ | ----------- | ------------ | -------- |
| bigint unsigned | bigint unsigned | double  | double       | varchar(25) | varchar(200) | datetime |
| 自增主键        | 用户表外键      | 收入    | 支出         | 收支类型    | 账单备注     | 操作时间 |

## 证券账户表（accounts）

| id              | user_id         |
| --------------- | --------------- |
| bigint unsigned | bigint unsigned |
| 自增主键        | 用户表外键      |

## 股票操作表（operations）

| id              | account_id      | stock_id        | share_price        | buy_num  | sale_num | time     |
| --------------- | --------------- | --------------- | ------------------ | -------- | -------- | -------- |
| bigint unsigned | bigint unsigned | bigint unsigned | double             | int      | int      | datetime |
| 自增主键        | 证券账户表外键  | 股票表外键      | 交易时每股股票价格 | 买入股数 | 卖出股数 | 操作时间 |

## 股票持仓表（stocks）

| id              | account_id      | code        | position_num | profit       |
| --------------- | --------------- | ----------- | ------------ | ------------ |
| bigint unsigned | bigint unsigned | varchar(10) | int          | double       |
| 自增主键        | 证券账户表外键  | 股票代码    | 持仓股数     | 当前交易盈亏 |

# 接口设计

> Response 中：
> `code`为 0 时为正确处理，非 0 值皆为错误码。
> `msg`为处理信息。

## 用户管理

### 用户注册

- **HTTP Method**

​     [POST]

- **PATH**

​     /api/user/register

- **Request**

```json
{
    "username": "951753852456",    //string，注册账号
    "password": "123456789",    //string，用户密码
    "name": "张三",    //string，用户姓名
    "inviteUsername": "123123",    //string，邀请人账号
    "invitePassword": "321321"    //string，邀请人密码
}
```

- **Response**

```json
{
    "code": 0,    //int，状态码
    "msg": "注册成功"    //string，返回信息
}
```

> 注册成功后自动调用登录接口
> 注册失败则返回错误信息

### 用户注销 

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/user/logout

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "password": "123456789"    //string，用户密码
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "注销成功"    //string，返回信息
}
```

### 用户登录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/user/login

- **Request** 

```json
{
    "username": "951753852456",    //string，注册账号
    "password": "123456789"    //string，用户密码
}
```

- **Response** 

```json
{
    "data":
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
        "isAdmin": false    //bool，是否是管理员
    },
    "code": 0,    //int，状态码
    "msg": "登录成功"    //string，返回信息
}
```

### 获取用户信息

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/user/info

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

- **Response** 

```json
{
    "data":
    {
        "name": "qwer",    //string，用户名
        "username": "951753852456",    //string，用户账号
        "receiptSum": 500,    //double，收入总和
        "disbursementSum": 1000,    //double，支出总和
        "advanceConsumption": 1500    //double，预消费金额
    },
    "code": 0,    //int，状态码
    "msg": "获取成功"    //string，返回信息
}
```

### 修改用户信息

#### 修改用户姓名

- **HTTP Method** 

​     [PUT]

- **PATH** 

​     /api/user/modify/name

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "name": "qwer"    //string，新用户姓名
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "修改成功"    //string，返回信息
}
```

#### 修改用户密码

- **HTTP Method**

​     [PUT]

- **PATH** 

​     /api/user/modify/password

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "oldPassword": "qwer123...",    //string，用户旧密码
    "newPassword": "qwer456..."    //string，用户新密码
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "修改成功"    //string，返回信息
}
```

#### 修改用户预消费金额

- **HTTP Method**

​     [PUT]

- **PATH** 

​     /api/user/modify/advance_consumption

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "advanceConsumption": 1500    //double，用户预消费金额
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "修改成功"    //string，返回信息
}
```

## 管理员用户管理

> 管理员账户在配置文件中固定

### 搜索用户账号获取用户信息

（支持模糊查询）

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/admin/search_info

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "username": "951753852456"    //string，用户名
}
```

- **Response** 

```json
{
    "data":
    {
        "users": [
            {
                "name": "qwer",    //string，用户名
                "username": "951753852456",    //string，用户账号
                "receiptSum": 500,    //double，收入总和
                "disbursementSum": 1000,    //double，支出总和
                "advanceConsumption": 1500    //double，预消费金额
            },
            {
                "name": "abcd",    //string，用户名
                "username": "123456789",    //string，用户账号
                "receiptSum": 500,    //double，收入总和
                "disbursementSum": 1000,    //double，支出总和
                "advanceConsumption": 1500    //double，预消费金额
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "获取成功"    //string，返回信息
}
```

### 获取所有用户信息

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/admin/all_info

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

- **Response** 

```json
{
    "data":
    {
        "users": [
            {
                "name": "qwer",    //string，用户名
                "username": "951753852456",    //string，用户账号
                "receiptSum": 500,    //double，收入总和
                "disbursementSum": 1000,    //double，支出总和
                "advanceConsumption": 1500    //double，预消费金额
            },
            {
                "name": "abcd",    //string，用户名
                "username": "123456789",    //string，用户账号
                "receiptSum": 500,    //double，收入总和
                "disbursementSum": 1000,    //double，支出总和
                "advanceConsumption": 1500    //double，预消费金额
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "注册成功"    //string，返回信息
}
```

### 删除用户

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/admin/delete

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "username": "951753852456"    //string，用户用户名
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "删除成功"    //string，返回信息
}
```

### 添加用户

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/admin/add

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "username": "951753852456",    //string，用户用户名
    "password": "qwer123...",    //string，用户密码
    "name": "张三三"
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "添加成功"    //string，返回信息
}
```

### 修改用户密码

- **HTTP Method**

​     [PUT]

- **PATH** 

​     /api/admin/modify_password

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "username": "951753852456",    //string，用户用户名
    "password": "qwer123..."    //string，用户新密码
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "修改成功"    //string，返回信息
}
```

## 数据库管理

> 管理员管理数据库

### 清空数据库

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/database/empty

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "password": "1231321",    //string，管理员账户密码
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "清空成功"    //string，返回信息
}
```

### 备份数据库

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/database/save

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "name": "test"    //string，备份的数据库名称
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "备份成功"    //string，返回信息
}
```

### 获取数据库备份信息

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/database/get

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

- **Response** 

```json
{
    "data": 
    {
        "save": [
            {
                "name": "save1"    //string，备份数据库名称
            },
            {
                "name": "save2"
            },
            {
                "name": "save3"
            },
        ]    //备份的数据库信息
    },
    "code": 0,    //int，状态码
    "msg": "备份成功"    //string，返回信息
}
```

### 恢复数据库

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/database/recover

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "password": "1231321",    //string，管理员账户密码
    "name": "test"    //string，所要恢复到的数据库的名称
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "恢复成功"    //string，返回信息
}
```

### 删除数据库备份

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/database/delete

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "password": "1231321",    //string，管理员账户密码
    "name": "test"    //string，所要恢复到的数据库的名称
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "删除成功"    //string，返回信息
}
```

## 用户收支管理

### 添加收支记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/bill/add

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "receipt": 20,    //double，用户收入（该记录为支出记录则为0）
    "disbursement": 0,    //double，用户支出（该记录为收入记录则为0）
    "type": "学习",    //string，收支类型
    "note": "一支笔"    //string，备注信息
}
```

- **Response** 

```json
{
    "data":
    {
        "id": 2    //int，账单主键
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```

### 根据时间段获取收支记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/bill/get_by_time

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "beginTime": "2021-06-25 12:00:00",    //string，查询起始时间
    "endTime": "2021-06-29 12:00:00"    //string，查询结束时间
}
```

> 前端传输格式化的时间数据

- **Response** 

```json
{
    "data":
    {
        "records":[
            {
                "id": 1,    //int，收支记录id
                "receipt": 0,    //double，收入金额
                "disbursement": 15,    //double，支出金额
                "type": "衣食住行",    //string，收支类型
                "note": "午餐",    //string，备注
                "time": "2021-06-20 12:05:30"    //string，记录时间
            },
            {
                "id": 2,    //int，收支记录id
                "receipt": 0,    //double，收入金额
                "disbursement": 10,    //double，支出金额
                "type": "医疗",    //string，收支类型
                "note": "感冒",    //string，备注
                "time": "2021-06-20 12:05:35"    //string，记录时间
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "添加成功"    //string，返回信息
}
```

### 根据类型获取收支记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/bill/get_by_type

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "type": "衣食住行"
}
```

> 收入类型：工资、股票、分红、奖金
> 支出类型：税收、衣食住行、医疗、其他
> type字段中不允许出现其他类型

- **Response** 

```json
{
    "data":
    {
        "records":[
            {
                "id": 1,    //int，收支记录id
                "receipt": 0,    //double，收入金额
                "disbursement": 15,    //double，支出金额
                "type": "衣食住行",    //string，收支类型
                "note": "午餐",    //string，备注
                "time": "2021-06-20 12:05:30"    //string，记录时间
            },
            {
                "id": 2,    //int，收支记录id
                "receipt": 0,    //double，收入金额
                "disbursement": 10,    //double，支出金额
                "type": "衣食住行",    //string，收支类型
                "note": "晚餐",    //string，备注
                "time": "2021-06-20 12:05:35"    //string，记录时间
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```

### 根据类型获取收支记录总和

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/bill/get_sum_by_type

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

> 收入类型：工资、股票、分红、奖金
> 支出类型：税收、衣食住行、医疗、其他
> type字段中不允许出现其他类型

- **Response** 

```json
{
    "data":
    {
        "receipt": [20, 30, 40, 50],
        "disbursement": [40, 50, 10, 0]
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```

### 获取所有收支记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/bill/get_all

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

- **Response** 

```json
{
    "data":
    {
        "records":[
            {
                "id": 1,    //int，收支记录id
                "receipt": 0,    //double，收入金额
                "disbursement": 15,    //double，支出金额
                "type": "衣食住行",    //string，收支类型
                "note": "午餐",    //string，备注
                "time": "2021-06-20 12:05:30"    //string，记录时间
            },
            {
                "id": 2,    //int，收支记录id
                "receipt": 0,    //double，收入金额
                "disbursement": 10,    //double，支出金额
                "type": "衣食住行",    //string，收支类型
                "note": "晚餐",    //string，备注
                "time": "2021-06-20 12:05:35"    //string，记录时间
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```

### 删除收支记录

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/bill/delete

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "id": 2    //int，收支记录id
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "删除成功"    //string，返回信息
}
```

### 修改收支记录

- **HTTP Method** 

​     [PUT]

- **PATH** 

​     /api/bill/update

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "id": 2,    //int，收支记录id
    "receipt": 0,    //double，用户收入
    "disbursement": 30,    //double，用户支出
    "type": "衣食住行",    //string，类型
    "note": "午餐"    //string，备注
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "修改成功"    //string，返回信息
}
```

## 财务管理

### 证券账户管理

#### 创建证券账户

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/account/add

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

- **Response** 

```json
{
    "data": 
    {
        "id": 1    //int，证券账户主键
    },
    "code": 0,    //int，状态码
    "msg": "添加成功"    //string，返回信息
}
```

#### 删除证券账户

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/security/account/delete

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "id": 1,    //int，证券账户主键
    "password": "123456789"    //string，用户密码，再次确认，防止他人恶意操作
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "删除成功"    //string，返回信息
}
```

#### 查看当前证券账户

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/account/get

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI"    //string，token
}
```

- **Response** 

```json
{
    "data":
    {
        "records": [    
            {
                "id": 1,    //int，股票账户主键
                "profit": 800    //double，股票账户总收益
            },
            {
                "id": 2,    //int，股票账户主键
                "profit": -200    //double，股票账户总收益
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "添加成功"    //string，返回信息
}
```

### 股票持仓管理

#### 查看该证券账户下持仓信息

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/stock/get

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "accountID": 2,    //int，所属证券账户外键
}
```

- **Response** 

```json
{
    "data":
    {
        "stocks": [
            {
                "id": 1,    //int，股票主键
                "name": "股票名称",    //string，股票名称
                "code": "sz030303",    //string，股票代码
                "positionNum": 300,    //int，持仓股数
                "price": 4.51,    //double，当前每股价格
                "profit": 800    //double，该股票总收益
            },
            {
                "id": 2,    //int，股票主键
                "name": "股票名称",    //string，股票名称
                "code": "sz515151",    //string，股票代码
                "positionNum": 7000,    //int，持仓股数
                "price": 50.20,    //double，当前每股价格
                "profit": -2000    //double，该股票总收益
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "获取成功"    //string，返回信息
}
```

### 股票交易管理

#### 添加股票交易记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/operation/add

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "accountID": 2,    //int，证券账户外键
    "code": "sz030303",    //string，股票代码
    "buyNum": 20,    //int，买入股数
    "saleNum": 0,    //int，卖出股数
    "sharePrice": 5.01    //double，交易时价格
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "记录成功"    //string，返回信息
}
```

#### 删除股票交易记录

- **HTTP Method** 

​     [DELETE]

- **PATH** 

​     /api/security/operation/delete

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "accountID": 2,    //int，证券账户外键
    "code": "sz030303",    //string，股票代码
    "id": 1    //int，股票交易记录主键
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "删除成功"    //string，返回信息
}
```

#### 修改股票交易记录

- **HTTP Method** 

​     [PUT]

- **PATH** 

​     /api/security/operation/update

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "id": 1,    //int，交易记录表主键
    "accountID": 2,    //int，证券账户外键
    "code": "sz003030",    //string，股票代码
    "buyNum": 20,    //int，买入股数
    "saleNum": 0,    //int，卖出股数
    "sharePrice": 5.01    //double，交易时价格
}
```

- **Response** 

```json
{
    "code": 0,    //int，状态码
    "msg": "修改成功"    //string，返回信息
}
```

#### 获取所有股票交易记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/operation/get_all

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "accountID": 2    //int，证券账户主键
}
```

- **Response** 

```json
{
    "data":
    {
        "records":[
            {
                "id": 1,    //交易表主键
                "name": "股票名称",    //string，股票名称
                "code": "sz030303",    //string，股票代码
                "buyNum": 300,    //int，买入股数
                "saleNum": 0,    //int，卖出股数
                "sharePrice": 50.01,    //double，交易是价格
                "time": "2021-07-01 12:00:00"    //string，交易时间
            },
            {
                "id": 2,    //交易表主键
                "name": "股票名称",    //string，股票名称
                "code": "sz686868",    //string，股票代码
                "buyNum": 0,    //int，买入股数
                "saleNum": 100,    //int，卖出股数
                "sharePrice": 52.05,    //double，交易是价格
                "time": "2021-07-02 12:00:00"    //string，交易时间
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```

#### 根据时间段获取股票交易记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/operation/get_by_time

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "accountID": 2,    //int，证券账户主键
    "beginTime": "2021-07-01 11:50:50",    //string，起始时间
    "endTime": "2021-07-01 13:00:00"    //string，结束时间
}
```

- **Response** 

```json
{
    "data":
    {
        "records":[
            {
                "id": 1,    //交易表主键
                "name": "股票名称",    //string，股票名称
                "code": "sz030303",    //string，股票代码
                "buyNum": 300,    //int，买入股数
                "saleNum": 0,    //int，卖出股数
                "sharePrice": 50.01,    //double，交易是价格
                "time": "2021-07-01 12:00:00"    //string，交易时间
            },
            {
                "id": 1,    //交易表主键
                "name": "股票名称",    //string，股票名称
                "code": "sz686868",    //string，股票代码
                "buyNum": 0,    //int，买入股数
                "saleNum": 100,    //int，卖出股数
                "sharePrice": 52.05,    //double，交易是价格
                "time": "2021-07-01 12:00:00"    //string，交易时间
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```

#### 根据股票获取股票交易记录

- **HTTP Method** 

​     [POST]

- **PATH** 

​     /api/security/operation/get_by_stock

- **Request** 

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAxMTIyMzIsIm9wZW5JRCI6IjEyMzM0NTM0NSJ9.U5bTxP6VJcIYKVolaYKobOm5oEn_-nydr01aHWz72cI",    //string，token
    "accountID": 2,    //int，证券账户主键
    "code": "sz003030"    //string，股票代码
}
```

- **Response** 

```json
{
    "data":
    {
        "records":[
            {
                "id": 1,    //交易表主键
                "name": "股票名称",    //string，股票名称
                "code": "sz030303",    //string，股票代码
                "buyNum": 300,    //int，买入股数
                "saleNum": 0,    //int，卖出股数
                "sharePrice": 50.01,    //double，交易是价格
                "time": "2021-07-01 12:00:00"    //string，交易时间
            },
            {
                "id": 2,    //交易表主键
                "name": "股票名称",    //string，股票名称
                "code": "sz686868",    //string，股票代码
                "buyNum": 0,    //int，买入股数
                "saleNum": 100,    //int，卖出股数
                "sharePrice": 52.05,    //double，交易是价格
                "time": "2021-07-01 12:00:00"    //string，交易时间
            }
        ]
    },
    "code": 0,    //int，状态码
    "msg": "查询成功"    //string，返回信息
}
```
