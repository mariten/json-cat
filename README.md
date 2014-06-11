What is This?
=============
A utility for printing JSON to the Linux command line (like the `cat` command) in a human-readable format

It is written in PHP and takes advantage of a few of its convenient built-in functions like `json_decode`

Download
========
Fetch [**jsoncat**](https://github.com/mariten/json-cat/blob/master/bin/jsoncat) from the master branch of this repository, located in the `bin` directory:

### Download from *nix Command Line
```bash
cd ~/bin
wget --no-check-certificate 'https://raw.githubusercontent.com/mariten/json-cat/master/bin/jsoncat'
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
### Command
```bash
jsoncat ~/tmp/bookend.json
```

### Printed Result
```
array (
  'id' => 105,
  'name' => 'Cat bookend',
  'price' => 29.99,
  'tags' =>
  array (
    0 => 'home',
    1 => 'pets',
    2 => 'cats',
  ),
)
```

### Original JSON Content
```
{"id":105,"name":"Cat bookend","price":29.99,"tags":["home","pets","cats"]}
```
