package serviceimpl

import (
	"context"
	"goamartha/domain/entity"
	"goamartha/domain/model"
	"goamartha/exception"
	"goamartha/repository"
	"goamartha/service"
	"time"
)

func NewPaymentServiceImpl(paymentRepository *repository.PaymentRepository) service.PaymentService {
	return &PaymentServiceImpl{PaymentRepository: *paymentRepository}
}

type PaymentServiceImpl struct {
	repository.PaymentRepository
}

// Make Payment for every terms
func (service *PaymentServiceImpl) MakePayment(ctx context.Context, borrower *model.Borrower) (model.Payment, error) {

	requestEntity := entity.Payment{
		BorrowId:     borrower.Id,
		TotalPayment: paymentPerWeek(borrower.FinalAmount),
		IsPaid:       true,
	}
	if borrower.PaymentSummary.TotalPaid == borrower.FinalAmount { // its already settled
		panic(exception.NotFoundError{
			Message: "Your debt has been settled in full.",
		})
	}

	if borrower.PaymentSummary.TotalDelinquent > 0 { // If delinquent cannot make payment, need repayment
		panic(exception.NotFoundError{
			Message: "Payment transaction failed. Please make a repayment to settle the outstanding balance.",
		})
	}
	err := service.PaymentRepository.Insert(ctx, &requestEntity)
	response := model.SinglePayment(requestEntity)
	//response.TotalPayment = paymentPerWeek(borrower.FinalAmount)
	return response, err
}

// Make RePayment If theres Delinquent
func (service *PaymentServiceImpl) MakeRePayment(ctx context.Context, borrower *model.Borrower) (model.Payment, error) {
	if borrower.PaymentSummary.TotalDelinquent == 0 { // If no delinquent cannot make repayment, use payment
		panic(exception.NotFoundError{
			Message: "There are no overdue bills on your account. Please continue with your regular payments.",
		})
	}
	totalPayment := paymentPerWeek(borrower.FinalAmount) + borrower.PaymentSummary.TotalDelinquent
	requestEntity := entity.Payment{
		BorrowId:     borrower.Id,
		TotalPayment: totalPayment,
		IsPaid:       true,
	}
	err := service.PaymentRepository.Insert(ctx, &requestEntity)
	response := model.SinglePayment(requestEntity)
	//response.TotalPayment = totalPayment
	return response, err
}

// Get Schedule Payment Terms
func (service *PaymentServiceImpl) GetSchedule(ctx context.Context, borrower *model.Borrower) error {
	var paymentTerms []model.PaymentSchedules

	// Define the layout corresponding to the format of the datetime string
	const layout = "2006-01-02 15:04:05"
	// Parse the string into a time.Time object
	week, err := time.Parse(layout, borrower.TransactionDate)
	if err != nil {
		exception.PanicLogging("Error parsing datetime:" + err.Error())
	}
	// Calculate Schedule
	weeks := model.PaymentTerms + 1 // Number of weeks to display
	totalPaid := borrower.PaymentSummary.TotalPaid
	paymentWeek := paymentPerWeek(borrower.FinalAmount)
	var countDel int
	for i := 0; i < weeks; i++ {
		if i > 0 { // start 1 week after borrow
			term := model.PaymentSchedules{
				Week:         i,
				TotalPayment: paymentWeek,
				DueDate:      week.Format("2006-01-02 15:04:05"),
			}
			if totalPaid > 0 { // check payment
				totalPaid = totalPaid - paymentWeek
				term.IsPaid = true
			}
			//dueWeek := week.AddDate(0, 0, 7)
			dueWeek := week
			if time.Now().After(dueWeek) && term.IsPaid == false { // is Deliquent
				term.IsDelinquent = true
				countDel++
			}
			paymentTerms = append(paymentTerms, term)
		}

		week = week.AddDate(0, 0, 7) // Add one week
	}
	borrower.PaymentSchedules = paymentTerms
	borrower.PaymentSummary.TotalDelinquent = paymentPerWeek(borrower.FinalAmount) * float64(countDel)
	return nil
}

// Get Payment Summary
func (service *PaymentServiceImpl) PaymentSummaries(ctx context.Context, borrower *model.Borrower) {
	requestEntity := entity.Payment{
		BorrowId: borrower.Id,
		IsPaid:   true,
	}
	count, _ := service.PaymentRepository.CountPayment(ctx, &requestEntity)
	countPayments := int(count)
	totalPaid := float64(countPayments) * paymentPerWeek(borrower.FinalAmount)
	borrower.PaymentSummary = model.PaymentSummary{
		CountPayment:        countPayments,
		TotalPaid:           totalPaid,
		OutstandingBallance: borrower.FinalAmount - totalPaid,
	}
}

func paymentPerWeek(finalAmount float64) float64 {
	// Payment perweek
	result := finalAmount / model.PaymentTerms
	return result
}
