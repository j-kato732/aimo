package proto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func (a GetAimsRequest) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Period, validation.Required, validation.Length(6, 6)),
		validation.Field(&a.UserId, validation.Required),
	)
}

func (a AimModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Period, validation.Required, validation.Length(6, 6)),
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.AimItem, validation.Length(0, 500)),
		validation.Field(&a.AchievementLevel, validation.Length(0, 500)),
		validation.Field(&a.AchievementWeight, validation.Length(0, 3)),
		validation.Field(&a.AchievementDifficultyBefore, validation.Length(0, 1)),
		validation.Field(&a.AimNumber, validation.Required),
	)
}

func (a PersonalEvaModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AimId, validation.Required),
		validation.Field(&a.PersonalEvaluation, validation.Length(0, 500)),
		validation.Field(&a.AchievementRate, validation.Length(0, 3)),
		validation.Field(&a.AchievementDifficulty, validation.Length(0, 1)),
	)
}

func (a EvaluationBeforeModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AimId, validation.Required),
		validation.Field(&a.EvaluatorNumber, validation.Required, validation.Length(0, 1)),
		validation.Field(&a.Comment, validation.Length(0, 500)),
	)
}

func (a EvaluationModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AimId, validation.Required),
		validation.Field(&a.EvaluatorNumber, validation.Required, validation.Length(0, 1)),
		validation.Field(&a.Comment, validation.Length(0, 500)),
		validation.Field(&a.AchievementRate, validation.Length(0, 100)),
		validation.Field(&a.AchievementDifficulty, validation.Length(0, 1)),
	)
}

func (a AchievementMeanModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Period, validation.Required, validation.Length(6, 6)),
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.AimNumber, validation.Required, validation.Length(1, 1)),
		validation.Field(&a.AchievementMeanNumber, validation.Required, validation.Length(1, 1)),
		validation.Field(&a.AchievementMean, validation.Length(0, 500)),
		validation.Field(&a.FirstMonth, validation.Length(4, 5)),
	)
}

func (a ComprehensiveCommentModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.Period, validation.Required, validation.Length(6, 6)),
		validation.Field(&a.ComprehensiveComment, validation.Length(0, 500)),
	)
}
