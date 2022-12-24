package khalti

import (
	"context"
	"errors"
	"os"

	"github.com/Sugaml/mrc-payment/api/models"
	"github.com/babulalt/go-khalti/khalti"
)

// InitiateTransactionRequest represents a Khalti Initaial Payment Request
type InitiateTransactionRequest struct {
	PubicKey        string `json:"public_key"`
	Amount          uint64 `json:"amount"`
	ProductIdentity string `json:"product_identity"`
	ProductName     string `json:"product_name"`
	ProductUrl      string `json:"product_url"`
}

type VerifyPaymentRequest struct {
	SID    uint   `json:"sid"`
	Token  string `json:"token"`
	Amount uint64 `json:"amount"`
}

func NewKhaltiClient() (*khalti.KhaltiService, error) {
	clientId := os.Getenv("KHALTI_CLIENT_ID")
	secretId := os.Getenv("KHALTI_SECRET_ID")
	if clientId == "" {
		return nil, errors.New("required khalti client id")
	}

	if secretId == "" {
		return nil, errors.New("required khalti secret key")
	}
	return khalti.NewKhaltiClient(clientId, secretId, nil)
}

func NewInitiateTransactionRequest(student *models.Student) (*InitiateTransactionRequest, error) {
	request := &InitiateTransactionRequest{}
	if student == nil {
		return nil, errors.New("required student object")
	}
	clientId := os.Getenv("KHALTI_CLIENT_ID")
	if clientId == "" {
		return nil, errors.New("required khalti client id")
	}
	request.Amount = uint64(student.Course.Fee)
	request.ProductIdentity = "Education"
	request.ProductName = student.Course.Name
	request.ProductUrl = ""
	request.PubicKey = clientId
	return request, nil
}

func VerifyPayment(token string, amount uint64) (*khalti.VerifyTransactionResponse, error) {
	paylaod := &khalti.VerifyTransactionRequest{
		Token:  token,
		Amount: amount,
	}
	khaltiClient, err := NewKhaltiClient()
	if err != nil {
		return nil, err
	}
	return khaltiClient.VerifyTransaction(context.Background(), paylaod)
}
