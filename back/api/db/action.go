package db

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "github.com/j-kato732/aimo/proto"
)

const (
	db_path = "./db/test.db"
)

var (
	aim_model              pb.AimModelORM
	achievement_mean_model pb.AchievementMeanModelORM
)

func PostAim(ctx context.Context, request_aim_model *pb.AimModel) (*pb.PostAimResponse_PostAimResult, error) {

	// 目標項目をORMに変換
	request_aim_model_orm, err := request_aim_model.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// dbに接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	con, err := db.DB()
	defer con.Close()

	// モデルに対応するテーブルを作成
	isExist := db.Migrator().HasTable("aimModel")
	if isExist == false {
		db.AutoMigrate(aim_model)
	}

	// 目標達成を新規作成
	if err = db.Create(&request_aim_model_orm).Error; err != nil {
		return nil, err
	}

	log.Println(request_aim_model_orm.Id)

	return &pb.PostAimResponse_PostAimResult{
		Id: request_aim_model_orm.Id,
	}, nil
}

func PostAchievementMean(ctx context.Context, request_achievement_mean *pb.AchievementMeanModel) (int64, error) {
	// db接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	con, err := db.DB()
	defer con.Close()

	// モデルに対応するテーブルを作成
	isExist := db.Migrator().HasTable("achievementMeanModel")
	if isExist == false {
		db.AutoMigrate(achievement_mean_model)
	}

	// requestをORMに変換、aim_idを追加し、レコード作成
	request_orm_achievement_mean, err := request_achievement_mean.ToORM(ctx)
	if err != nil {
		log.Println(err)
	}
	// create record
	if err = db.Create(&request_orm_achievement_mean).Error; err != nil {
		return 0, err
	}
	return request_orm_achievement_mean.Id, nil
}
