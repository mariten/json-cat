package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // Check input parameters
    argv := os.Args
    argc := len(argv)
    if argc < 2 {
        fmt.Println("Please specify file name")
        fmt.Println("EXAMPLE:")
        fmt.Println("  jsoncat ~/dev/test_out.json")
        os.Exit(1)
    }
    target_file_name := argv[1]

    // Read file content (and errors that occurred)
    file_content, read_error := ioutil.ReadFile(target_file_name)
    if read_error != nil {
        fmt.Println("Could not find file \"" + target_file_name + "\": ", read_error)
        os.Exit(2)
    }

    // Convert to JSON object
    var json_content map[string]interface{}
    json_decode_error := json.Unmarshal(file_content, &json_content)
    if json_decode_error != nil {
        fmt.Println("\"" + target_file_name + "\" contains invalid JSON: ", json_decode_error)
        os.Exit(3)
    }

    // Pretty-print
    pretty_json, prettify_error := json.MarshalIndent(json_content, "", "    ")
    if prettify_error != nil {
        fmt.Println("Failed converting JSON from \"" + target_file_name + "\" to pretty format: ", prettify_error)
        os.Exit(4)
    }
    os.Stdout.Write(pretty_json)
}
