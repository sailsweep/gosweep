package schema

import "entgo.io/ent"

/**
CREATE TABLE `ms_application` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '应用ID',
  `app_name` varchar(32) NOT NULL DEFAULT '' COMMENT '应用名称',
  `app_id` varchar(64) NOT NULL DEFAULT '' COMMENT '应用编码',
  `app_secret` varchar(255) NOT NULL DEFAULT '' COMMENT '应用密钥',
  `notify_url` varchar(255) NOT NULL DEFAULT '' COMMENT '回调地址',
  `auth` varchar(255) NOT NULL DEFAULT '' COMMENT '客户端授权',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已经注册,注册绑定之后不允许修改code',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用表';
*/

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
	return nil
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return nil
}
