# HTML Link Parser

Create a package that makes it easy to parse an HTML file and extract all of the links (```<a href="">...</a>``` tags). For each extracted link you should return a data structure that includes both the href, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Implementation of HTML Link Parser from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Commands:**

- go run examples\ex1\main.go
- go run examples\ex2\main.go
- go run examples\ex3\main.go
- go run examples\ex4\main.go

**Features:**

- using depth first search to traverse html document
- extracting relevant information from html document

**Packages explored:**

- fmt
- os - to open and close the files
- io - to read from file which satisfies io.Reader interface
- [golang.org/x/net/html](golang.org/x/net/html) - to parse the html document into Tree structure
- strings - to format relevant data

**Output:**

``` terminal
D:\gophercises\link-parser>go run examples\ex1\main.go
/other-page                                       : A link to another page some span
/page-two                                         : A link to second page

D:\gophercises\link-parser>go run examples\ex2\main.go
https://www.twitter.com/joncalhoun                : Check me out on twitter
https://github.com/gophercises                    : Gophercises is on Github!

D:\gophercises\link-parser>go run examples\ex3\main.go
#                                                 : Login
/lost                                             : Lost? Need help?
https://twitter.com/marcusolsson                  : @marcusolsson

D:\gophercises\link-parser>go run examples\ex4\main.go
/dog-cat                                          : dog cat
```
