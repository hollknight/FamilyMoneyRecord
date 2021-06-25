# FamilyMoneyRecord

[toc]
## 数据库设计
### 用户表（users）
|       id        |    username    |      password       |    name     | receipt_sum | disbursement_sum | advance_consumption |
| :-------------: | :------------: | :-----------------: | :---------: | :---------: | :--------------: | :-----------------: |
| bigint unsigned |  varchar(100)  |      char(32)       | varchar(20) |     int     |       int        |         int         |
|    自增主键     | 用户名（账号） | 密码（MD5加密存储） |    姓名     |    收入     |       支出       |     预消费金额      |

### 账单表（bills）
|       id        |     user_id     | receipt | disbursement |    type     |   time   |
| :-------------: | :-------------: | :-----: | :----------: | :---------: | :------: |
| bigint unsigned | bigint unsigned |   int   |     int      | varchar(25) | datetime |
|    自增主键     |   用户表外键    |  收入   |     支出     |  收支类型   | 操作时间 |

### 证券账户表（accounts）

|       id        |     user_id     |    profit    |
| :-------------: | :-------------: | :----------: |
| bigint unsigned | bigint unsigned |     int      |
|    自增主键     |   用户表外键    | 当前交易盈亏 |

### 股票操作表（operations）
|       id        |   account_id    |    code     |    share_price     | buy_num  | sale_num |   time   |
| :-------------: | :-------------: | :---------: | :----------------: | :------: | :------: | :------: |
| bigint unsigned | bigint unsigned | varchar(10) |        int         |   int    |   int    | datetime |
|    自增主键     | 证券账户表外键  |  股票代码   | 交易时每股股票价格 | 买入股数 | 卖出股数 | 操作时间 |

### 股票持仓表（stocks）
|       id        |   account_id    |    code     | position_num |    profit    |
| :-------------: | :-------------: | :---------: | :----------: | :----------: |
| bigint unsigned | bigint unsigned | varchar(10) |     int      |     int      |
|    自增主键     | 证券账户表外键  |  股票代码   |   持仓股数   | 当前交易盈亏 |
