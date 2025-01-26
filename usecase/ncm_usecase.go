package usecase

import (
	"meditrack/model"
	"meditrack/repository"
)

type NCMUsecase struct {
	repository repository.NCMRepository
}

func NewNCMUsecase(repo repository.NCMRepository) NCMUsecase {
	return NCMUsecase{
		repository: repo,
	}
}

func (n *NCMUsecase) GetAllNCM() ([]model.NCM, error) {
	ncm, err := n.repository.GetAllNCM()
	if err != nil {
		return []model.NCM{}, err
	}

	return ncm, nil
}

func (n *NCMUsecase) GetNCMByCode(code string) (*model.NCM, error) {
	ncm, err := n.repository.GetNCMByCode(code)
	if err != nil {
		return nil, err
	}

	return ncm, nil
}

func (n *NCMUsecase) CreateNCM(ncm model.NCM) (model.NCM, error) {
	createdNCM, err := n.repository.CreateNCM(ncm)
	if err != nil {
		return model.NCM{}, err
	}

	return createdNCM, nil
}
