package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/stripe/stripe-go/v74"
    "github.com/stripe/stripe-go/v74/customer"
    "github.com/stripe/stripe-go/v74/invoice"
)

func main() {
    // Ensure the Stripe API key is provided as an environment variable
    apiKey := os.Getenv("STRIPE_API_KEY")
    if apiKey == "" {
        log.Fatal("STRIPE_API_KEY environment variable is not set")
    }

    // Initialize the Stripe client
    stripe.Key = apiKey

    // Fetch all customers
    params := &stripe.CustomerListParams{}
    params.Limit = stripe.Int64(100)
    i := customer.List(params)

    for i.Next() {
        c := i.Customer()

        // Check if the customer has any successful payments for invoices
        hasSuccessfulPayment := false
        invoiceParams := &stripe.InvoiceListParams{Customer: stripe.String(c.ID)}
        invoiceList := invoice.List(invoiceParams)

        for invoiceList.Next() {
            inv := invoiceList.Invoice()
            if inv.Paid {
                hasSuccessfulPayment = true
                break
            }
        }

        if !hasSuccessfulPayment {
            // Delete the customer
            _, err := customer.Del(c.ID, nil)
            if err != nil {
                log.Printf("Failed to delete customer %s: %v\n", c.ID, err)
            } else {
                fmt.Printf("Deleted customer %s\n", c.ID)
            }
        }
    }

    if err := i.Err(); err != nil {
        log.Fatalf("Error fetching customers: %v\n", err)
    }

    fmt.Println("Customer cleanup complete.")
}
