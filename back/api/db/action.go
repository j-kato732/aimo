package db

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "github.com/j-kato732/aimo/proto"
)

const (
	db_path = "./db/test.db"
)

func PostAim(ctx context.Context, aim_model *pb.AimModel) (*pb.PostAimResponse_PostAimResult, error) {
	aim_model_orm, err := aim_model.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	con, err := db.DB()
	defer con.Close()

	isExist := db.Migrator().HasTable("aimModel")
	if isExist == false {
		db.AutoMigrate(&aim_model_orm)
	}

	if err := db.Create(&aim_model_orm).Error; err != nil {
		return nil, err
	}

	return &pb.PostAimResponse_PostAimResult{
		Id: aim_model_orm.Id,
	}, nil
}
