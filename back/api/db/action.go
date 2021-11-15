package db

import (
	"context"
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	errdetails "github.com/j-kato732/aimo/errors"
	pb "github.com/j-kato732/aimo/proto"
)

const (
	db_path = "./db/test.db"
)

var (
	aim_model              pb.AimModelORM
	achievement_mean_model pb.AchievementMeanModelORM
)

/*
/aim
*/

func GetAims(ctx context.Context, request *pb.GetAimsRequest) ([]*pb.AimModel, error) {
	var aims_ORM []*pb.AimModelORM

	// dbに接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	query := "period = ? AND user_id = ?"
	result := db.Where(query, request.GetPeriod(), request.GetUserId()).Find(&aims_ORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// ORMからPBへ変換
	var response []*pb.AimModel
	for _, aim_ORM := range aims_ORM {
		result, err := aim_ORM.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		response = append(response, &result)
	}
	return response, nil
}

/*
/aim
*/

func GetAim(ctx context.Context, request *pb.AimModel) (*pb.AimModel, error) {
	var responseORM *pb.AimModelORM

	// dbに接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	query := "period = ? AND user_id = ? AND aim_number = ?"
	result := db.Where(query, request.GetPeriod(), request.GetUserId(), request.GetAimNumber()).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// ORMからPBへ変換
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &response, nil
}

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
	if err != nil {
		return nil, err
	}
	defer con.Close()

	// モデルに対応するテーブルを作成
	isExist := db.Migrator().HasTable("aimModel")
	if !isExist {
		db.AutoMigrate(aim_model)
	}

	// レコード存在チェック
	_, err = GetAim(ctx, request_aim_model)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return nil, errdetails.ErrRecordExist
	}

	// 目標達成を新規作成
	if err = db.Create(&request_aim_model_orm).Error; err != nil {
		return nil, err
	}

	return &pb.PostAimResponse_PostAimResult{
		Id: request_aim_model_orm.Id,
	}, nil
}

func PutAim(ctx context.Context, request *pb.AimModel) error {
	// db接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return err
	}
	con, err := db.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	// aim_modelをORMに変換
	request_ORM, err := request.ToORM(ctx)
	if err != nil {
		return err
	}

	// update実行
	if err = db.Model(&request_ORM).Updates(request_ORM).Error; err != nil {
		return err
	}

	return nil
}

/*
/achievementMeans
*/

func GetAchievementMeans(ctx context.Context, request *pb.AchievementMeanModel) ([]*pb.AchievementMeanModel, error) {
	// db接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	var achievement_means_orm []*pb.AchievementMeanModelORM
	query := "period = ? AND user_id = ? AND aim_number = ?"
	if err = db.Where(query, request.GetPeriod(), request.GetUserId(), request.GetAimNumber()).Find(&achievement_means_orm).Error; err != nil {
		return nil, err
	}

	var achievement_means []*pb.AchievementMeanModel
	for _, achievemachievement_mean_orm := range achievement_means_orm {
		result, err := achievemachievement_mean_orm.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		achievement_means = append(achievement_means, &result)
	}

	log.Printf("%+v", achievement_means)

	return achievement_means, nil
}

// func PostAchievementMean(ctx context.Context, request_achievement_mean *pb.AchievementMeanModel) (int64, error) {
// 	// db接続
// 	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
// 	if err != nil {
// 		return 0, err
// 	}
// 	con, err := db.DB()
// 	defer con.Close()

// 	// モデルに対応するテーブルを作成
// 	isExist := db.Migrator().HasTable("achievementMeanModel")
// 	if isExist == false {
// 		db.AutoMigrate(achievement_mean_model)
// 	}

// 	// requestをORMに変換、aim_idを追加し、レコード作成
// 	request_orm_achievement_mean, err := request_achievement_mean.ToORM(ctx)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// create record
// 	if err = db.Create(&request_orm_achievement_mean).Error; err != nil {
// 		return 0, err
// 	}
// 	return request_orm_achievement_mean.Id, nil
// }

// func UpdateAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) error {
// 	// db接続
// 	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}
// 	con, err := db.DB()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}
// 	defer con.Close()

// 	// requestをORMに変換
// 	request_ORM, err := request.ToORM(ctx)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	if err = db.Model(&request_ORM).Updates(request_ORM).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

/*
/achievementMean
*/

func GetAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.AchievementMeanModel, error) {
	var responseORM *pb.AchievementMeanModelORM

	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer con.Close()

	// get実行
	query := "period = ? AND user_id = ? AND aim_number = ? AND achievement_mean_number = ?"
	result := db.Where(query, request.Period, request.UserId, request.AimNumber, request.AchievementMeanNumber).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		return 0, err
	}
	defer con.Close()

	// モデルに対応するテーブルを作成
	isExist := db.Migrator().HasTable("achievementMeanModel")
	if !isExist {
		db.AutoMigrate(achievement_mean_model)
	}

	// レコード存在チェック
	_, err = GetAchievementMean(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// requestをORMに変換、aim_idを追加し、レコード作成
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
	}
	// create record
	if err = db.Create(&requestORM).Error; err != nil {
		return 0, err
	}

	return requestORM.Id, nil
}

func PutAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) error {
	// achievementMeanModelをORMに変換
	request_ORM, err := request.ToORM(ctx)
	if err != nil {
		return err
	}

	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer con.Close()

	// update実行
	if err = db.Model(&request_ORM).Updates(request_ORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

/*
// personalEva
*/
func GetPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.PersonalEvaModel, error) {
	var responseORM *pb.PersonalEvaModelORM = new(pb.PersonalEvaModelORM)

	// personalEvaModelをORMに変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// db接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	//get実行
	query := "aim_id = ?"
	result := db.Where(query, requestORM.AimId).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (int64, error) {
	// DB接続
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetPersonalEva(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// create実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

/*
/evaluationBefore
*/
func GetEvaluationBefore(ctx context.Context, request *pb.EvaluationBeforeModel) (*pb.EvaluationBeforeModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.EvaluationBeforeModelORM

	//get実行
	query := "aim_id = ? AND evaluator_number = ?"
	result := db.Where(query, requestORM.AimId, requestORM.EvaluatorNumber).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostEvaluationBefore(ctx context.Context, request *pb.EvaluationBeforeModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetEvaluationBefore(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutEvaluationBefore(ctx context.Context, request *pb.EvaluationBeforeModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func GetEvaluation(ctx context.Context, request *pb.EvaluationModel) (*pb.EvaluationModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.EvaluationModelORM

	//get実行
	query := "aim_id = ? AND evaluator_number = ?"
	result := db.Where(query, requestORM.AimId, requestORM.EvaluatorNumber).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostEvaluation(ctx context.Context, request *pb.EvaluationModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetEvaluation(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutEvaluation(ctx context.Context, request *pb.EvaluationModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func GetComprehensiveComment(ctx context.Context, request *pb.ComprehensiveCommentModel) (*pb.ComprehensiveCommentModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.ComprehensiveCommentModelORM

	//get実行
	query := "user_id = ? AND period = ?"
	result := db.Where(query, requestORM.UserId, requestORM.Period).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostComprehensiveComment(ctx context.Context, request *pb.ComprehensiveCommentModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetComprehensiveComment(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutComprehensiveComment(ctx context.Context, request *pb.ComprehensiveCommentModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

/*
/user
*/
func GetUser(ctx context.Context, request *pb.UserModel) (*pb.UserModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.UserModelORM

	//get実行
	query := "auth_id = ? AND period = ?"
	result := db.Where(query, requestORM.AuthId, requestORM.Period).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostUser(ctx context.Context, request *pb.UserModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetUser(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutUser(ctx context.Context, request *pb.UserModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

/*
/policy
*/
func GetPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.PolicyModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.PolicyModelORM

	//get実行
	query := "period = ?"
	result := db.Where(query, requestORM.Period).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostPolicy(ctx context.Context, request *pb.PolicyModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetPolicy(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutPolicy(ctx context.Context, request *pb.PolicyModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

/*
/departmentGoal
*/
func GetDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.DepartmentGoalModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.DepartmentGoalModelORM

	//get実行
	query := "period = ? AND department_id"
	result := db.Where(query, requestORM.Period, requestORM.DepartmentId).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetDepartmentGoal(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

/*
/role
*/
func GetRole(ctx context.Context, request *pb.RoleModel) (*pb.RoleModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseORM *pb.RoleModelORM

	//get実行
	query := "period = ? AND department_id = ? AND job_id = ?"
	result := db.Where(query, requestORM.Period, requestORM.DepartmentId, requestORM.JobId).Find(&responseORM)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	// convert to PB from ORM
	response, err := responseORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func PostRole(ctx context.Context, request *pb.RoleModel) (int64, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer con.Close()

	// requestをORM型へ変換
	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// personalEvaテーブルが存在しない場合は作成する
	isExist := db.Migrator().HasTable(requestORM.TableName())
	if !isExist {
		db.AutoMigrate(requestORM)
	}

	// レコード存在チェック
	_, err = GetRole(ctx, request)
	if !errors.Is(err, errdetails.ErrNotFound) {
		return 0, errdetails.ErrRecordExist
	}

	// post実行
	if err = db.Create(&requestORM).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return requestORM.Id, nil
}

func PutRole(ctx context.Context, request *pb.RoleModel) error {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer con.Close()

	requestORM, err := request.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	// put実行
	if err = db.Model(&requestORM).Updates(requestORM).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func GetPeriods(ctx context.Context) ([]*pb.PeriodModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	var periods []*pb.PeriodModel
	//get実行
	result := db.Find(&periods)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	return periods, nil
}

func GetDepartments(ctx context.Context) ([]*pb.DepartmentModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	var departments []*pb.DepartmentModel

	//get実行
	result := db.Find(&departments)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	return departments, nil
}

func GetJobs(ctx context.Context) ([]*pb.JobModel, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	con, err := db.DB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer con.Close()

	var jobs []*pb.JobModel

	//get実行
	result := db.Find(&jobs)
	if result.Error != nil {
		return nil, result.Error
	}

	// 取得レコードが0の場合はError
	if result.RowsAffected == 0 {
		return nil, errdetails.ErrNotFound
	}

	return jobs, nil
}
