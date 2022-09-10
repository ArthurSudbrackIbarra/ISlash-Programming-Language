#!/bin/bash

# Color.
PURPLE='\033[0;94m'
RESET='\033[0m'

# Programs that won't be tested because they require user input.
NO_TEST=(
  "example-programs/arrays-examples/fibbonaci.isl"
  "example-programs/arrays-examples/multiply-n-numbers.isl"
  "example-programs/arrays-examples/simulating-a-map.isl"
  "example-programs/arrays-examples/sum-of-inputed-numbers.isl"
  "example-programs/arrays-examples/multiply-n-numbers.isl"
  "example-programs/files-examples/write-to-file.isl"
  "example-programs/games-examples/hangman.isl"
  "example-programs/games-examples/tic-tac-toe.isl"
  "example-programs/instructions-examples/input-output.isl"
  "example-programs/numbers-examples/calculator.isl"
  "example-programs/numbers-examples/fizz-buzz.isl"
  "example-programs/numbers-examples/guess-the-number.isl"
  "example-programs/strings-examples/is-palindrome.isl"
  "example-programs/strings-examples/contains.isl"
  "example-programs/strings-examples/spelling.isl"
)

echo

# Recursively iterating through .isl files inside the '../../programs' directory.
for FILE in programs/**/*.isl
do
  if [[ " ${NO_TEST[*]} " =~ " ${FILE} " ]]; then
    continue
  fi
  echo -e "${PURPLE}======= [NOW RUNNING] $FILE =======${RESET}"
  echo
  # Running the ISlash program.
  ./islash $FILE
  # Checking for errors.
  EXIT_CODE=$?
  if [ $EXIT_CODE -ne 0 ]; then
    exit $EXIT_CODE
  fi
  echo
done
