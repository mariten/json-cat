#!/usr/bin/env python

import sys
import os.path
import json

# Check input parameters
if len(sys.argv) < 2:
    print 'Please specify file name'
    print 'EXAMPLE:'
    print '  jsoncat ~/dev/test_out.json'
    sys.exit(1)

# Check file exists
target_file_name = sys.argv[1]
does_file_exist = os.path.isfile(target_file_name)
if does_file_exist == False:
    print 'Could not find file "' + target_file_name + '"'
    sys.exit(2)

# Pretty print JSON to command line
try:
    with open(target_file_name, 'r') as file_handle:
        parsed_json = json.load(file_handle)
        print json.dumps(parsed_json, indent=4, sort_keys=True)

except ValueError:
    print '"' + target_file_name + '" contains invalid JSON'
