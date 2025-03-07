# Simple Blockchain Client

This is a simple blockchain client implemented in Go. It exposes a REST API to interact with the blockchain.

## Features
- `GET /blocks`: Retrieve the entire blockchain.
- `POST /mine`: Mine a new block and add it to the blockchain.

## Running Locally
1. Build the Docker image:
   ```bash
   docker build -t blockchain-client .