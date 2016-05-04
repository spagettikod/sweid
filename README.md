# sweid
Go API to validate Swedish personal identification (personnummer) and
organisational (organisationsnummer) numbers. Specifications for these
identifiers can be found at the Swedish Tax Agency, http://www.skatteverket.se.

## Usage
Call one of the two methods to test a Personal Identification Number (ValidPN) or Organisation Number (ValidON).

Valid formats for Personal Identification Numbers:
* YYYYMMDD-NNNN
* YYMMDD-NNNN
* YYYYMMDDNNNN
* YYMMDDNNNN

Valid formats for Organisation Numbers:
* NNNNNN-NNNN
* NNNNNNNNNN

## Example
```
package main

import (
  "fmt"
  "github.com/spagettikod/sweid"
)

func main() {
  if ValidPN("640823-3234") {
		fmt.Println("Correct Personal Identification Number!")
	} else {
		fmt.Println("Invalid Personal Identification Number!")
	}
}
```
