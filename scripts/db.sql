-- 创建用户表
DROP TABLE IF EXISTS "user" CASCADE;
CREATE TABLE "user" (
  "user_id" SERIAL PRIMARY KEY, -- 使用 SERIAL 作为自增主键
  "username" VARCHAR(255) NOT NULL UNIQUE, -- 用户名
  "email" VARCHAR(255) NOT NULL UNIQUE, -- 用户邮箱
  "password" VARCHAR(255) NOT NULL, -- 用户密码
  "private" BOOL NOT NULL DEFAULT false, -- 是否私密，默认为 false
  "experience" INTEGER DEFAULT 0, -- 用户经验值，默认为 0
  "inviter" INTEGER REFERENCES "user" ("user_id") ON DELETE SET NULL, -- 邀请人ID
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 账户创建时间
  "last_active" TIMESTAMP WITH TIME ZONE, -- 最后活跃时间
  "avatar" VARCHAR(255), -- 用户头像链接
  "signature" TEXT, -- 用户签名
  "background" VARCHAR(255), -- 用户背景图片链接
  "deleted_at" TIMESTAMP WITH TIME ZONE -- 账户删除时间
);

-- 创建用户角色表
CREATE TABLE "user_role" (
  "role_id" SERIAL PRIMARY KEY,  -- 角色ID，作为自增主键
  "user_id" INTEGER NOT NULL,    -- 用户ID
  "role" VARCHAR(255) NOT NULL,  -- 角色名称
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
  "deleted_at" TIMESTAMP WITH TIME ZONE, -- 软删除时间戳
  -- 唯一性约束，确保每个用户和角色的组合唯一
  UNIQUE ("user_id", "role")
);