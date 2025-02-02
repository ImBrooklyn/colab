// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"colab/app/pkg/data/model"
)

func newMenu(db *gorm.DB, opts ...gen.DOOption) menu {
	_menu := menu{}

	_menu.menuDo.UseDB(db, opts...)
	_menu.menuDo.UseModel(&model.Menu{})

	tableName := _menu.menuDo.TableName()
	_menu.ALL = field.NewAsterisk(tableName)
	_menu.ID = field.NewInt64(tableName, "id")
	_menu.CreatedAt = field.NewTime(tableName, "created_at")
	_menu.UpdatedAt = field.NewTime(tableName, "updated_at")
	_menu.DeletedAt = field.NewField(tableName, "deleted_at")
	_menu.Pid = field.NewInt64(tableName, "pid")
	_menu.Title = field.NewString(tableName, "title")
	_menu.Icon = field.NewString(tableName, "icon")
	_menu.URL = field.NewString(tableName, "url")
	_menu.Filepath = field.NewString(tableName, "filepath")
	_menu.Params = field.NewString(tableName, "params")
	_menu.Node = field.NewString(tableName, "node")
	_menu.Sort = field.NewInt32(tableName, "sort")
	_menu.Status = field.NewInt32(tableName, "status")
	_menu.IsInner = field.NewBool(tableName, "is_inner")
	_menu.DefaultValues = field.NewString(tableName, "default_values")
	_menu.ShowSlider = field.NewBool(tableName, "show_slider")

	_menu.fillFieldMap()

	return _menu
}

type menu struct {
	menuDo

	ALL           field.Asterisk
	ID            field.Int64
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Pid           field.Int64
	Title         field.String
	Icon          field.String
	URL           field.String
	Filepath      field.String
	Params        field.String
	Node          field.String
	Sort          field.Int32
	Status        field.Int32
	IsInner       field.Bool
	DefaultValues field.String
	ShowSlider    field.Bool

	fieldMap map[string]field.Expr
}

func (m menu) Table(newTableName string) *menu {
	m.menuDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m menu) As(alias string) *menu {
	m.menuDo.DO = *(m.menuDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *menu) updateTableName(table string) *menu {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Pid = field.NewInt64(table, "pid")
	m.Title = field.NewString(table, "title")
	m.Icon = field.NewString(table, "icon")
	m.URL = field.NewString(table, "url")
	m.Filepath = field.NewString(table, "filepath")
	m.Params = field.NewString(table, "params")
	m.Node = field.NewString(table, "node")
	m.Sort = field.NewInt32(table, "sort")
	m.Status = field.NewInt32(table, "status")
	m.IsInner = field.NewBool(table, "is_inner")
	m.DefaultValues = field.NewString(table, "default_values")
	m.ShowSlider = field.NewBool(table, "show_slider")

	m.fillFieldMap()

	return m
}

func (m *menu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *menu) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 16)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["pid"] = m.Pid
	m.fieldMap["title"] = m.Title
	m.fieldMap["icon"] = m.Icon
	m.fieldMap["url"] = m.URL
	m.fieldMap["filepath"] = m.Filepath
	m.fieldMap["params"] = m.Params
	m.fieldMap["node"] = m.Node
	m.fieldMap["sort"] = m.Sort
	m.fieldMap["status"] = m.Status
	m.fieldMap["is_inner"] = m.IsInner
	m.fieldMap["default_values"] = m.DefaultValues
	m.fieldMap["show_slider"] = m.ShowSlider
}

func (m menu) clone(db *gorm.DB) menu {
	m.menuDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m menu) replaceDB(db *gorm.DB) menu {
	m.menuDo.ReplaceDB(db)
	return m
}

type menuDo struct{ gen.DO }

type IMenuDo interface {
	gen.SubQuery
	Debug() IMenuDo
	WithContext(ctx context.Context) IMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMenuDo
	WriteDB() IMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMenuDo
	Not(conds ...gen.Condition) IMenuDo
	Or(conds ...gen.Condition) IMenuDo
	Select(conds ...field.Expr) IMenuDo
	Where(conds ...gen.Condition) IMenuDo
	Order(conds ...field.Expr) IMenuDo
	Distinct(cols ...field.Expr) IMenuDo
	Omit(cols ...field.Expr) IMenuDo
	Join(table schema.Tabler, on ...field.Expr) IMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMenuDo
	Group(cols ...field.Expr) IMenuDo
	Having(conds ...gen.Condition) IMenuDo
	Limit(limit int) IMenuDo
	Offset(offset int) IMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMenuDo
	Unscoped() IMenuDo
	Create(values ...*model.Menu) error
	CreateInBatches(values []*model.Menu, batchSize int) error
	Save(values ...*model.Menu) error
	First() (*model.Menu, error)
	Take() (*model.Menu, error)
	Last() (*model.Menu, error)
	Find() ([]*model.Menu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Menu, err error)
	FindInBatches(result *[]*model.Menu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Menu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMenuDo
	Assign(attrs ...field.AssignExpr) IMenuDo
	Joins(fields ...field.RelationField) IMenuDo
	Preload(fields ...field.RelationField) IMenuDo
	FirstOrInit() (*model.Menu, error)
	FirstOrCreate() (*model.Menu, error)
	FindByPage(offset int, limit int) (result []*model.Menu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m menuDo) Debug() IMenuDo {
	return m.withDO(m.DO.Debug())
}

func (m menuDo) WithContext(ctx context.Context) IMenuDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m menuDo) ReadDB() IMenuDo {
	return m.Clauses(dbresolver.Read)
}

func (m menuDo) WriteDB() IMenuDo {
	return m.Clauses(dbresolver.Write)
}

func (m menuDo) Session(config *gorm.Session) IMenuDo {
	return m.withDO(m.DO.Session(config))
}

func (m menuDo) Clauses(conds ...clause.Expression) IMenuDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m menuDo) Returning(value interface{}, columns ...string) IMenuDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m menuDo) Not(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m menuDo) Or(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m menuDo) Select(conds ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m menuDo) Where(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m menuDo) Order(conds ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m menuDo) Distinct(cols ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m menuDo) Omit(cols ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m menuDo) Join(table schema.Tabler, on ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m menuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMenuDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m menuDo) RightJoin(table schema.Tabler, on ...field.Expr) IMenuDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m menuDo) Group(cols ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m menuDo) Having(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m menuDo) Limit(limit int) IMenuDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m menuDo) Offset(offset int) IMenuDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m menuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMenuDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m menuDo) Unscoped() IMenuDo {
	return m.withDO(m.DO.Unscoped())
}

func (m menuDo) Create(values ...*model.Menu) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m menuDo) CreateInBatches(values []*model.Menu, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m menuDo) Save(values ...*model.Menu) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m menuDo) First() (*model.Menu, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Menu), nil
	}
}

func (m menuDo) Take() (*model.Menu, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Menu), nil
	}
}

func (m menuDo) Last() (*model.Menu, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Menu), nil
	}
}

func (m menuDo) Find() ([]*model.Menu, error) {
	result, err := m.DO.Find()
	return result.([]*model.Menu), err
}

func (m menuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Menu, err error) {
	buf := make([]*model.Menu, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m menuDo) FindInBatches(result *[]*model.Menu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m menuDo) Attrs(attrs ...field.AssignExpr) IMenuDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m menuDo) Assign(attrs ...field.AssignExpr) IMenuDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m menuDo) Joins(fields ...field.RelationField) IMenuDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m menuDo) Preload(fields ...field.RelationField) IMenuDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m menuDo) FirstOrInit() (*model.Menu, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Menu), nil
	}
}

func (m menuDo) FirstOrCreate() (*model.Menu, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Menu), nil
	}
}

func (m menuDo) FindByPage(offset int, limit int) (result []*model.Menu, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m menuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m menuDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m menuDo) Delete(models ...*model.Menu) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *menuDo) withDO(do gen.Dao) *menuDo {
	m.DO = *do.(*gen.DO)
	return m
}
