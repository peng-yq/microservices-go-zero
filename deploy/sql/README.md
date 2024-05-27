## sql脚本

### microservices_order

#### homestay_order

| 字段名                  | 类型                                          | 默认值                           | 注释                              |
|-----------------------|---------------------------------------------|---------------------------------|---------------------------------|
| `id`                  | bigint NOT NULL AUTO_INCREMENT              |                                 | 主键，自动增长                        |
| `create_time`         | datetime NOT NULL                           | CURRENT_TIMESTAMP               | 记录创建时间                        |
| `update_time`         | datetime NOT NULL                           | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time`         | datetime NOT NULL                           | CURRENT_TIMESTAMP               | 记录删除时间                        |
| `del_state`           | tinyint NOT NULL                            | '0'                             | 删除状态，0表示未删除                |
| `version`             | bigint NOT NULL                             | '0'                             | 版本号，用于乐观锁                     |
| `sn`                  | char(25)                                     | ''                              | 订单号，唯一                        |
| `user_id`             | bigint NOT NULL                             | '0'                             | 下单用户ID                          |
| `homestay_id`         | bigint NOT NULL                             | '0'                             | 民宿ID                              |
| `title`               | varchar(32)                                 | ''                              | 标题                               |
| `sub_title`           | varchar(128)                                | ''                              | 副标题                              |
| `cover`               | varchar(1024)                               | ''                              | 封面图片地址                          |
| `info`                | varchar(4069)                               | ''                              | 介绍文本                            |
| `people_num`          | tinyint(1) NOT NULL                         | '0'                             | 容纳人数                            |
| `row_type`            | tinyint(1) NOT NULL                         | '0'                             | 售卖类型，0按房间，1按人次             |
| `need_food`           | tinyint(1) NOT NULL                         | '0'                             | 是否需要餐食，0不需要，1需要          |
| `food_info`           | varchar(1024)                               | ''                              | 餐食信息                            |
| `food_price`          | bigint NOT NULL                             |                                 | 餐食价格（单位：分）                  |
| `homestay_price`      | bigint NOT NULL                             |                                 | 民宿价格（单位：分）                  |
| `market_homestay_price` | bigint NOT NULL                            | '0'                             | 民宿市场价格（单位：分）              |
| `homestay_business_id` | bigint NOT NULL                            | '0'                             | 店铺ID                              |
| `homestay_user_id`    | bigint NOT NULL                             | '0'                             | 店铺房东ID                          |
| `live_start_date`     | date NOT NULL                               |                                 | 开始入住日期                          |
| `live_end_date`       | date NOT NULL                               |                                 | 结束入住日期                          |
| `live_people_num`     | tinyint NOT NULL                            | '0'                             | 实际入住人数                          |
| `trade_state`         | tinyint(1) NOT NULL                         | '0'                             | 交易状态，-1已取消，0待支付等          |
| `trade_code`          | char(8)                                     | ''                              | 确认码                              |
| `remark`              | varchar(64)                                 | ''                              | 用户下单备注                          |
| `order_total_price`   | bigint NOT NULL                             | '0'                             | 订单总价格（单位：分）                |
| `food_total_price`    | bigint NOT NULL                             | '0'                             | 餐食总价格（单位：分）                |
| `homestay_total_price` | bigint NOT NULL                            | '0'                             | 民宿总价格（单位：分）                |

### microservices_payment

#### third_payment

| 字段名             | 类型                                              | 默认值                            | 注释                              |
|------------------|-------------------------------------------------|----------------------------------|---------------------------------|
| `id`             | bigint NOT NULL AUTO_INCREMENT                  |                                  | 主键，自动增长                        |
| `sn`             | char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 流水单号，唯一                      |
| `create_time`    | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time`    | datetime NOT NULL                               | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time`    | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`      | tinyint(1) NOT NULL                             | '0'                              | 删除状态，0表示未删除                |
| `version`        | bigint NOT NULL                                 | '0'                              | 乐观锁版本号                        |
| `user_id`        | bigint NOT NULL                                 | '0'                              | 用户id                             |
| `pay_mode`       | varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 支付方式，如"1:微信支付"             |
| `trade_type`     | varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 第三方支付类型                      |
| `trade_state`    | varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 第三方交易状态                      |
| `pay_total`      | bigint NOT NULL                                 | '0'                              | 支付总金额（单位：分）                |
| `transaction_id` | char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 第三方支付单号                      |
| `trade_state_desc` | varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 支付状态描述                        |
| `order_sn`       | char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 业务单号                            |
| `service_type`   | varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 业务类型                            |
| `pay_status`     | tinyint(1) NOT NULL                             | '0'                              | 平台内交易状态，如-1:支付失败，0:未支付，1:支付成功，2:已退款 |
| `pay_time`       | datetime NOT NULL                               | '1970-01-01 08:00:00'            | 支付成功时间                        |

### microservices_travel

以下是您提供的各个表结构的详细描述，转换成 Markdown 格式的表格：

#### homestay

| 字段名                | 类型                                               | 默认值                            | 注释                              |
|---------------------|--------------------------------------------------|----------------------------------|---------------------------------|
| `id`                | bigint NOT NULL AUTO_INCREMENT                   |                                  | 主键，自动增长                        |
| `create_time`       | datetime NOT NULL                                | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time`       | datetime NOT NULL                                | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time`       | datetime NOT NULL                                | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`         | tinyint NOT NULL                                 | '0'                              | 删除状态，0表示未删除                |
| `version`           | bigint NOT NULL                                  | '0'                              | 版本号                             |
| `title`             | varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 标题                               |
| `sub_title`         | varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 副标题                             |
| `banner`            | varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 轮播图，第一张封面                    |
| `info`              | varchar(4069) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 介绍                               |
| `people_num`        | tinyint(1) NOT NULL                              | '0'                              | 容纳人的数量                         |
| `homestay_business_id` | bigint NOT NULL                                | '0'                              | 民宿店铺id                          |
| `user_id`           | bigint NOT NULL                                  | '0'                              | 房东id，冗余字段                     |
| `row_state`         | tinyint(1) NOT NULL                              | '0'                              | 0:下架 1:上架                       |
| `row_type`          | tinyint(1) NOT NULL                              | '0'                              | 售卖类型0：按房间出售 1:按人次出售    |
| `food_info`         | varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL |                                  | 餐食标准                            |
| `food_price`        | bigint NOT NULL                                  | '0'                              | 餐食价格（分）                      |
| `homestay_price`    | bigint NOT NULL                                  | '0'                              | 民宿价格（分）                      |
| `market_homestay_price` | bigint NOT NULL                              | '0'                              | 民宿市场价格（分）                   |

#### homestay_activity

| 字段名              | 类型                                              | 默认值                            | 注释                              |
|-------------------|-------------------------------------------------|----------------------------------|---------------------------------|
| `id`              | bigint NOT NULL AUTO_INCREMENT                  |                                  | 主键，自动增长                        |
| `create_time`     | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time`     | datetime NOT NULL                               | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time`     | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`       | tinyint NOT NULL                                | '0'                              | 删除状态，0表示未删除                |
| `row_type`        | varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 活动类型                            |
| `data_id`         | bigint NOT NULL                                 | '0'                              | 业务表id（id跟随活动类型走）          |
| `row_status`      | tinyint(1) NOT NULL                             | '0'                              | 0:下架 1:上架                       |
| `version`         | bigint NOT NULL                                 | '0'                              | 版本号                             |

#### homestay_business
| 字段名             | 类型                                              | 默认值                            | 注释                              |
|------------------|-------------------------------------------------|----------------------------------|---------------------------------|
| `id`             | bigint NOT NULL AUTO_INCREMENT                  |                                  | 主键，自动增长                        |
| `create_time`    | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time`    | datetime NOT NULL                               | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time`    | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`      | tinyint NOT NULL                                | '0'                              | 删除状态，0表示未删除                |
| `title`          | varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 店铺名称                            |
| `user_id`        | bigint NOT NULL                                 | '0'                              | 关联的用户id                        |
| `info`           | varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 店铺介绍                            |
| `boss_info`      | varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 房东介绍                            |
| `license_fron`   | varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 营业执照正面                        |
| `license_back`   | varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 营业执照背面                        |
| `row_state`      | tinyint(1) NOT NULL                             | '0'                              | 0:禁止营业 1:正常营业               |
| `star`           | double(2,1) NOT NULL                            | '0.0'                            | 店铺整体评价，冗余                   |
| `tags`           | varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 店家标签                            |
| `cover`          | varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 封面图                              |
| `header_img`     | varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 店招门头图片                        |

#### homestay_comment

| 字段名            | 类型                                              | 默认值                            | 注释                              |
|-----------------|-------------------------------------------------|----------------------------------|---------------------------------|
| `id`            | bigint NOT NULL AUTO_INCREMENT                  |                                  | 主键，自动增长                        |
| `create_time`   | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time`   | datetime NOT NULL                               | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time`   | datetime NOT NULL                               | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`     | tinyint NOT NULL                                | '0'                              | 删除状态，0表示未删除                |
| `homestay_id`   | bigint NOT NULL                                 | '0'                              | 民宿id                             |
| `user_id`       | bigint NOT NULL                                 | '0'                              | 用户id                             |
| `content`       | varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 评论内容                            |
| `star`          | json NOT NULL                                   |                                  | 星星数,多个维度                     |

### microservices_usercenter

#### user

| 字段名          | 类型                                                     | 默认值                            | 注释                              |
|---------------|--------------------------------------------------------|----------------------------------|---------------------------------|
| `id`          | bigint NOT NULL AUTO_INCREMENT                         |                                  | 主键，自动增长                        |
| `create_time` | datetime NOT NULL                                      | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time` | datetime NOT NULL                                      | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time` | datetime NOT NULL                                      | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`   | tinyint NOT NULL                                       | '0'                              | 删除状态，0表示未删除                |
| `version`     | bigint NOT NULL                                        | '0'                              | 版本号                             |
| `mobile`      | char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 手机号                             |
| `password`    | varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 密码                               |
| `nickname`    | varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 昵称                               |
| `sex`         | tinyint(1) NOT NULL                                    | '0'                              | 性别 0:男 1:女                     |
| `avatar`      | varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 头像                               |
| `info`        | varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 用户信息                           |

#### user_auth

| 字段名          | 类型                                                     | 默认值                            | 注释                              |
|---------------|--------------------------------------------------------|----------------------------------|---------------------------------|
| `id`          | bigint NOT NULL AUTO_INCREMENT                         |                                  | 主键，自动增长                        |
| `create_time` | datetime NOT NULL                                      | CURRENT_TIMESTAMP                | 记录创建时间                        |
| `update_time` | datetime NOT NULL                                      | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 记录更新时间，每次更新时自动设置当前时间 |
| `delete_time` | datetime NOT NULL                                      | CURRENT_TIMESTAMP                | 记录删除时间                        |
| `del_state`   | tinyint NOT NULL                                       | '0'                              | 删除状态，0表示未删除                |
| `version`     | bigint NOT NULL                                        | '0'                              | 版本号                             |
| `user_id`     | bigint NOT NULL                                        | '0'                              | 用户ID                             |
| `auth_key`    | varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 平台唯一id                          |
| `auth_type`   | varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL | ''                               | 平台类型                            |


