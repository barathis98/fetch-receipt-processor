# Receipt Processor API

This project implements a web service that processes receipts and calculates points based on specific rules. The API allows users to submit receipts and retrieve the calculated points for each receipt.

## Table of Contents

- [Receipt Processor API](#receipt-processor-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Installing Go](#installing-go)
    - [Installing the Project](#installing-the-project)
  - [Setup](#setup)
  - [Running the Application](#running-the-application)
  - [API Endpoints](#api-endpoints)
    - [1. Process Receipt](#1-process-receipt)
    - [2. Get Points](#2-get-points)
  - [Troubleshooting](#troubleshooting)

## Installation

### Installing Go

Before installing the project, you need to have Go installed on your system. Follow these steps to install Go:

1. Visit the official Go download page: https://golang.org/dl/

2. Download the installer for your operating system:
   - For Windows: Download the MSI installer
   - For macOS: Download the package installer
   - For Linux: Follow the instructions for your specific distribution

3. Run the installer and follow the prompts. The default installation options are usually sufficient.

4. After installation, open a new terminal or command prompt and verify the installation by running:
   ```
   go version
   ```
   You should see the installed Go version printed.

5. Set up your Go workspace:
   - For Go 1.13 and later (recommended), you don't need to set GOPATH explicitly.
   - For earlier versions, set the GOPATH environment variable to your desired workspace location.

### Installing the Project

Now that Go is installed, you can proceed with installing the project:

1. Install Git if you haven't already (https://git-scm.com/downloads)

2. Clone the repository:
   ```
   git clone https://github.com/barathis98/fetch-receipt-processor.git
   cd fetch-recept-processor/receipt-processor
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

## Setup

No additional setup is required for this application as it uses in-memory storage. However, make sure you have the necessary permissions to run a web server on your machine.

## Running the Application

This project uses a Makefile to simplify common tasks. To start the server, run the following command in the project root directory:

```
make run
```

The server will start and listen on `http://localhost:8080`.

If you want to see other available make commands, you can run:

```
make help
```

This will display a list of all available make commands with their descriptions.

## API Endpoints

### 1. Process Receipt

- **URL:** `/receipts/process`
- **Method:** POST
- **Body:** JSON Receipt object
- **Response:** JSON containing an ID for the receipt

Example request:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    }
  ],
  "total": "18.74"
}
```

Example response:
```json
{
  "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

### 2. Get Points

- **URL:** `/receipts/{id}/points`
- **Method:** GET
- **Response:** JSON object containing the number of points awarded

Example response:
```json
{
  "points": 32
}
```


## Troubleshooting

If you encounter any issues while setting up or running the application, try the following:

1. Ensure you have the correct version of Go installed by running `go version`.
2. Check that all dependencies are properly installed by running `go mod tidy`.
3. If you get "make: *** No rule to make target `run'.  Stop."  error. Make sure you cd (cd receipt-processor) into the directory properly.
4. Make sure no other application is using port 8080.
5. If you're having trouble with Go commands, make sure your PATH includes the Go binary directory.
6. If make commands are not working, ensure you have Make installed on your system.

For any other issues, please open an issue on the GitHub repository.