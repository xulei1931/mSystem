create schema if not exists dbname;
create table if not exists dbname.user (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '类型id',
  `user_name` char(50) NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '用户的密码',
  `create_at` char(50) NOT NULL DEFAULT '' COMMENT '用户的注册时间',
  `email` char(50) NOT NULL DEFAULT '' COMMENT '用户的email',
  `phone` varchar(12) NOT NULL DEFAULT '0' COMMENT '用户联系方式',
  PRIMARY KEY (`user_id`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4