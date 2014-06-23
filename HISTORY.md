1.1.0 (2014-06-23)
====================
* Added new version using Node.js (`jsoncat-node`)
* To handle multiple language versions of JSON pretty-printing, renamed existing PHP version from `jsoncat` to `jsoncat-php`
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
==================
* Initial version of `jsoncat`, originally named `jcat`
* Uses PHP build-in functions `json_decode` to turn a JSON string (found in a file) in a PHP array, then prints that array to stdout using `var_export`
