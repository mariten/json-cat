#!/bin/sh

# Check input parameters
JSON_FILE="$1"
if [ "$JSON_FILE" == "" ]; then
    echo 'Please specify a valid file name'
    echo 'EXAMPLE:'
    echo '  jsoncat ~/dev/test_out.json'
    exit 1
fi

# Check file exists
if [ ! -f "$JSON_FILE" ]; then
    echo "File '$JSON_FILE' does not exist"
    exit 2
fi

# Pretty print JSON to command line
node -e "console.log(JSON.stringify(JSON.parse(require('fs').readFileSync(process.argv[1])), null, 4));" "$1"
