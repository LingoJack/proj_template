package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/lingojack/proj_template/model/entity"
	"github.com/lingojack/proj_template/model/query"
	"gorm.io/gorm"
)

// TUserDao 用户表的Dao实现
type TUserDao struct {
	*gorm.DB
}

func (dao *TUserDao) Database() string {
	// TODO 补全 db 名称
	return "@database_name"
}

// NewTUserDao 创建TUserDao实例
// 参数:
//   - db: GORM数据库连接实例
//
// 返回:
//   - *TUserDao: Dao实例
func NewTUserDao(db *gorm.DB) *TUserDao {
	return &TUserDao{DB: db}
}

// ==================== 事务支持方法 ====================

// WithTx 使用指定的事务对象创建新的 DAO 实例
// 参数:
//   - tx: GORM事务对象
//
// 返回:
//   - *TUserDao: 使用事务的新 DAO 实例
//
// 使用示例:
//
//	db.Transaction(func(tx *gorm.DB) error {
//	    txDao := dao.WithTx(tx)
//	    return txDao.Insert(ctx, poBean)
//	})
func (dao *TUserDao) WithTx(tx *gorm.DB) *TUserDao {
	return &TUserDao{DB: tx}
}

// Transaction 在事务中执行操作
// 参数:
//   - ctx: 上下文对象
//   - fn: 事务处理函数，接收使用事务的 DAO 实例
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 自动管理事务的开始、提交和回滚
//   - 如果 fn 返回 error，事务会自动回滚
//   - 如果 fn 执行成功，事务会自动提交
//
// 使用示例:
//
//	err := dao.Transaction(ctx, func(txDao *TUserDao) error {
//	    if err := txDao.Insert(ctx, poBean1); err != nil {
//	        return err
//	    }
//	    if err := txDao.Insert(ctx, poBean2); err != nil {
//	        return err
//	    }
//	    return nil
//	})
func (dao *TUserDao) Transaction(ctx context.Context, fn func(*TUserDao) error) error {
	return dao.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txDao := &TUserDao{DB: tx}
		return fn(txDao)
	})
}

// ==================== 查询条件构建 ====================

// buildTUserQueryCondition 构建查询条件
// 参数:
//   - db: GORM数据库连接实例
//   - queryDto: 查询条件Dto对象
//
// 返回:
//   - *gorm.DB: 应用了查询条件的数据库连接
//   - error: 错误信息
//
// 说明:
//   - 支持精确匹配、模糊查询、IN查询、范围查询等多种查询方式
//   - 零值字段会被忽略，不会作为查询条件
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *TUserDao) buildTUserQueryCondition(db *gorm.DB, queryDto *query.TUserDto) (*gorm.DB, error) {
	if queryDto == nil {
		return db, nil
	}

	// 基础字段精确查询
	if queryDto.Id != 0 {
		db = db.Where("id = ?", queryDto.Id)
	}
	if queryDto.Username != "" {
		db = db.Where("username = ?", queryDto.Username)
	}
	if queryDto.Email != "" {
		db = db.Where("email = ?", queryDto.Email)
	}
	if queryDto.Password != "" {
		db = db.Where("password = ?", queryDto.Password)
	}
	if queryDto.Role != "" {
		db = db.Where("role = ?", queryDto.Role)
	}
	if !queryDto.CreateTime.IsZero() {
		db = db.Where("createTime = ?", queryDto.CreateTime)
	}
	if !queryDto.UpdateTime.IsZero() {
		db = db.Where("updateTime = ?", queryDto.UpdateTime)
	}

	// 模糊查询条件
	if queryDto.UsernameFuzzy != "" {
		db = db.Where("username LIKE ?", "%"+queryDto.UsernameFuzzy+"%")
	}
	if queryDto.EmailFuzzy != "" {
		db = db.Where("email LIKE ?", "%"+queryDto.EmailFuzzy+"%")
	}
	if queryDto.PasswordFuzzy != "" {
		db = db.Where("password LIKE ?", "%"+queryDto.PasswordFuzzy+"%")
	}
	if queryDto.RoleFuzzy != "" {
		db = db.Where("role LIKE ?", "%"+queryDto.RoleFuzzy+"%")
	}

	// 日期范围查询
	if !queryDto.CreateTimeStart.IsZero() {
		db = db.Where("createTime >= ?", queryDto.CreateTimeStart)
	}
	if !queryDto.CreateTimeEnd.IsZero() {
		db = db.Where("createTime < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.CreateTimeEnd)
	}
	if !queryDto.UpdateTimeStart.IsZero() {
		db = db.Where("updateTime >= ?", queryDto.UpdateTimeStart)
	}
	if !queryDto.UpdateTimeEnd.IsZero() {
		db = db.Where("updateTime < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.UpdateTimeEnd)
	}

	// IN 查询条件
	// 校验 IdList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.IdList != nil {
		if len(queryDto.IdList) == 0 {
			return nil, fmt.Errorf("IdList 不能为空列表")
		}
		db = db.Where("id IN ?", queryDto.IdList)
	}
	// 校验 UsernameList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.UsernameList != nil {
		if len(queryDto.UsernameList) == 0 {
			return nil, fmt.Errorf("UsernameList 不能为空列表")
		}
		db = db.Where("username IN ?", queryDto.UsernameList)
	}
	// 校验 EmailList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.EmailList != nil {
		if len(queryDto.EmailList) == 0 {
			return nil, fmt.Errorf("EmailList 不能为空列表")
		}
		db = db.Where("email IN ?", queryDto.EmailList)
	}

	return db, nil
}

// ==================== 基础查询方法 ====================

// SelectList 查询列表
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象，支持分页、排序、多条件查询
//
// 返回:
//   - [] *entity.TUser: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *TUserDao) SelectList(ctx context.Context, queryDto *query.TUserDto) ([]*entity.TUser, error) {
	var resultList []*entity.TUser
	db := dao.WithContext(ctx).Model(&entity.TUser{})

	// 应用查询条件
	var err error
	db, err = dao.buildTUserQueryCondition(db, queryDto)
	if err != nil {
		return nil, err
	}

	// 排序
	if queryDto != nil && queryDto.OrderBy != "" {
		if dao.isValidOrderBy(queryDto.OrderBy) {
			db = db.Order(queryDto.OrderBy)
		}
	}

	// 分页
	if queryDto != nil && queryDto.PageSize > 0 {
		db = db.Offset(queryDto.PageOffset * queryDto.PageSize).Limit(queryDto.PageSize)
	}

	err = db.Find(&resultList).Error
	return resultList, err
}

// SelectCount 查询数量
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象
//
// 返回:
//   - int64: 符合条件的记录数量
//   - error: 错误信息
//
// 说明:
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *TUserDao) SelectCount(ctx context.Context, queryDto *query.TUserDto) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&entity.TUser{})

	// 应用查询条件
	var err error
	db, err = dao.buildTUserQueryCondition(db, queryDto)
	if err != nil {
		return 0, err
	}

	err = db.Count(&count).Error
	return count, err
}

// SelectListWithAppendConditionFunc 查询列表（支持自定义条件函数）
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象，支持分页、排序、多条件查询
//   - appendConditionFunc: 自定义条件函数，用于添加额外的查询条件
//   - 如果为 nil，则不添加额外条件
//   - 函数签名: func(ctx context.Context, db *gorm.DB) *gorm.DB
//
// 返回:
//   - [] *entity.TUser: 查询结果列表
//   - error: 错误信息
//
// 使用示例:
//
//	// 示例1: 添加复杂的自定义条件
//	resultList, err := dao.SelectListWithAppendConditionFunc(ctx, queryDto, func(ctx context.Context, db *gorm.DB) *gorm.DB {
//	    return db.Where("status IN (?, ?)", "active", "pending").
//	              Where("created_at > ?", time.Now().AddDate(0, -1, 0))
//	})
//
//	// 示例2: 不添加额外条件
//	resultList, err := dao.SelectListWithAppendConditionFunc(ctx, queryDto, nil)
//
// 说明:
//   - 自定义条件函数在基础查询条件、排序和分页之后执行
//   - 适用于需要动态添加复杂查询条件的场景
//   - IN 查询条件校验规则同 SelectList 方法
func (dao *TUserDao) SelectListWithAppendConditionFunc(ctx context.Context, queryDto *query.TUserDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) ([]*entity.TUser, error) {
	var resultList []*entity.TUser
	db := dao.WithContext(ctx).Model(&entity.TUser{})

	// 应用查询条件
	var err error
	db, err = dao.buildTUserQueryCondition(db, queryDto)
	if err != nil {
		return nil, err
	}

	// 排序
	if queryDto != nil && queryDto.OrderBy != "" {
		if dao.isValidOrderBy(queryDto.OrderBy) {
			db = db.Order(queryDto.OrderBy)
		}
	}

	// 分页
	if queryDto != nil && queryDto.PageSize > 0 {
		db = db.Offset(queryDto.PageOffset * queryDto.PageSize).Limit(queryDto.PageSize)
	}

	// 应用自定义条件函数
	if appendConditionFunc != nil {
		db = appendConditionFunc(ctx, db)
	}

	err = db.Find(&resultList).Error
	return resultList, err
}

// SelectCountWithAppendConditionFunc 查询数量（支持自定义条件函数）
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象
//   - appendConditionFunc: 自定义条件函数，用于添加额外的查询条件
//   - 如果为 nil，则不添加额外条件
//   - 函数签名: func(ctx context.Context, db *gorm.DB) *gorm.DB
//
// 返回:
//   - int64: 符合条件的记录数量
//   - error: 错误信息
//
// 使用示例:
//
//	// 示例1: 添加自定义条件统计
//	count, err := dao.SelectCountWithAppendConditionFunc(ctx, queryDto, func(ctx context.Context, db *gorm.DB) *gorm.DB {
//	    return db.Where("status = ?", "active")
//	})
//
//	// 示例2: 不添加额外条件
//	count, err := dao.SelectCountWithAppendConditionFunc(ctx, queryDto, nil)
//
// 说明:
//   - 自定义条件函数在基础查询条件之后执行
//   - 适用于需要动态添加复杂查询条件的统计场景
//   - IN 查询条件校验规则同 SelectCount 方法
func (dao *TUserDao) SelectCountWithAppendConditionFunc(ctx context.Context, queryDto *query.TUserDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&entity.TUser{})

	// 应用查询条件
	var err error
	db, err = dao.buildTUserQueryCondition(db, queryDto)
	if err != nil {
		return 0, err
	}

	// 应用自定义条件函数
	if appendConditionFunc != nil {
		db = appendConditionFunc(ctx, db)
	}

	err = db.Count(&count).Error
	return count, err
}

// ==================== 基础插入方法 ====================

// Insert 单行插入
// 参数:
//   - ctx: 上下文对象
//   - poBean: 要插入的PO对象
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 插入所有字段，包括零值字段
//   - 自增主键会在插入后自动填充到poBean中
func (dao *TUserDao) Insert(ctx context.Context, poBean *entity.TUser) error {
	if poBean == nil {
		return fmt.Errorf("插入对象不能为空")
	}
	return dao.WithContext(ctx).Create(poBean).Error
}

// InsertBatch 批量插入
// 参数:
//   - ctx: 上下文对象
//   - poBeanList: 要插入的PO对象列表
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 批量插入所有记录，在一个事务中执行
//   - 自增主键会在插入后自动填充到各个poBean中
func (dao *TUserDao) InsertBatch(ctx context.Context, poBeanList []*entity.TUser) error {
	if len(poBeanList) == 0 {
		return fmt.Errorf("批量插入列表不能为空")
	}
	return dao.WithContext(ctx).Create(&poBeanList).Error
}

// InsertOrUpdateNullable 插入或更新（会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 要插入或更新的PO对象
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 如果记录不存在（根据主键判断），则执行插入操作
//  2. 如果记录已存在，则执行全字段更新操作
//  3. **重要**: 更新时会用传入对象的所有字段值覆盖数据库中的值，包括零值（nil、""、0、false等）
//     例如: 如果 poBean.Content = nil，会将数据库中的 content 字段更新为 NULL
//     例如: 如果 poBean.ArtifactName = ""，会将数据库中的 artifactName 字段更新为空字符串
//  4. 这种行为适用于需要"完整替换"记录的场景
//  5. 如果不希望零值覆盖数据库中的非零值，应使用 UpdateByXxx 等方法（内部使用 Updates）
func (dao *TUserDao) InsertOrUpdateNullable(ctx context.Context, poBean *entity.TUser) error {
	if poBean == nil {
		return fmt.Errorf("插入或更新对象不能为空")
	}
	// 使用 GORM 的 Save 方法:
	// - 根据主键判断记录是否存在
	// - 存在则更新所有字段（包括零值字段）
	// - 不存在则插入新记录
	return dao.WithContext(ctx).Save(poBean).Error
}

// InsertOrUpdateBatchNullable 批量插入或更新（会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBeanList: 要插入或更新的PO对象列表
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 对列表中的每条记录，根据主键判断是插入还是更新
//  2. 如果记录不存在，则执行插入操作
//  3. 如果记录已存在，则执行全字段更新操作
//  4. **重要**: 更新时会用传入对象的所有字段值覆盖数据库中的值，包括零值（nil、""、0、false等）
//     这意味着如果某个字段在传入对象中为零值，会将数据库中对应字段更新为零值
//  5. 批量操作在一个事务中执行，要么全部成功，要么全部失败
//  6. 适用场景: 需要完整替换多条记录的场景
//  7. 性能提示: 批量操作比逐条调用 InsertOrUpdateNullable 效率更高
//  8. 如果不希望零值覆盖，建议逐条调用 UpdateByXxx 等方法
func (dao *TUserDao) InsertOrUpdateBatchNullable(ctx context.Context, poBeanList []*entity.TUser) error {
	if len(poBeanList) == 0 {
		return fmt.Errorf("批量插入或更新列表不能为空")
	}
	// 使用 GORM 的 Save 方法批量保存:
	// - 对每条记录根据主键判断是插入还是更新
	// - 更新时会覆盖所有字段（包括零值字段）
	// - 在一个事务中执行，保证原子性
	return dao.WithContext(ctx).Save(&poBeanList).Error
}

// ==================== 主键索引方法 ====================

// SelectById 根据主键Id查询单条记录
// 参数:
//   - ctx: 上下文对象
//   - id: 主键值
//
// 返回:
//   - *entity.TUser: 查询结果，如果不存在返回nil
//   - error: 错误信息
func (dao *TUserDao) SelectById(ctx context.Context, id uint64) (*entity.TUser, error) {
	var resultBean entity.TUser
	err := dao.WithContext(ctx).Where("id = ?", id).First(&resultBean).Error
	if err != nil {
		return nil, err
	}
	return &resultBean, nil
}

// SelectByIdList 根据主键Id列表批量查询
// 参数:
//   - ctx: 上下文对象
//   - idList: 主键值列表
//
// 返回:
//   - [] *entity.TUser: 查询结果列表
//   - error: 错误信息
func (dao *TUserDao) SelectByIdList(ctx context.Context, idList []uint64) ([]*entity.TUser, error) {
	if len(idList) == 0 {
		return []*entity.TUser{}, nil
	}
	var resultList []*entity.TUser
	err := dao.WithContext(ctx).Where("id IN ?", idList).Find(&resultList).Error
	return resultList, err
}

// UpdateById 根据主键Id更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - id: 主键值
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 根据指定的 id 更新记录
//  2. **重要**: 只更新非零值字段，零值字段会被忽略，不会覆盖数据库中的值
//     例如: 如果 poBean.Content = nil，不会更新数据库中的 content 字段
//     例如: 如果 poBean.ArtifactName = ""，不会更新数据库中的 artifactName 字段
//  3. 这种行为适用于"部分更新"场景，保留数据库中未传入的字段值
//  4. 如果需要将某个字段更新为零值，应使用 UpdateByIdWithMap 方法显式指定
//  5. 与 InsertOrUpdateNullable 的区别: InsertOrUpdateNullable 会用零值覆盖，UpdateById 不会
func (dao *TUserDao) UpdateById(ctx context.Context, poBean *entity.TUser, id uint64) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	// 使用 Updates 方法:
	// - 只更新结构体中的非零值字段
	// - 零值字段会被忽略，保留数据库中的原值
	// - 适合部分更新场景
	return dao.WithContext(ctx).Model(&entity.TUser{}).Where("id = ?", id).Updates(poBean).Error
}

// UpdateByIdWithMap 根据主键Id使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - id: 主键值
//   - updatedMap: 要更新的字段Map，key为字段名（数据库列名），value为字段值
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 根据指定的 id 更新记录
//  2. 使用 map 可以显式指定要更新的字段，包括零值字段
//  3. **重要**: 与 UpdateById 不同，使用 map 可以将字段更新为零值
//     例如: updatedMap["content"]  = nil 会将 content 字段更新为 NULL
//     例如: updatedMap["artifactName"]  = "" 会将 artifactName 字段更新为空字符串
//  4. 只更新 map 中指定的字段，未指定的字段保持不变
//  5. 适用场景: 需要精确控制更新哪些字段，包括需要将某些字段设置为零值的场景
//  6. 使用建议: 字段名必须与数据库列名一致（或使用 GORM 的字段映射名）
func (dao *TUserDao) UpdateByIdWithMap(ctx context.Context, id uint64, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	// 使用 Updates 方法配合 map:
	// - 可以显式更新零值字段
	// - 只更新 map 中指定的字段
	// - 提供最精确的字段更新控制
	return dao.WithContext(ctx).Model(&entity.TUser{}).Where("id = ?", id).Updates(updatedMap).Error
}

// UpdateByIdWithCondition 根据主键Id和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - id: 主键值
//   - conditionMap: 额外的查询条件Map，key为字段名，value为字段值
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 根据指定的 id 和额外的条件更新记录
//  2. 只更新非零值字段，零值字段会被忽略
//  3. 适用场景: 需要在主键基础上增加额外的更新条件，如乐观锁、状态检查等
//  4. 示例: conditionMap["version"]  = 1 可以实现乐观锁，只有版本号匹配才更新
func (dao *TUserDao) UpdateByIdWithCondition(ctx context.Context, poBean *entity.TUser, id uint64, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&entity.TUser{}).Where("id = ?", id)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByIdWithMapAndCondition 根据主键Id和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - id: 主键值
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 根据指定的 id 和额外的条件更新记录
//  2. 使用 map 可以显式指定要更新的字段，包括零值字段
//  3. 提供最灵活的更新控制方式
func (dao *TUserDao) UpdateByIdWithMapAndCondition(ctx context.Context, id uint64, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&entity.TUser{}).Where("id = ?", id)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteById 根据主键Id删除
// 参数:
//   - ctx: 上下文对象
//   - id: 主键值
//
// 返回:
//   - error: 错误信息
func (dao *TUserDao) DeleteById(ctx context.Context, id uint64) error {
	return dao.WithContext(ctx).Where("id = ?", id).Delete(&entity.TUser{}).Error
}

// ==================== 唯一索引 uk_username 方法 ====================

// SelectByUsername 根据唯一索引uk_username查询单条记录
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//
// 返回:
//   - *entity.TUser: 查询结果，如果不存在返回nil
//   - error: 错误信息
func (dao *TUserDao) SelectByUsername(ctx context.Context, username string) (*entity.TUser, error) {
	var resultBean entity.TUser
	err := dao.WithContext(ctx).Where("username = ?", username).First(&resultBean).Error
	if err != nil {
		return nil, err
	}
	return &resultBean, nil
}

// SelectByUsernameList 根据唯一索引uk_username批量查询
// 参数:
//   - ctx: 上下文对象
//   - usernameList: 用户名列表
//
// 返回:
//   - [] *entity.TUser: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 虽然是唯一索引，但支持批量查询多个唯一键对应的记录
//   - 适用场景: 根据多个唯一键（如用户名列表）批量查询记录
func (dao *TUserDao) SelectByUsernameList(ctx context.Context, usernameList []string) ([]*entity.TUser, error) {
	if len(usernameList) == 0 {
		return []*entity.TUser{}, nil
	}
	var resultList []*entity.TUser
	err := dao.WithContext(ctx).Where("username IN ?", usernameList).Find(&resultList).Error
	return resultList, err
}

// UpdateByUsername 根据唯一索引uk_username更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - username: 用户名
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
func (dao *TUserDao) UpdateByUsername(ctx context.Context, poBean *entity.TUser, username string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&entity.TUser{}).Where("username = ?", username).Updates(poBean).Error
}

// UpdateByUsernameWithMap 根据唯一索引uk_username使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
func (dao *TUserDao) UpdateByUsernameWithMap(ctx context.Context, username string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&entity.TUser{}).Where("username = ?", username).Updates(updatedMap).Error
}

// UpdateByUsernameWithCondition 根据唯一索引uk_username和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - username: 用户名
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在唯一键基础上增加额外的更新条件
func (dao *TUserDao) UpdateByUsernameWithCondition(ctx context.Context, poBean *entity.TUser, username string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&entity.TUser{}).Where("username = ?", username)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByUsernameWithMapAndCondition 根据唯一索引uk_username和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *TUserDao) UpdateByUsernameWithMapAndCondition(ctx context.Context, username string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&entity.TUser{}).Where("username = ?", username)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByUsername 根据唯一索引uk_username删除
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//
// 返回:
//   - error: 错误信息
func (dao *TUserDao) DeleteByUsername(ctx context.Context, username string) error {
	return dao.WithContext(ctx).Where("username = ?", username).Delete(&entity.TUser{}).Error
}

// ==================== 唯一索引 uk_email 方法 ====================

// SelectByEmail 根据唯一索引uk_email查询单条记录
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//
// 返回:
//   - *entity.TUser: 查询结果，如果不存在返回nil
//   - error: 错误信息
func (dao *TUserDao) SelectByEmail(ctx context.Context, email string) (*entity.TUser, error) {
	var resultBean entity.TUser
	err := dao.WithContext(ctx).Where("email = ?", email).First(&resultBean).Error
	if err != nil {
		return nil, err
	}
	return &resultBean, nil
}

// SelectByEmailList 根据唯一索引uk_email批量查询
// 参数:
//   - ctx: 上下文对象
//   - emailList: 邮箱列表
//
// 返回:
//   - [] *entity.TUser: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 虽然是唯一索引，但支持批量查询多个唯一键对应的记录
//   - 适用场景: 根据多个唯一键（如用户名列表）批量查询记录
func (dao *TUserDao) SelectByEmailList(ctx context.Context, emailList []string) ([]*entity.TUser, error) {
	if len(emailList) == 0 {
		return []*entity.TUser{}, nil
	}
	var resultList []*entity.TUser
	err := dao.WithContext(ctx).Where("email IN ?", emailList).Find(&resultList).Error
	return resultList, err
}

// UpdateByEmail 根据唯一索引uk_email更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - email: 邮箱
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
func (dao *TUserDao) UpdateByEmail(ctx context.Context, poBean *entity.TUser, email string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&entity.TUser{}).Where("email = ?", email).Updates(poBean).Error
}

// UpdateByEmailWithMap 根据唯一索引uk_email使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
func (dao *TUserDao) UpdateByEmailWithMap(ctx context.Context, email string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&entity.TUser{}).Where("email = ?", email).Updates(updatedMap).Error
}

// UpdateByEmailWithCondition 根据唯一索引uk_email和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - email: 邮箱
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在唯一键基础上增加额外的更新条件
func (dao *TUserDao) UpdateByEmailWithCondition(ctx context.Context, poBean *entity.TUser, email string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&entity.TUser{}).Where("email = ?", email)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByEmailWithMapAndCondition 根据唯一索引uk_email和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *TUserDao) UpdateByEmailWithMapAndCondition(ctx context.Context, email string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&entity.TUser{}).Where("email = ?", email)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByEmail 根据唯一索引uk_email删除
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//
// 返回:
//   - error: 错误信息
func (dao *TUserDao) DeleteByEmail(ctx context.Context, email string) error {
	return dao.WithContext(ctx).Where("email = ?", email).Delete(&entity.TUser{}).Error
}

// ==================== 原生SQL执行方法 ====================

// ExecSql 执行原生SQL查询
// 参数:
//   - ctx: 上下文对象
//   - recvPtr: 接收查询结果的指针（必须是指针类型）
//   - 查询单条记录时传入结构体指针，如 &entity.TNode{}
//   - 查询多条记录时传入 slice 指针，如 &[] *entity.TNode{}
//   - sql: SQL语句，支持占位符 ?
//   - args: SQL参数，按顺序对应 SQL 中的占位符
//
// 返回:
//   - error: 错误信息，查询失败或记录不存在时返回错误
//
// 使用示例:
//
//	// 示例1: 查询单条记录
//	var result entity.TUser
//	err := dao.ExecSql(ctx, &result, "SELECT * FROM t_user WHERE id = ?", 1)
//	if err != nil {
//	    // 处理错误（包括记录不存在的情况）
//	    return err
//	}
//
//	// 示例2: 查询多条记录
//	var resultList [] *entity.TUser
//	err := dao.ExecSql(ctx, &resultList, "SELECT * FROM t_user WHERE skill_id = ?", "skill123")
//	if err != nil {
//	    return err
//	}
//
//	// 示例3: 查询聚合结果
//	type CountResult struct {
//	    SkillId string `gorm:"column:skill_id"`
//	    Count   int64  `gorm:"column:count"`
//	}
//	var countList [] *CountResult
//	err := dao.ExecSql(ctx, &countList, "SELECT skill_id, COUNT(*) as count FROM t_user GROUP BY skill_id")
//
// 注意事项:
//   - recvPtr 必须传指针，否则无法接收查询结果
//   - 查询单条记录时，如果返回多行，只会取第一行
//   - 查询多条记录时，如果没有结果，会返回空 slice（不是 nil）
//   - 结构体字段需要通过 gorm 标签与数据库列名匹配
//   - 如果取了别名，gorm 的 column 标签需要和 SQL 取的别名一致
//   - gorm 的 column 标签默认为下划线格式
func (dao *TUserDao) ExecSql(ctx context.Context, recvPtr any, sql string, args ...any) error {
	return dao.WithContext(ctx).Raw(sql, args...).Scan(recvPtr).Error
}

// ==================== 辅助方法 ====================

// getValidOrderByFields 获取允许排序的字段白名单
// 返回:
//   - map[string] bool: 字段白名单，key为字段名，value为true表示允许排序
func (dao *TUserDao) getValidOrderByFields() map[string]bool {
	return map[string]bool{
		"id":         true,
		"username":   true,
		"email":      true,
		"password":   true,
		"role":       true,
		"createTime": true,
		"updateTime": true,
	}
}

// isValidOrderBy 验证排序字符串是否安全（基于字段白名单）
// 支持格式:
//   - 单字段: id DESC
//   - 多字段: id DESC, createTime ASC
//
// 参数:
//   - orderBy: 排序字符串
//
// 返回:
//   - true: 排序字符串合法且所有字段都在白名单中
//   - false: 排序字符串不合法或包含非白名单字段
func (dao *TUserDao) isValidOrderBy(orderBy string) bool {
	if orderBy == "" {
		return false
	}

	// 获取字段白名单
	validFields := dao.getValidOrderByFields()

	// 按逗号分割多个排序字段
	orderParts := strings.Split(orderBy, ",")

	for _, part := range orderParts {
		part = strings.TrimSpace(part)
		if part == "" {
			return false
		}

		// 按空格分割字段名和排序方向
		tokens := strings.Fields(part)
		if len(tokens) == 0 || len(tokens) > 2 {
			// 格式错误: 必须是 "字段名" 或 "字段名 方向"
			return false
		}

		// 验证字段名是否在白名单中
		fieldName := tokens[0]
		if !validFields[fieldName] {
			// 字段不在白名单中
			return false
		}

		// 如果指定了排序方向，验证是否为 ASC 或 DESC
		if len(tokens) == 2 {
			direction := strings.ToUpper(tokens[1])
			if direction != "ASC" && direction != "DESC" {
				// 排序方向无效
				return false
			}
		}
	}

	return true
}
