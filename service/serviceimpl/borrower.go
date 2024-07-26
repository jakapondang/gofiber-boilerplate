package serviceimpl

import (
	"context"
	"goamartha/common"
	"goamartha/domain/entity"
	"goamartha/domain/model"
	"goamartha/exception"
	"goamartha/repository"
	"goamartha/service"
	"strconv"
	"time"
)

func NewBorrowerServiceImpl(borrowerRepository *repository.BorrowerRepository) service.BorrowerService {
	return &BorrowerServiceImpl{BorrowerRepository: *borrowerRepository}
}

type BorrowerServiceImpl struct {
	repository.BorrowerRepository
}

func (service *BorrowerServiceImpl) FindAll(ctx context.Context) (responses []model.Borrower, err error) {
	entities, err := service.BorrowerRepository.FindAll(ctx)
	for _, entity := range entities {
		responses = append(responses, model.Borrower{
			Id:              entity.Id,
			Username:        entity.Username,
			Amount:          entity.Amount,
			TaxRate:         entity.TaxRate * 100,
			TaxAmount:       entity.TaxAmount,
			FinalAmount:     entity.FinalAmount,
			IsSettled:       entity.IsSettled,
			TransactionDate: entity.TransactionDate.Format("2006-01-02 15:04:05"),
			Created:         entity.Created.Format("2006-01-02 15:04:05"),
		})
	}

	if len(entities) == 0 {
		panic(exception.NotFoundError{
			Message: "Borrowers Not Found",
		})
	}
	return responses, err
}

func (service *BorrowerServiceImpl) FindById(ctx context.Context, id string) (response model.Borrower, err error) {
	valid, err := strconv.Atoi(id)
	requestEntity := entity.Borrower{
		Id: uint(valid),
	}
	err = service.BorrowerRepository.FindById(ctx, &requestEntity)
	response = model.SingleBorrower(requestEntity)
	return response, err
}

func (service *BorrowerServiceImpl) FindByUsername(ctx context.Context, username string) (response model.Borrower, err error) {
	requestEntity := entity.Borrower{
		Username: username,
	}
	err = service.BorrowerRepository.FindByUsername(ctx, &requestEntity)
	response = model.SingleBorrower(requestEntity)

	return response, err
}

func (service *BorrowerServiceImpl) Create(ctx context.Context, request model.BorrowerRequest) (model.Borrower, error) {

	common.Validate(request)
	requestEntity := entity.Borrower{
		Username:        request.Username,
		Amount:          request.Amount,
		TaxRate:         model.PaymentTax,
		TaxAmount:       request.Amount * model.PaymentTax,
		TransactionDate: time.Now(),
	}

	err := service.BorrowerRepository.Insert(ctx, &requestEntity)
	response := model.SingleBorrower(requestEntity)

	return response, err
}

func (service *BorrowerServiceImpl) Updates(ctx context.Context, request *model.Borrower) error {
	const layout = "2006-01-02 15:04:05"
	trxDate, err := time.Parse(layout, request.TransactionDate)
	if err != nil {
		exception.PanicLogging("Error parsing datetime:" + err.Error())
	}
	borrower := entity.Borrower{
		Id:              request.Id,
		Username:        request.Username,
		Amount:          request.Amount,
		TaxRate:         request.TaxRate / 100,
		TaxAmount:       request.TaxAmount,
		FinalAmount:     request.FinalAmount,
		IsSettled:       request.IsSettled,
		TransactionDate: trxDate,
	}
	err = service.BorrowerRepository.Update(ctx, &borrower)

	return err
}
