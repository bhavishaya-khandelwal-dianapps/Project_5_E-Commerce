package config

import (
	"log"
	"os"

	"github.com/razorpay/razorpay-go"
)

var RazorpayClient *razorpay.Client

// Initialize razorpay client
func InitRazorpay() {
	key := os.Getenv("RAZORPAY_KEY_ID")
	secret := os.Getenv("RAZORPAY_KEY_SECRET")

	if key == "" || secret == "" {
		log.Fatal("❌ Razorpay environment variables not found")
	}

	RazorpayClient = razorpay.NewClient(key, secret)
	log.Println("✅ Razorpay client initialized")
}
