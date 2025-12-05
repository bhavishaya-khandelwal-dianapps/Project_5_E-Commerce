package services

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
	"github.com/razorpay/razorpay-go/utils"
)

// 1. Function to create payment
func CreatePayment(order *models.Order) (*models.Payment, error) {
	// 1. Prepare Razorpay order payload
	data := map[string]interface{}{
		"amount":          float64(order.TotalPrice * 100), // in paise
		"currency":        "INR",
		"receipt":         "ORDER_RCPTID_" + strconv.Itoa(int(order.ID)),
		"payment_capture": 1, // auto-capture
	}

	rzpOrder, err := config.RazorpayClient.Order.Create(data, nil)
	if err != nil {
		return nil, err
	}

	// 2. Create payment record in DB
	payment := &models.Payment{
		OrderId:         order.ID,
		Amount:          order.TotalPrice,
		Status:          "PENDING",
		Method:          "RAZORPAY",
		RazorpayOrderId: rzpOrder["id"].(string),
	}

	if err := repositories.CreatePayment(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

// 2. Function to verify payment
func VerifyPayment(orderId, paymentId, signature string) error {
	// 1. Verify signature using razorpay utility
	params := map[string]interface{}{
		"razorpay_order_id":   orderId,
		"razorpay_payment_id": paymentId,
	}

	if isValid := utils.VerifyPaymentSignature(params, signature, config.RazorpayClient.Auth.Secret); !isValid {
		log.Println("❌ Razorpay signature verification failed")
		return errors.New("payment verification failed")
	}

	// 2. Fetch payment from DB using razorpayOrderId
	payment, err := repositories.GetPaymentByRazorpayOrderId(orderId)
	if err != nil {
		return err
	}

	// 3. Update payment record
	payment.PaymentId = paymentId
	payment.RazorpaySignature = signature
	payment.Status = "SUCCESS"

	err = repositories.UpdatePayment(payment)
	if err != nil {
		return err
	}

	// 4. Update corresponding Order Status
	order, err := repositories.GetOrderById(payment.OrderId)
	if err != nil {
		return err
	}

	order.Status = "PAID"
	err = repositories.UpdateOrder(order)
	if err != nil {
		return err
	}

	return nil
}

// 3. Function to process webhook
type RazorpayWebhookPayload struct {
	Event   string                 `json:"event"`
	Payload map[string]interface{} `json:"payload"`
}

func ProcessWebhookEvent(body []byte, signature string) error {
	var data RazorpayWebhookPayload

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	switch data.Event {
	case "payment.captured":
		return handlePaymentCapture(data, signature)

	case "payment.failed":
		return handlePaymentFailed(data)

	case "order.paid":
		return nil

	default:
		return nil
	}
}

// Payment success, handler
func handlePaymentCapture(data RazorpayWebhookPayload, signature string) error {
	paymentObj := data.Payload["payment"].(map[string]interface{})["entity"].(map[string]interface{})

	razorpayPaymentId := paymentObj["id"].(string)
	orderId := paymentObj["order_id"].(string)

	// Find payment by razorpayOrderId
	payment, err := repositories.GetPaymentByRazorpayOrderId(orderId)
	if err != nil {
		return errors.New("payment not found")
	}

	payment.Status = "SUCCESS"
	payment.PaymentId = razorpayPaymentId
	payment.RazorpaySignature = signature

	order, err := repositories.GetOrderById(payment.OrderId)
	if err != nil {
		return err
	}

	order.Status = "PAID"
	err = repositories.UpdateOrder(order)
	if err != nil {
		return err
	}

	return repositories.UpdatePayment(payment)
}

// Payment failed, handler
func handlePaymentFailed(data RazorpayWebhookPayload) error {
	paymentObj := data.Payload["payment"].(map[string]interface{})["entity"].(map[string]interface{})

	orderId := paymentObj["order_id"].(string)
	payment, err := repositories.GetPaymentByRazorpayOrderId(orderId)
	if err != nil {
		return errors.New("payment not found")
	}

	payment.Status = "FAILED"

	order, err := repositories.GetOrderById(payment.OrderId)
	if err != nil {
		return err
	}

	order.Status = "FAILED"
	err = repositories.UpdateOrder(order)
	if err != nil {
		return err
	}

	return repositories.UpdatePayment(payment)
}
