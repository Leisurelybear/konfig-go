package config

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	CollectionId int64  `gorm:"column:collection_id;type:INT(11);NOT NULL"`
	Name         string `gorm:"column:name;type:VARCHAR(255);"`
	Key          string `gorm:"column:key;type:VARCHAR(255);NOT NULL"`
	Value        string `gorm:"column:value;type:TEXT;"`
	IsDraft      int8   `gorm:"column:is_draft;type:TINYINT(2);NOT NULL"`
	UpdatedBy    string `gorm:"column:updated_by;type:VARCHAR(40);"`
	CreatedBy    string `gorm:"column:created_by;type:VARCHAR(40);"`
}

func (table *Config) TableName() string {
	return "config"
}

type ConfigDAO struct {
	db *gorm.DB
}

func NewConfigDAO(db *gorm.DB) *ConfigDAO {
	return &ConfigDAO{db: db}
}

// Create creates a new config
func (dao *ConfigDAO) Create(config *Config) error {
	return dao.db.Create(config).Error
}

// Update updates an existing config
func (dao *ConfigDAO) Update(config *Config) error {
	return dao.db.Save(config).Error
}

// Delete deletes a config
func (dao *ConfigDAO) Delete(config *Config) error {
	return dao.db.Delete(config).Error
}

// GetByID gets a config by ID
func (dao *ConfigDAO) GetByID(id uint) (*Config, error) {
	var config Config
	if err := dao.db.First(&config, id).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// GetAll gets all configs
func (dao *ConfigDAO) GetAll() ([]*Config, error) {
	var configs []*Config
	if err := dao.db.Find(&configs).Error; err != nil {
		return nil, err
	}
	return configs, nil
}

func (dao *ConfigDAO) GetByCollectionID(collectionId int32) ([]*Config, error) {
	var configs []*Config
	err := dao.db.Where("collection_id = ?", collectionId).Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
}
