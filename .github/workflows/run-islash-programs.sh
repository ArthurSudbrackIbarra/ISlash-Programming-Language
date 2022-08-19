#!/bin/bash

# Programs that won't be tested because they require user input.
NO_TEST=(
  "programs/arrays-examples/fibbonaci.isl"
  "programs/arrays-examples/multiply-n-numbers.isl"
  "programs/arrays-examples/simulating-a-map.isl"
  "programs/arrays-examples/sum-of-inputed-numbers.isl"
  "programs/arrays-examples/multiply-n-numbers.isl"
  "programs/files-examples/write-to-file.isl"
  "programs/games-examples/hangman.isl"
  "programs/games-examples/tic-tac-toe.isl"
  "programs/numbers-examples/calculator.isl"
  "programs/numbers-examples/guess-the-number.isl"
  "programs/strings-examples/is-palindrome.isl"
  "programs/strings-examples/spelling.isl"
)

# Recursively iterating through .isl files inside the '../../programs' directory.
for FILE in programs/**/*.isl
do
  if [[ " ${NO_TEST[*]} " =~ " ${FILE} " ]]; then
    continue
  fi
  # Running the ISlash program.
  echo "=== NOW RUNNING: $FILE ==="
  ./islash $FILE
  # Checking for errors.
  EXIT_CODE=$?
  if [ $EXIT_CODE -ne 0 ]; then
    exit $EXIT_CODE
  fi
  # New line.
  echo
done
