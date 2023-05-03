# Parallel SHA256 HEX dump

This software processes each line of the read file in parallel and outputs the processing results in the same order as the original lines.
The content of the process is a HEX dump of the SHA256 checksum of the row data.

## Requirement

| Language | Version |
| :---: | :---: |
| Golang | 1.20.3 |

## Installation

To use this software, please follow the instructions below to install it.

1. Make sure you have the Go language installed. Download and install the latest version of Go from the official Go website [https://golang.org/](https://golang.org/).
2. Clone this repository.
```bash
$ git clone https://github.com/Lium1126/file-hex-dump.git
```
3. Go to the cloned directory.
```bash
$ cd file-hex-dump
```

## Usage
### Run sample case
```bash
$ make run
```

### Run with source code
```bash
$ go run ./main.go <input file>
```

### Run with built binary
```bash
$ make go-build
$ ./myhexdump <input file>
```

# Note

- This software supports common file formats, but not all file formats. Correct results may not be obtained for certain file formats.
- This program can process a large amount of data, but the processing speed depends on the performance of your system. Processing can take a long time, especially for very large files.
- We recommend that you take appropriate data protection measures, such as backing up your files, when using this software. The author assumes no liability for data loss or damage.

# Author

* Name: Lium1126
* Organization: AICHI INSTITUTE OF TECHNOLOGY
* GitHub: [Lium1126](https://github.com/Lium1126)

