# Input-Output.

# Get user input (INPUT):
# Usage: input VARIABLE_TO_STORE_VALUE PROMPT
input x "Type a value for x: "
say "X value: $(x)\n"

# Notice: Numbers typed will turn into number variables and other things will turn into string variables.

# Print to the terminal (OUTPUT):
# Usage: say X

# Let's declare some variables:
var num 1
var str "Hi"
var array [1,2,3]

# You can print everything to the terminal:
say num
say str
say array
# You can combine values using string interpolation:
say "This is a number: $(num), but this is an array: $(array)"
