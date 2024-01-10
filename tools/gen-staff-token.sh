#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
    source .env
else
    echo "Error: .env file not found!"
    exit 1
fi

# Check if faculty and department are provided as command-line arguments
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <faculty> <department>"
    exit 1
fi

# Extract faculty and department from command-line arguments
faculty=$1
department=$2

# Ensure that the faculty and department are not empty
if [ -z "$faculty" ] || [ -z "$department" ]; then
    echo "Invalid input. Faculty and Department cannot be empty."
    exit 1
fi

# Ensure that JWT_SECRET_KEY is set
if [ -z "$JWT_SECRET_KEY" ]; then
    echo "Error: JWT_SECRET_KEY is not set in .env file!"
    exit 1
fi

# Define the secret key for JWT token generation
SECRET_KEY="$JWT_SECRET_KEY"

# Define the header for JWT token (Base64 encoded {"alg":"HS256","typ":"JWT"})
HEADER="{\"alg\":\"HS256\",\"typ\":\"JWT\"}"

# Define the payload for JWT token (Base64 encoded {"role":"staff","faculty":<faculty>,"department":<department>})
PAYLOAD="{\"role\":\"staff\",\"faculty\":$faculty,\"department\":$department}"

# Encode the header and payload using base64
HEADER_ENCODED=$(echo -n "$HEADER" | base64 | tr -d '=' | tr '/+' '_-')
PAYLOAD_ENCODED=$(echo -n "$PAYLOAD" | base64 | tr -d '=' | tr '/+' '_-')

# Generate the JWT token by combining the encoded header and payload with a "." separator
TOKEN="$HEADER_ENCODED.$PAYLOAD_ENCODED"

# Sign the JWT token using the secret key and encode the signature
SIGNATURE=$(echo -n "$TOKEN" | openssl dgst -sha256 -hmac "$SECRET_KEY" -binary | base64 | tr -d '=' | tr '/+' '_-')

# Print the generated JWT token
echo "Token: $TOKEN.$SIGNATURE"
