
# Stripe Customer Cleanup CLI

This Go CLI application deletes all customers from a Stripe account who have not had any successful payments for invoices.

## Prerequisites

- Go 1.20 or later
- A Stripe account and API key

## Installation

1. **Clone the repository**:

    ```sh
    git clone https://github.com/altafino/stripe-customer-cleanup-cli.git
    cd stripe-customer-cleanup-cli
    ```

2. **Set up the environment variables**:

    Copy the example environment variables file and edit it to include your Stripe API key:

    ```sh
    cp .env.example .env
    ```

    Open `.env` in a text editor and replace `your_stripe_api_key_here` with your actual Stripe API key:

    ```plaintext
    STRIPE_API_KEY=sk_test_your_actual_stripe_api_key
    ```

3. **Initialize the Go module and install dependencies**:

    Initialize the Go module (if not already initialized):

    ```sh
    go mod init stripe-customer-cleanup-cli
    ```

    Install the necessary dependencies:

    ```sh
    go get github.com/stripe/stripe-go/v74
    go get github.com/joho/godotenv
    ```

## Usage

1. **Run the CLI**:

    ```sh
    go run main.go
    ```

    This will delete all customers who have not had any successful payments for invoices. The output will indicate which customers were deleted.

## Project Structure

```
/your-project-directory
    |-- main.go                  # Main Go script
    |-- go.mod                   # Go module file
    |-- .env.example             # Example environment variables file
    |-- .env                     # Environment variables file (ignored by version control)
    |-- README.md                # This README file
```

## Environment Variables

The application requires the following environment variables:

- `STRIPE_API_KEY`: Your Stripe API key. You can obtain it from the Stripe Dashboard.

## Dependencies

This project uses the following Go packages:

- [Stripe Go](https://github.com/stripe/stripe-go)
- [Godotenv](https://github.com/joho/godotenv)

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Contact

If you have any questions or need further assistance, please contact me at X @gloangch.
