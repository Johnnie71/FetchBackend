# Receipt Points Processor

A simple Go application that processes receipts, calculates points based on various criteria, and allows retrieving the calculated points using a unique receipt ID.

## Overview

This application accepts receipt data via a `POST` request, calculates points based on the receipt details (retailer name, purchase date, purchase time, items, and total), and returns a unique receipt ID. You can then use the receipt ID to fetch the total points awarded via a `GET` request.

## Technologies
- **Language**: GO v1.23.3
- **Framework**: GIN v1.10.0
- **Hotreload**: AIR v1.61.1 

## Features

- **Process receipts**: Calculate points based on receipt details.
- **Generate unique ID**: Each receipt receives a unique ID upon processing.
- **Retrieve points**: Query the total points awarded for a specific receipt by ID.

## API Endpoints

### `POST /receipts/process`

Processes a receipt and returns the id of the receipt.

#### Request Body (JSON)
```json
{
  "retailer": "SuperMart",
  "purchaseDate": "2024-11-20",
  "purchaseTime": "15:30",
  "items": [
    {
      "shortDescription": "Apple",
      "price": "1.50"
    },
    {
      "shortDescription": "Banana",
      "price": "0.75"
    }
  ],
  "total": "3.25"
}

### `GET /receipts/id/points`
Gets the points associated with the receipts ID

## Setup Application

To get the application up and running locally, follow these steps:

### Prerequisites
Ensure you have the following installed:
- **Go** (v1.23.3) – The programming language used for the application
- **Gin Framework** (v1.10.0) – A web framework for Go used in this application
- **AIR** (v1.61.1) – Hot reload tool for Go to speed up development

### Steps to Set Up

1. **Clone the repository**:
    ```bash
    git clone git@github.com:Johnnie71/FetchBackend.git GoFetch
    cd GoFetch
    ```

2. **Install dependencies**:
    This project uses Go modules to manage dependencies. Run the following command to install all dependencies:
    ```bash
    go mod tidy
    ```

3. **Run the application**:
    To start the application in development mode with hot reloading, use **AIR**:
    ```bash
    air
    ```
    If **AIR** is not installed, you can run the application without hot reloading by running:
    ```bash
    go run main.go
    ```
    The app will now be running at `http://localhost:8080`.

4. **Access the API**:
    - **POST** `/receipts/process`: To process a receipt and get an ID assigned to it.
    - **GET** `/receipts/:id/points`: To get the points awarded for a specific receipt by its ID.
