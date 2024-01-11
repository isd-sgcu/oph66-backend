#!/bin/bash

## Usage:
##   ./tools/gen-staff-token.sh <faculty> <department> <faculty-wide>
##   ./tools/gen-staff-token.sh 23 0 true

if [ -f .env ]; then
    source .env
else
    echo "Error: .env file not found!"
    exit 1
fi

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <faculty> <department> <faculty-wide>"
    exit 1
fi

faculty=$1
department=$2
faculty_wide=$3

if [ -z "$faculty" ] || [ -z "$department" ]; then
    echo "Invalid input. Faculty and Department cannot be empty."
    exit 1
fi

if [ -z "$JWT_SECRET_KEY" ]; then
    echo "Error: JWT_SECRET_KEY is not set in .env file!"
    exit 1
fi

SECRET_KEY="$JWT_SECRET_KEY"

HEADER="{\"alg\":\"HS256\",\"typ\":\"JWT\"}"

PAYLOAD="{\"role\":\"staff\",\"faculty\":$faculty,\"department\":$department,\"faculty-wide\":$faculty_wide}"

HEADER_ENCODED=$(echo -n "$HEADER" | base64 | tr -d '=' | tr '/+' '_-')
PAYLOAD_ENCODED=$(echo -n "$PAYLOAD" | base64 | tr -d '=' | tr '/+' '_-')

TOKEN="$HEADER_ENCODED.$PAYLOAD_ENCODED"

SIGNATURE=$(echo -n "$TOKEN" | openssl dgst -sha256 -hmac "$SECRET_KEY" -binary | base64 | tr -d '=' | tr '/+' '_-')

echo "Token: $TOKEN.$SIGNATURE"
