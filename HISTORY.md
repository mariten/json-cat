1.3.0 (2015-01-09)
====================
* Added new version using GoLang (`golang-jsoncat.go`)

1.2.0 (2014-10-23)
====================
* Added new version using Node.js (`python-jsoncat.py`)
* Added download instructions for Python version to README
* Slight revamp of README, including a usage example using `grep` as well as a picture of a **JSON cat** for good measure

1.1.0 (2014-06-23)
====================
* Added new version using Node.js (`node-jsoncat.sh`)
* To handle multiple language versions of JSON pretty-printing, renamed existing PHP version from `jsoncat` to `php-jsoncat.php`
* Rewrote README to reflect that same logic exists for multiple programming languages (multiple download options)

1.0.0 (2014-06-11)
====================
* Uploaded to GitHub
* Added this HISTORY file and README instructions

0.2.0 (2014-02-06)
====================
* Show warning message if JSON string found in file is invalid
* Cleaned up output, added an "Example" help-like message
* Renamed `jcat` (PHP JSON printer) to `jsoncat`

0.1.0 (2014-02-04)
====================
* Initial version of `jsoncat`, originally named `jcat`
* Uses PHP build-in functions `json_decode` to turn a JSON string (found in a file) in a PHP array, then prints that array to stdout using `var_export`
