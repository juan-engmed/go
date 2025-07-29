package service

import (
	"errors"
	"pizzaria/pizza_project/internal/models"
)

func ValidateRating(review *models.Review) error {

	if review.Rating <= 0 || review.Rating > 5 {
		return errors.New("a avaliação aceita apenas notas entre 1 e 5")
	}

	return nil
}