package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"libs/api-core/features/vote/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
)

func (a *VoteService) Vote(payload dto.Vote, userID string) error {
	var found int64

	err := a.db.Model(&models.VoteModel{}).Where("user_id = ? AND parent_id = ?", userID, payload.ParentID).Count(&found).Error

	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_GET_VOTE", "failed to get vote", err)
	}

	if found > 0 {
		return utils.NewError(fiber.StatusForbidden, "E_FORBIDEN", "you already own this vote", errors.New("you already own this vote"))
	}

	err = a.db.Transaction(func(tx *gorm.DB) error {

		if payload.ParentType == models.ParentTypeQuestion {
			questionModel := models.QuestionModel{}

			err := tx.Model(questionModel).Where("id = ?", payload.ParentID).First(&questionModel).Error
			if err != nil {
				return utils.NewError(fiber.StatusNotFound, "E_NOT_FOUND", "question not found", err)
			}

			votes := questionModel.Votes
			votes++
			if err := tx.Model(questionModel).Update("votes", votes).Error; err != nil {
				return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_VOTES_QUESTION", "update votes on question model failed", err)
			}
		}

		if payload.ParentType == models.ParentTypeAnswer {
			answerModel := models.AnswerModel{}

			err := tx.Model(answerModel).Where("id = ?", payload.ParentID).First(&answerModel).Error
			if err != nil {
				return utils.NewError(fiber.StatusNotFound, "E_ANSWER_NOT_FOUND", "answer not found", err)
			}
			votes := answerModel.Votes
			votes++
			if err := tx.Model(answerModel).Where("ID = ?", answerModel.ID).Update("votes", votes).Error; err != nil {
				return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_VOTES_ANSWER", "update votes on answer model failed", err)
			}
		}

		if payload.ParentType == models.ParentTypeComment {
			commentModel := models.CommentModel{}
			err := tx.Model(commentModel).Where("id = ?", payload.ParentID).First(&commentModel).Error
			if err != nil {
				return utils.NewError(fiber.StatusNotFound, "E_COMMENT_NOT_FOUND", "comment not found", err)
			}
			votes := commentModel.Votes
			votes++
			if err := tx.Model(commentModel).Where("ID = ?", commentModel.ID).Update("votes", votes).Error; err != nil {
				return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_VOTES_COMMENT", "update votes on comment model failed", err)
			}
		}

		voteModel := models.VoteModel{
			ParentType: payload.ParentType,
			ParentID:   payload.ParentID,
			UserID:     userID,
		}
		err := tx.Model(voteModel).Create(&voteModel).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
