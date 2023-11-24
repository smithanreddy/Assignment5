# Crypto Price Server

This Go application provides a simple server that fetches and converts cryptocurrency prices. It uses the CoinCap API to retrieve the latest cryptocurrency data.

## Getting Started

Follow the steps below to set up and run the Crypto Price Server on your local machine.

### Prerequisites

Make sure you have Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/smithanreddy/Assignment5


Open your web browser and navigate to http://localhost:8080/price to get cryptocurrency prices in CAD.

API Endpoints
GET /price
This endpoint returns cryptocurrency prices in CAD for Bitcoin, Ethereum, and Tether.

Example Response:

json
{
  "crypto": "Bitcoin",
  "priceCad": "66645.00"
}
