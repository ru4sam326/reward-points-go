# PROBLEM

A retailer offers a rewards program to its customers, awarding points based on each recorded purchase.

A customer receives 2 points for every dollar spent over $100 in each transaction, plus 1 point for every dollar spent over $50 in each transaction

(e.g. a $120 purchase = 2x$20 + 1x$50 = 90 points).

Given a record of every transaction during a three-month period, calculate the reward points earned for each customer per month and total.

Make up a data set to best demonstrate your solution
Check solution into GitHub

## ABOUT CODING

Written a separate package for movies and reviews with the structure, Written an utility for reading the content from the files.

Tweet functionality is independent, so that we can use it without reading it from file as well. written utilities for deriving the rating and movie title based on with and without trim flags

## Running the application

Clone the project into the project directory

Run the below to get the dependencies

`go mod tidy`

Pass the purchase history data in a  json file with the below format for each customer

- customer_name         string          Name of the customer
- transaction_date      timestamp       Transaction time in ISO8601 / RFC3339 format
- transaction_amount    float64         Transaction amount

To run the application pass the purchase_details file path

`go run main.go purchase_details.json`

## Running all Test Cases

go test -v ./...

## Test Coverage Command line

go test -v ./... -cover

## Specific test

go clean -testcache && go test -run TestGetTweetWithNoMovie -v gitlab.com/99designs-coding-test/tweet