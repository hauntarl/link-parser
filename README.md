# HTML Link Parser

Create a package that makes it easy to parse an HTML file and extract all of the links (```<a href="">...</a>``` tags). For each extracted link you should return a data structure that includes both the href, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Implementation of HTML Link Parser from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

<!--
**Run Commands:**

- go run main\main.go
- go run main\main.go -h (--help) (to get information regarding flags)
- go run main\main.go --yaml file-name.yaml -json=file-name.json

**Features:**

- grouping packages using go.mod
- using command-line flags
- parsing yaml bytes and files
- parsing json bytes and files
- setting up a basic http server
- redirecting requests using http
- persisting the url mapping in database

**Packages explored:**

- fmt
- net/http - to setup a basic http server and redirect requests
- [gopkg.in/yaml.v2](gopkg.in/yaml.v2) - to work with yaml data
- flag - to get yaml/json file name
- os - to open and close the file
- io - to read from file which satisfies io.Reader interface
- encoding/json - to work with json data
- [github.com/boltdb/bolt](github.com/boltdb/bolt) - to store and retrieve urls for specified path

**Output:**

``` terminal
```
-->
