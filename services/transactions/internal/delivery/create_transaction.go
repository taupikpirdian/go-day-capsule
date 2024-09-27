package delivery

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/internal/delivery/request"
	"github.com/google/uuid"
)

func (t TransactionHandler) CreateTransaction(ctx context.Context, req *transactionPb.Transaction) (*transactionPb.GeneralResponse, error) {
	l := logger.Logger{
		EndPoint:    "/api/v1/transaction/create",
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()
	data := request.NewTransactionRequest(req)

	err := t.transaction.CreateTransaction(ctx, data)
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := &transactionPb.GeneralResponse{
		Data: &transactionPb.General{
			Message: "success",
		},
	}
	l.ResponseData = rsp
	return rsp, nil
}
