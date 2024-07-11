-- 创建用户表
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` INT NOT NULL AUTO_INCREMENT, -- 使用 AUTO_INCREMENT 作为自增主键
  `username` VARCHAR(255) NOT NULL UNIQUE, -- 用户名
  `email` VARCHAR(255) NOT NULL UNIQUE, -- 用户邮箱
  `password` VARCHAR(255) NOT NULL, -- 用户密码
  `private` BOOLEAN NOT NULL DEFAULT false, -- 是否私密，默认为 false
  `experience` INT DEFAULT 0, -- 用户经验值，默认为 0
  `inviter` INT NOT NULL DEFAULT 0, -- 邀请人ID
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 账户创建时间
  `last_active` DATETIME, -- 最后活跃时间
  `avatar` VARCHAR(255), -- 用户头像链接
  `signature` TEXT, -- 用户签名
  `background` VARCHAR(255), -- 用户背景图片链接
  `deleted_at` DATETIME, -- 账户删除时间
  PRIMARY KEY (`user_id`)
);

-- 创建用户角色表
CREATE TABLE `user_role` (
  `role_id` INT NOT NULL AUTO_INCREMENT,  -- 角色ID，作为自增主键
  `user_id` INT NOT NULL,    -- 用户ID
  `role` VARCHAR(255) NOT NULL,  -- 角色名称
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
  `deleted_at` DATETIME, -- 软删除时间戳
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `uk_user_role` (`user_id`, `role`)
);
