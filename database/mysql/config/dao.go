package config

import (
	"context"
	"gorm.io/gorm"
	"konfig-go/common/logger"
	"konfig-go/common/storage/mysql"
	"konfig-go/pb"
)

type Config struct {
	gorm.Model
	CollectionId int64  `gorm:"column:collection_id;type:INT(11);NOT NULL"`
	Name         string `gorm:"column:name;type:VARCHAR(255);"`
	Key          string `gorm:"column:key;type:VARCHAR(255);NOT NULL"`
	Value        string `gorm:"column:value;type:TEXT;"`
	Version      int8   `gorm:"column:version;type:TINYINT(2);NOT NULL"`
	UpdatedBy    string `gorm:"column:updated_by;type:VARCHAR(40);"`
	CreatedBy    string `gorm:"column:created_by;type:VARCHAR(40);"`
}

func (table *Config) ConvertToDTO() *pb.Config {
	return &pb.Config{
		ID:           int64(table.ID),
		CollectionID: table.CollectionId,
		Name:         table.Name,
		Key:          table.Key,
		Value:        table.Value,
		CreatedBy:    table.CreatedBy,
		UpdatedBy:    table.UpdatedBy,
		CreatedAt:    table.CreatedAt.Unix(),
		UpdatedAt:    table.UpdatedAt.Unix(),
	}
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
func (dao *ConfigDAO) Create(ctx context.Context, config *Config) error {
	return dao.db.WithContext(ctx).Create(config).Error
}

// Update updates an existing config
func (dao *ConfigDAO) Update(ctx context.Context, config *Config) error {
	return dao.db.WithContext(ctx).Save(config).Error
}

// Delete deletes a config
func (dao *ConfigDAO) Delete(ctx context.Context, config *Config) error {
	return dao.db.WithContext(ctx).Delete(config).Error
}

// GetByID gets a config by id
func (dao *ConfigDAO) GetByID(ctx context.Context, id uint) (*Config, error) {
	var config Config
	if err := dao.db.WithContext(ctx).First(&config, id).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// GetAll gets all configs
func (dao *ConfigDAO) GetAll(ctx context.Context) ([]*Config, error) {
	var configs []*Config
	if err := dao.db.WithContext(ctx).Find(&configs).Error; err != nil {
		return nil, err
	}
	return configs, nil
}

func (dao *ConfigDAO) GetByCollectionID(ctx context.Context, collectionId int32) ([]*Config, error) {
	var configs []*Config
	err := dao.db.WithContext(ctx).Where("collection_id = ?", collectionId).Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func (dao *ConfigDAO) SearchWithPaging(ctx context.Context, collectionID int64, name string, sort int64, pageNum int, pageSize int) ([]*Config, error) {
	logger.Logger.Info(ctx, "[config]search with paging, %v_%v_%v_%v_%v_", collectionID, name, sort, pageSize, pageNum)
	var configs []*Config
	offset := (pageNum - 1) * pageSize
	query := dao.db.WithContext(ctx).Offset(offset).Limit(pageSize).Where("collection_id = ?", collectionID)
	if len(name) != 0 {
		nameFuzey := "%" + name + "%"
		query.Where("`name` LIKE ?", nameFuzey)
	}
	mysql.AppendOrderBy(query, "id", sort)
	err := query.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
}

// GetByCollectionIDAndKey gets a config by cid and key to check dup
func (dao *ConfigDAO) GetByCollectionIDAndKey(ctx context.Context, collectionID int64, key string) (*Config, error) {
	var config Config
	if err := dao.db.WithContext(ctx).Where("collection_id = ? AND key = ?", collectionID, key).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (dao *ConfigDAO) CountByCollectionIDs(ctx context.Context, collectionIDs []int64) (map[int64]int64, error) {
	var results []struct {
		CollectionID int64
		Count        int64
	}
	countMap := make(map[int64]int64)

	err := dao.db.WithContext(ctx).Model(&Config{}).
		Select("collection_id, COUNT(*) as count").
		Where("collection_id IN ?", collectionIDs).
		Group("collection_id").
		Find(&results).Error
	if err != nil {
		logger.Logger.Error(ctx, "error to count configs by collectionIDs, err:%v", err)
		return countMap, nil
	}

	for _, result := range results {
		countMap[result.CollectionID] = result.Count
	}
	return countMap, nil
}
