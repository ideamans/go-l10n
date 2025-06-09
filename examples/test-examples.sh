#!/bin/bash

# Test script for go-l10n examples
set -e

echo "Testing go-l10n examples..."
echo "=========================="
echo

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test basic example
echo -e "${YELLOW}Testing basic example...${NC}"
cd basic

echo "1. Testing with default language detection:"
go run main.go | head -10
echo

echo "2. Testing with LANG=ja_JP.UTF-8:"
LANG=ja_JP.UTF-8 go run main.go | head -10
echo

echo "3. Testing with LANG=en_US.UTF-8:"
LANG=en_US.UTF-8 go run main.go | head -10
echo

echo "4. Testing with LC_ALL=ja_JP.UTF-8:"
LC_ALL=ja_JP.UTF-8 go run main.go | head -10
echo

echo "5. Testing with LANGUAGE=ja:"
LANGUAGE=ja go run main.go | head -10
echo

echo "6. Testing with L10N_DEFAULT_LANGUAGE=ja:"
L10N_DEFAULT_LANGUAGE=ja go run main.go | head -10
echo

echo "7. Testing with L10N_SKIP_DETECTION=1:"
L10N_SKIP_DETECTION=1 LANG=ja_JP.UTF-8 go run main.go | head -10
echo

cd ..

# Test modular example
echo -e "${YELLOW}Testing modular example...${NC}"
cd modular

echo "1. Testing with default language detection:"
go run . | grep -A 5 "Main Module"
echo

echo "2. Testing with LANG=ja_JP.UTF-8:"
LANG=ja_JP.UTF-8 go run . | grep -A 5 "Main Module"
echo

cd ..

echo -e "${GREEN}All tests completed successfully!${NC}"