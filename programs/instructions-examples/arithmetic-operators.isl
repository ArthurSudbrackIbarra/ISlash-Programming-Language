# Arithmetic Operators.

# Usage:
# INSTRUCTION NUMBER_1 NUMBER_2 VARIABLE_TO_STORE_THE_RESULT

# Let's define our operands:
var num1 3
var num2 5

# Addition (+)
# 3 + 5 = 8
add num1 num2 result
say "$(num1) + $(num2) = $(result)\n"

# Subtraction (-)
# 3 - 5 = -2
sub num1 num2 result
say "$(num1) - $(num2) = $(result)\n"

# Multiplication (*)
# 3 * 5 = 15
mult num1 num2 result
say "$(num1) * $(num2) = $(result)\n"

# Division (/)
# 3 / 5 = 0.6
div num1 num2 result
say "$(num1) / $(num2) = $(result)\n"

# Exponentiation
# 3 ^ 5 = 243
power num1 num2 result
say "$(num1) ^ $(num2) = $(result)\n"

# Nth root
# Square root of 49 = 7
root 49 2 result
say "Square root of 49: $(result)\n"
# Cubic root of 125 = 5
root 125 3 result
say "Cubic root of 125: $(result)"
