# PROBLEM

A retailer offers a rewards program to its customers, awarding points based on each recorded purchase.

A customer receives 2 points for every dollar spent over $100 in each transaction, plus 1 point for every dollar spent over $50 in each transaction

(e.g. a $120 purchase = 2x$20 + 1x$50 = 90 points).

Given a record of every transaction during a three-month period, calculate the reward points earned for each customer per month and total.

Make up a data set to best demonstrate your solution
Check solution into GitHub

## ABOUT CODING

Prepared a test data for customer purchase details in the `purchase_details.json`, Provide the sample/customized data as a file.
Written a  separate package for fetching the purchase details from file and for calculating the monthly wise reward points.

Added test cases as well and The coverage report: 100.0% of statements

## Running the application

Clone the project into the project directory

Run the below to get the dependencies

`go mod tidy`

Pass the purchase history data in a  json file with the below format for each customer

- customer_name         string          Name of the customer
- transaction_date      timestamp       Transaction time in ISO8601 / RFC3339 format
- transaction_amount    float64         Transaction amount

To run the application with sample the purchase_details file path

`go run main.go`

To run the application with custom the purchase_details file path

`go run main.go purchase_details.json`

## OUTPUT

Sample Output will be in the below format

Reward point Summary Of Customer:John
Feb'2022 : 500
Mar'2022 : 1400
Apr'2022 : 1050
Total Reward points: 2950
Reward point Summary Of Customer:Diana
Feb'2022 : 152
Mar'2022 : 354
Apr'2022 : 1906
Total Reward points: 2412

## Running all Test Cases

go test -v ./...

## Test Coverage Command line

go test -v ./... -cover
