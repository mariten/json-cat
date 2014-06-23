What is This?
=============
Utilities for printing JSON to the Linux command line (like the `cat` command) in a human-readable format

Available in multiple languages (PHP and Node.js for now), making use of any built-in JSON processing capabilities of a given language

Download
========
Fetch any [**jsoncat**](https://github.com/mariten/json-cat/blob/master/bin/) from the master branch `bin` directory of this repository, located in the `bin` directory.  It is recommended to rename the downloaded version of your choice to `jsoncat`.

### Download Node.js Version
```bash
cd ~/bin
wget --no-check-certificate -O jsoncat \
'https://raw.githubusercontent.com/mariten/json-cat/master/bin/jsoncat-node'
```

### Download PHP Version
```bash
cd ~/bin
wget --no-check-certificate -O jsoncat \
'https://raw.githubusercontent.com/mariten/json-cat/master/bin/jsoncat-php'
```

Install
=======
Installing it is as simple as setting it to be executable
```bash
chmod +x ~/bin/jsoncat
```

Usage
=====
Assuming that your `~/bin` folder is set in your PATH environment variable:

```bash
cd ~/data/some_directory_with_json_files/
jsoncat JSON_FILENAME
```

Examples
========
### Without `jsoncat` (Ugly)
```
cat ~/tmp/test.json
-------------------
{"id":105,"name":"Cat bookend","price":29.99,"tags":["home","pets","cats"]}
```

### With `jsoncat` (Pretty)
```
jsoncat ~/tmp/test.json
-------------------
{
    "id": 105,
    "name": "Cat bookend",
    "price": 29.99,
    "tags": [
        "home",
        "pets",
        "cats"
    ]
}
```
