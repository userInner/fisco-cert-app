### Database document
> 2024-04-17 22:42:57
#### t_auth  
NO | KEY | COLUMN | COMMENT | DATA_TYPE | NOTNULL | REMARK
:---: | :---: | --- |--------| --- | :---: | ---
1|PRI|id| 主键     |BIGINT(20) UNSIGNED|Y|
2| |auth_id| 存证ID   |BIGINT(20) UNSIGNED|Y|
3| |user_id| 存证用户ID |BIGINT(20) UNSIGNED|Y|
4| |image_id| 存储图片ID |BIGINT(20) UNSIGNED|Y|
5| |address| 存证人地址  |VARCHAR(256)|Y|
6| |created_at| 创建时间   |DATETIME|N|
7| |updated_at| 更新时间   |DATETIME|N|
8| |deleted_at| 软删除键   |DATETIME|N|
#### t_cert  
NO | KEY | COLUMN | COMMENT | DATA_TYPE | NOTNULL | REMARK
:---: | :---: | --- |---------| --- | :---: | ---
1|PRI|id| 主键      |BIGINT(20) UNSIGNED|Y|
2| |tx_id| 交易ID    |VARCHAR(256)|Y|
3| |user_id| 存证用户ID  |BIGINT(20) UNSIGNED|Y|
4| |image_id| 存储图片ID  |BIGINT(20) UNSIGNED|Y|
5| |evidence| 存证信息    |VARCHAR(256)|Y|
6| |created_at|  创建时间     |DATETIME|N|
7| |updated_at|  更新时间     |DATETIME|N|
8| |deleted_at|  软删除键     |DATETIME|N|
#### t_img  
NO | KEY | COLUMN | COMMENT | DATA_TYPE | NOTNULL | REMARK
:---: | :---: | --- |---------| --- | :---: | ---
1|PRI|id| 主键      |BIGINT(20) UNSIGNED|Y|
2| |img_id| 图像ID    |BIGINT(20) UNSIGNED|Y|
3| |user_id| 存储用户ID  |BIGINT(20) UNSIGNED|Y|
4| |file_name| 图片名     |VARCHAR(256)|Y|
5| |file_size| 图片大小    |BIGINT(20) UNSIGNED|Y|
6| |suffix| 图片后缀    |VARCHAR(256)|Y|
7| |created_at|  创建时间     |DATETIME|N|
8| |updated_at|  更新时间     |DATETIME|N|
9| |deleted_at|  软删除键     |DATETIME|N|
#### t_transaction_monitor  
NO | KEY | COLUMN | COMMENT  | DATA_TYPE | NOTNULL | REMARK
:---: | :---: | --- |----------| --- | :---: | ---
1|PRI|id| 主键       |BIGINT(20) UNSIGNED|Y|
2| |user_id| 用户ID     |BIGINT(20) UNSIGNED|Y|
3| |trans_count| 交易量      |INT(11)|Y|
4| |trans_hash_lastest| 最新交易hash |VARCHAR(128)|N|
5| |created_at|  创建时间      |DATETIME|N|
6| |updated_at|  更新时间      |DATETIME|N|
7| |deleted_at|  软删除键      |DATETIME|N|
#### t_user  
NO | KEY | COLUMN | COMMENT | DATA_TYPE | NOTNULL | REMARK
:---: | :---: | --- |------| --- | :---: | ---
1|PRI|id| 主键   |BIGINT(20) UNSIGNED|Y|
2| |user_id| 用户ID |BIGINT(20) UNSIGNED|Y|
3| |user_name| 用户名  |VARCHAR(256)|Y|
4| |password| 用户密码 |VARCHAR(256)|Y|
5| |email| 用户邮箱 |VARCHAR(256)|N|
6| |gender| 用户性别 |TINYINT(3) UNSIGNED|Y|
7| |address| 用户私钥转换的地址 |VARCHAR(256)|Y|
8| |status| 用户状态 |TINYINT(3) UNSIGNED|Y|
9| |created_at| 创建时间   |DATETIME|N|
10| |updated_at| 更新时间     |DATETIME|N|
11| |deleted_at| 软删除键     |DATETIME|N|
