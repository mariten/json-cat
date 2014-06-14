#!/usr/bin/php
<?php
// Check input parameters
if(count($argv) < 2) {
    echo('Please specify file name' . "\n");
    echo('EXAMPLE:' . "\n");
    echo('  jsoncat ~/dev/test_out.json' . "\n");
    exit(1);
}

// Check file exists
$json_file = $argv[1];
if(file_exists($json_file) == false) {
    echo('Could not find file "' . $json_file . '"' . "\n");
    exit(2);
}

// Pretty print JSON to command line
$json_contents_raw = file_get_contents($json_file);
$json_contents_pretty = json_decode($json_contents_raw, true);
if($json_contents_pretty === null) {
    echo('"' . $json_file . '" contains invalid JSON'. "\n");
} else {
    echo(var_export($json_contents_pretty, true));
    echo("\n");
}

exit(0);
