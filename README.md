# galaxy_merchant_trading

This is a program to convert intergalactic transactions queries into human readable format.

## Pre-requisites

This program requires Golang to be installed. [Refrence](https://go.dev/doc/install)
to check if the installation is succesfull use the following command and it should show the version of golang installed.

```
go version
```

## Approach

This code is written in Golang 1.19.
The input is taken from `input.txt` file which is supplied as a the command line argument.
to run the program use the following command

```
go run main.go ./input.txt
```

The lines are scanned for `?` to determine if it is an instruction or a question. If it is an instruction the line is sent to `internal/parse.go` for processing. if it is a question it is sent to `internal/print.go` for processing and the result is sent back to `main.go` file.

The parsing happens based on text recognition and parsing to get the values of the materials and values.
There are 3 structs which store the values after parsing which are

- transactionToNumber
- metalValue
- materialToRoman

These structs store the values which will be later used to process the questions.

The questions are first parsed based on the question type and the materials and metals are identified from it.
Then the structs are used to obtain the value of the text and the output is displayed with the answer.

The design followed is a typical golang structure with `main.go` on the root.
The bussiness logic used is part of `internal` folder
The helper functions are placed in `utils` folder.

The code uses structs to define the maps which are used to answer the questions.

The go.mod go.sum are package manager which handle the dependencies.

## Testing

The test files contain test cases for respective files of parse.go and print.go.
The test structure is typical go structure where the input and outputs of functions are defined. The expected output or expected erros are also defined which is used to asses the test cases.
