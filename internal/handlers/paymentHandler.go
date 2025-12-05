package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/services"
	"github.com/gin-gonic/gin"
)

// 1. Function to create payment

type CreatePaymentInput struct {
	OrderId uint `json:"orderId" binding:"required"`
}

func CreatePayment(c *gin.Context) {
	var input CreatePaymentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order ID",
			"status":  false,
		})
		return
	}

	order, err := services.GetOrderById(uint(input.OrderId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
			"status":  false,
		})
		return
	}

	payment, err := services.CreatePayment(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment created successfully",
		"status":  true,
		"data":    payment,
	})
}

// 2. Functiont to verify payment
type VerifyPaymentInput struct {
	RazorpayOrderId   string `json:"razorpayOrderId" binding:"required"`
	RazorpayPaymentId string `json:"razorpayPaymentId" binding:"required"`
	RazorpaySignature string `json:"razorpaySignature" binding:"required"`
}

func VerifyPayment(c *gin.Context) {
	var input VerifyPaymentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"status":  false,
		})
		return
	}

	err := services.VerifyPayment(input.RazorpayOrderId, input.RazorpayPaymentId, input.RazorpaySignature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment verified successfully",
		"status":  true,
	})
}

// 3. Function to handler razorpay webhook
func HandlerRazorpayWebhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Inavalid Request",
			"status":  false,
		})
		return
	}

	signature := c.GetHeader("X-Razorpay-Signature")
	secret := os.Getenv("RAZORPAY_WEBHOOK_SECRET")

	if !verifyWebhookSignature(body, signature, secret) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Signature mismatched",
			"status":  false,
		})
		return
	}

	// Process webhook event
	err = services.ProcessWebhookEvent(body, signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Webhook received",
		"status":  true,
	})
}

func verifyWebhookSignature(body []byte, signature, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	computedSignature := hex.EncodeToString(h.Sum(nil))

	return computedSignature == signature
}
