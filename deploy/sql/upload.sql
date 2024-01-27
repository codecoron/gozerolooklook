create table upload_file (
                                  id int not null primary key auto_increment,
                                  user_id int not null comment '上传用户id',
                                  file_name varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci comment '文件名',
                                  ext varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci comment '扩展名',
                                  size int comment '文件大小',
                                  url varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci not null comment '下载链接',
                                  create_time datetime not null default current_timestamp,
                                  update_time datetime not null default current_timestamp ON UPDATE CURRENT_TIMESTAMP,
                                  delete_time datetime
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文件信息表';
