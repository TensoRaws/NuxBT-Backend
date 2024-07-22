-- 创建用户表
CREATE TABLE `user` (
  `user_id` INT NOT NULL AUTO_INCREMENT, -- 使用 AUTO_INCREMENT 作为自增主键
  `username` VARCHAR(255) NOT NULL UNIQUE, -- 用户名
  `email` VARCHAR(255) NOT NULL UNIQUE, -- 用户邮箱
  `password` VARCHAR(255) NOT NULL, -- 用户密码
  `private` BOOLEAN NOT NULL DEFAULT false, -- 是否私密，默认为 false
  `experience` INT DEFAULT 0, -- 用户经验值，默认为 0
  `inviter` INT NOT NULL DEFAULT 0, -- 邀请人ID
  `avatar` VARCHAR(255), -- 用户头像链接
  `signature` TEXT, -- 用户签名
  `background` VARCHAR(255), -- 用户背景图片链接
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 账户创建时间
  `deleted_at` DATETIME, -- 账户删除时间
  `last_active` DATETIME, -- 最后活跃时间
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

-- 创建种子表
CREATE TABLE `torrent` (
  `torrent_id` INT NOT NULL AUTO_INCREMENT,
  `hash` VARCHAR(255) NOT NULL UNIQUE,
  `uploader_id` INT NOT NULL,
  `official` BOOLEAN NOT NULL DEFAULT FALSE,
  `size` BIGINT NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `subtitle` VARCHAR(255) NOT NULL,
  `essay` TEXT,
  `description` TEXT NOT NULL,
  `genre` VARCHAR(255) NOT NULL,
  `anidb_id` INT NOT NULL,
  `img` VARCHAR(255) NOT NULL,
  `resolution` VARCHAR(255) NOT NULL,
  `video_codec` VARCHAR(255) NOT NULL,
  `audio_codec` VARCHAR(255) NOT NULL,
  `language` VARCHAR(255) NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  `file_list` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME, -- 软删除时间戳
  PRIMARY KEY (`torrent_id`)
);
