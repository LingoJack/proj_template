package entity

import (
	"encoding/json"
	"time"
)

// TUser 用户表
type TUser struct {
	Id         uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:主键ID;not null" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(64);comment:用户名;not null" json:"username"`
	Email      string    `gorm:"column:email;type:varchar(128);comment:邮箱;not null" json:"email"`
	Password   string    `gorm:"column:password;type:varchar(256);comment:密码（哈希）;not null" json:"password"`
	Role       string    `gorm:"column:role;type:varchar(32);default:user;comment:角色: user | admin;not null" json:"role"`
	CreateTime time.Time `gorm:"column:createTime;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间;not null" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime;type:datetime;default:CURRENT_TIMESTAMP;comment:更新时间;not null" json:"updateTime"`
}

// TableName 返回表名
func (t *TUser) TableName() string {
	return "t_user"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *TUser) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *TUser) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// TUserBuilder 用于构建 TUser 实例的 Builder
type TUserBuilder struct {
	instance *TUser
}

// NewTUserBuilder 创建一个新的 TUserBuilder 实例
// 返回:
//   - *TUserBuilder: Builder 实例，用于链式调用
func NewTUserBuilder() *TUserBuilder {
	return &TUserBuilder{
		instance: &TUser{},
	}
}

// WithUsername 设置 username 字段
// 参数:
//   - username: 用户名
//
// 返回:
//   - *TUserBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserBuilder) WithUsername(username string) *TUserBuilder {
	b.instance.Username = username
	return b
}

// WithEmail 设置 email 字段
// 参数:
//   - email: 邮箱
//
// 返回:
//   - *TUserBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserBuilder) WithEmail(email string) *TUserBuilder {
	b.instance.Email = email
	return b
}

// WithPassword 设置 password 字段
// 参数:
//   - password: 密码（哈希）
//
// 返回:
//   - *TUserBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserBuilder) WithPassword(password string) *TUserBuilder {
	b.instance.Password = password
	return b
}

// WithRole 设置 role 字段
// 参数:
//   - role: 角色: user | admin
//
// 返回:
//   - *TUserBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserBuilder) WithRole(role string) *TUserBuilder {
	b.instance.Role = role
	return b
}

// WithCreateTime 设置 createTime 字段
// 参数:
//   - createTime: 创建时间
//
// 返回:
//   - *TUserBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserBuilder) WithCreateTime(createTime time.Time) *TUserBuilder {
	b.instance.CreateTime = createTime
	return b
}

// WithUpdateTime 设置 updateTime 字段
// 参数:
//   - updateTime: 更新时间
//
// 返回:
//   - *TUserBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserBuilder) WithUpdateTime(updateTime time.Time) *TUserBuilder {
	b.instance.UpdateTime = updateTime
	return b
}

// Build 构建并返回 TUser 实例
// 返回:
//   - *TUser: 构建完成的实例
func (b *TUserBuilder) Build() *TUser {
	return b.instance
}
