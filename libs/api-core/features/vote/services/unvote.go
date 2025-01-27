package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"libs/api-core/features/vote/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
)

func (a *VoteService) UnVote(payload dto.UnVote, userID string) error {
	found := models.VoteModel{}

	err := a.db.Model(models.VoteModel{}).Where("user_id = ? AND parent_id = ?", userID, payload.ParentID).First(&found).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewError(fiber.StatusNotFound, "E_NOT_FOUND", "vote not found", err)
		}
		return utils.NewError(fiber.StatusInternalServerError, "E_GET_VOTE", "failed to get vote", err)
	}

	err = a.db.Transaction(func(tx *gorm.DB) error {

		if found.ParentType == models.ParentTypeQuestion {
			questionModel := models.QuestionModel{}

			err := tx.Model(questionModel).Where("id = ?", payload.ParentID).First(&questionModel).Error
			if err != nil {
				return utils.NewError(fiber.StatusNotFound, "E_NOT_FOUND", "question not found", err)
			}

			votes := questionModel.Votes
			votes--
			if err := tx.Model(questionModel).Where("ID = ?", questionModel.ID).Update("votes", votes).Error; err != nil {
				return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_VOTES_QUESTION", "update votes on question model failed", err)
			}
		}

		if found.ParentType == models.ParentTypeAnswer {
			answerModel := models.AnswerModel{}

			err := tx.Model(answerModel).Where("id = ?", payload.ParentID).First(&answerModel).Error
			if err != nil {
				return utils.NewError(fiber.StatusNotFound, "E_ANSWER_NOT_FOUND", "answer not found", err)
			}
			votes := answerModel.Votes
			votes--
			if err := tx.Model(answerModel).Where("ID = ?", answerModel.ID).Update("votes", votes).Error; err != nil {
				return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_VOTES_ANSWER", "update votes on answer model failed", err)
			}
		}

		if found.ParentType == models.ParentTypeComment {
			commentModel := models.CommentModel{}
			err := tx.Model(commentModel).Where("id = ?", payload.ParentID).First(&commentModel).Error
			if err != nil {
				return utils.NewError(fiber.StatusNotFound, "E_COMMENT_NOT_FOUND", "comment not found", err)
			}
			votes := commentModel.Votes
			votes--
			if err := tx.Model(commentModel).Where("ID = ?", commentModel.ID).Update("votes", votes).Error; err != nil {
				return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_VOTES_COMMENT", "update votes on comment model failed", err)
			}
		}

		voteModel := models.VoteModel{}
		err := tx.Model(voteModel).Where("ID = ?", found.ID).Delete(&voteModel).Error
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
