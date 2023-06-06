# rbook
Utility for giving books (epub, pdf, txt, etc.) a predictable, descriptive file name

## Installation
* Make sure you have `go` and `git` installed
* Clone this repository
  ```sh
  git clone https://github.com/JessebotX/rbook.git --depth 1
  ```
* Run the following commands
  ```sh
  cd rbook
  go build
  ```
* Move executable `rbook` (or `rbook.exe`) to your [PATH](https://katiek2.github.io/path-doc/).

## Usage
Suppose you have a file called "The Odyssey (Homer).epub"

```sh
# FORMAT: rbook r <INPUT_FILE_PATH> <BOOK TITLE> [additional tags]
rbook r "The Odyssey (Homer).epub" "The Odyssey" "Homer"
fantasy mythology greek
```

The result is the epub file being renamed to
`the-odyssey__homer_fantasy_greek_mythology.epub`.
