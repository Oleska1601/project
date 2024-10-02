package usecase

import (
	"project/internal/entity"
)

func GetLastTenOperations(responseData entity.ResponseData) []entity.Operation {
	start := responseData.NumsOfOperations - 10
	if start < 0 {
		start = 0
	}
	return responseData.Operations[start:]
}
