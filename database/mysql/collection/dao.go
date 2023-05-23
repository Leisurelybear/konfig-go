package collection

import (
	"context"
	"gorm.io/gorm"
	"konfig-go/common/storage/mysql"
	"konfig-go/pb"
)

const (
	Draft      = 0
	Production = 1
)

type Collection struct {
	gorm.Model
	Name      string `gorm:"column:name;type:VARCHAR(255);"`
	Version   int32  `gorm:"column:version;type:INT(11);NOT NULL"`
	Desc      string `gorm:"column:desc;type:VARCHAR(255);"`
	UpdatedBy string `gorm:"column:updated_by;type:VARCHAR(40);"`
	CreatedBy string `gorm:"column:created_by;type:VARCHAR(40);"`
}

func (table *Collection) ConvertToDTO() *pb.Collection {
	return &pb.Collection{
		ID:        int64(table.ID),
		Name:      table.Name,
		Desc:      table.Desc,
		IsDraft:   table.IsDraft(),
		CreatedBy: table.CreatedBy,
		UpdatedBy: table.UpdatedBy,
		CreatedAt: table.CreatedAt.Unix(),
		UpdatedAt: table.UpdatedAt.Unix(),
	}
}

func (table *Collection) IsDraft() bool {
	return table.Version == Draft
}

func (table *Collection) TableName() string {
	return "collection"
}

type CollectionDAO struct {
	db *gorm.DB
}

func NewCollectionDAO(db *gorm.DB) *CollectionDAO {
	return &CollectionDAO{
		db: db,
	}
}

func (dao *CollectionDAO) Create(ctx context.Context, collection *Collection) error {
	return dao.db.WithContext(ctx).Create(collection).Error
}

func (dao *CollectionDAO) Update(ctx context.Context, collection *Collection) error {
	collection.Version = Draft
	return dao.db.WithContext(ctx).Save(collection).Error
}

func (dao *CollectionDAO) Delete(ctx context.Context, collection *Collection) error {
	return dao.db.WithContext(ctx).Delete(collection).Error
}

func (dao *CollectionDAO) GetById(ctx context.Context, id int64) (*Collection, error) {
	var collection Collection
	err := dao.db.WithContext(ctx).First(&collection, id).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (dao *CollectionDAO) GetAll(ctx context.Context) ([]*Collection, error) {
	var collections []*Collection
	err := dao.db.WithContext(ctx).Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (dao *CollectionDAO) SearchWithPaging(ctx context.Context, name string, createdBy string, sort int64, pageNum int, pageSize int) ([]*Collection, error) {
	var colls []*Collection
	offset := (pageNum - 1) * pageSize
	query := dao.db.WithContext(ctx).Offset(offset).Limit(pageSize)
	if len(name) != 0 {
		nameFuzey := "%" + name + "%"
		query.Where("`name` LIKE ?", nameFuzey)
	}
	if len(createdBy) != 0 {
		createByFuzey := "%" + createdBy + "%"
		query.Where("`created_by` LIKE ?", createByFuzey)
	}
	mysql.AppendOrderBy(query, "id", sort)
	err := query.Find(&colls).Error
	if err != nil {
		return nil, err
	}
	return colls, nil
}
