package query

import (
	"encoding/json"
	"time"
)

// TUserDto 用户表 数据传输对象
type TUserDto struct {
	Id              uint64    `json:"id"`              // 主键ID
	Username        string    `json:"username"`        // 用户名
	Email           string    `json:"email"`           // 邮箱
	Password        string    `json:"password"`        // 密码（哈希）
	Role            string    `json:"role"`            // 角色: user | admin
	CreateTime      time.Time `json:"createTime"`      // 创建时间
	UpdateTime      time.Time `json:"updateTime"`      // 更新时间
	IdList          []uint64  `json:"idList"`          // 主键ID IN 查询
	UsernameFuzzy   string    `json:"usernameFuzzy"`   // 用户名 模糊查询
	UsernameList    []string  `json:"usernameList"`    // 用户名 IN 查询
	EmailFuzzy      string    `json:"emailFuzzy"`      // 邮箱 模糊查询
	EmailList       []string  `json:"emailList"`       // 邮箱 IN 查询
	PasswordFuzzy   string    `json:"passwordFuzzy"`   // 密码（哈希） 模糊查询
	RoleFuzzy       string    `json:"roleFuzzy"`       // 角色: user | admin 模糊查询
	CreateTimeStart time.Time `json:"createTimeStart"` // 创建时间 开始时间
	CreateTimeEnd   time.Time `json:"createTimeEnd"`   // 创建时间 结束时间
	UpdateTimeStart time.Time `json:"updateTimeStart"` // 更新时间 开始时间
	UpdateTimeEnd   time.Time `json:"updateTimeEnd"`   // 更新时间 结束时间
	OrderBy         string    `json:"orderBy"`         // 排序字段
	PageOffset      int       `json:"pageOffset"`      // 分页偏移量
	PageSize        int       `json:"pageSize"`        // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *TUserDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *TUserDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// TUserDtoBuilder 用于构建 TUserDto 实例的 Builder
type TUserDtoBuilder struct {
	instance *TUserDto
}

// NewTUserDtoBuilder 创建一个新的 TUserDtoBuilder 实例
// 返回:
//   - *TUserDtoBuilder: Builder 实例，用于链式调用
func NewTUserDtoBuilder() *TUserDtoBuilder {
	return &TUserDtoBuilder{
		instance: &TUserDto{},
	}
}

// WithUsername 设置 username 字段
// 参数:
//   - username: 用户名
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithUsername(username string) *TUserDtoBuilder {
	b.instance.Username = username
	return b
}

// WithEmail 设置 email 字段
// 参数:
//   - email: 邮箱
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithEmail(email string) *TUserDtoBuilder {
	b.instance.Email = email
	return b
}

// WithPassword 设置 password 字段
// 参数:
//   - password: 密码（哈希）
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithPassword(password string) *TUserDtoBuilder {
	b.instance.Password = password
	return b
}

// WithRole 设置 role 字段
// 参数:
//   - role: 角色: user | admin
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithRole(role string) *TUserDtoBuilder {
	b.instance.Role = role
	return b
}

// WithCreateTime 设置 createTime 字段
// 参数:
//   - createTime: 创建时间
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithCreateTime(createTime time.Time) *TUserDtoBuilder {
	b.instance.CreateTime = createTime
	return b
}

// WithUpdateTime 设置 updateTime 字段
// 参数:
//   - updateTime: 更新时间
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithUpdateTime(updateTime time.Time) *TUserDtoBuilder {
	b.instance.UpdateTime = updateTime
	return b
}

// WithUsernameFuzzy 设置 username_fuzzy 字段
// 参数:
//   - usernameFuzzy: 用户名 模糊查询
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithUsernameFuzzy(usernameFuzzy string) *TUserDtoBuilder {
	b.instance.UsernameFuzzy = usernameFuzzy
	return b
}

// WithUsernameList 设置 usernameList 字段
// 参数:
//   - usernameList: 用户名 IN 查询
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithUsernameList(usernameList []string) *TUserDtoBuilder {
	b.instance.UsernameList = usernameList
	return b
}

// WithEmailFuzzy 设置 email_fuzzy 字段
// 参数:
//   - emailFuzzy: 邮箱 模糊查询
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithEmailFuzzy(emailFuzzy string) *TUserDtoBuilder {
	b.instance.EmailFuzzy = emailFuzzy
	return b
}

// WithEmailList 设置 emailList 字段
// 参数:
//   - emailList: 邮箱 IN 查询
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithEmailList(emailList []string) *TUserDtoBuilder {
	b.instance.EmailList = emailList
	return b
}

// WithPasswordFuzzy 设置 password_fuzzy 字段
// 参数:
//   - passwordFuzzy: 密码（哈希） 模糊查询
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithPasswordFuzzy(passwordFuzzy string) *TUserDtoBuilder {
	b.instance.PasswordFuzzy = passwordFuzzy
	return b
}

// WithRoleFuzzy 设置 role_fuzzy 字段
// 参数:
//   - roleFuzzy: 角色: user | admin 模糊查询
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithRoleFuzzy(roleFuzzy string) *TUserDtoBuilder {
	b.instance.RoleFuzzy = roleFuzzy
	return b
}

// WithCreateTimeStart 设置 createTimeStart 字段
// 参数:
//   - createTimeStart: 创建时间 开始时间
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithCreateTimeStart(createTimeStart time.Time) *TUserDtoBuilder {
	b.instance.CreateTimeStart = createTimeStart
	return b
}

// WithCreateTimeEnd 设置 createTimeEnd 字段
// 参数:
//   - createTimeEnd: 创建时间 结束时间
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithCreateTimeEnd(createTimeEnd time.Time) *TUserDtoBuilder {
	b.instance.CreateTimeEnd = createTimeEnd
	return b
}

// WithUpdateTimeStart 设置 updateTimeStart 字段
// 参数:
//   - updateTimeStart: 更新时间 开始时间
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithUpdateTimeStart(updateTimeStart time.Time) *TUserDtoBuilder {
	b.instance.UpdateTimeStart = updateTimeStart
	return b
}

// WithUpdateTimeEnd 设置 updateTimeEnd 字段
// 参数:
//   - updateTimeEnd: 更新时间 结束时间
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithUpdateTimeEnd(updateTimeEnd time.Time) *TUserDtoBuilder {
	b.instance.UpdateTimeEnd = updateTimeEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithOrderBy(orderBy string) *TUserDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithPageOffset(pageOffset int) *TUserDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *TUserDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *TUserDtoBuilder) WithPageSize(pageSize int) *TUserDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 TUserDto 实例
// 返回:
//   - *TUserDto: 构建完成的实例
func (b *TUserDtoBuilder) Build() *TUserDto {
	return b.instance
}
