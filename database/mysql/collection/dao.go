package collection

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name      string `gorm:"column:name;type:VARCHAR(255);"`
	IsDraft   int32  `gorm:"column:is_draft;type:INT(11);NOT NULL"`
	UpdatedBy string `gorm:"column:updated_by;type:VARCHAR(40);"`
	CreatedBy string `gorm:"column:created_by;type:VARCHAR(40);"`
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

func (dao *CollectionDAO) Create(collection *Collection) error {
	return dao.db.Create(collection).Error
}

func (dao *CollectionDAO) Update(collection *Collection) error {
	return dao.db.Save(collection).Error
}

func (dao *CollectionDAO) Delete(collection *Collection) error {
	return dao.db.Delete(collection).Error
}

func (dao *CollectionDAO) GetById(id int64) (*Collection, error) {
	var collection Collection
	err := dao.db.First(&collection, id).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (dao *CollectionDAO) GetAll() ([]*Collection, error) {
	var collections []*Collection
	err := dao.db.Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}
