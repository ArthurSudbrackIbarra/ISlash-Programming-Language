# Comparison operators.

# Usage:
# INSTRUCTION X Y VARIABLE_TO_STORE_RESULT

# The result stored in 'VARIABLE_TO_STORe_RESULT' will be 1 if the comparison is true and 0 if false.

# Let's define the numbers to compare:
var num1 10
var num2 20

# Greater than (>)
greater num1 num2 result
say "$(num1) is greater than $(num2)? $(result)\n"

# Greater than or equal (>=)
greaterequal num1 num2 result
say "$(num1) is greater than or equal $(num2)? $(result)\n"

# Less than (<)
less num1 num2 result
say "$(num1) is less than $(num2)? $(result)\n"

# Less than or equal (<=)
lessequal num1 num2 result
say "$(num1) is less than or equal $(num2)? $(result)\n"

# Equal (==)
# The equal instruction can be used to compare not only numbers, but also strings and arrays.
equal num1 num2 result1
equal "Hello" "Hello" result2
equal [1,2,3] [1,2,3,4] result3
say "$(num1) is equal to $(num2)? $(result1)"
say "'Hello' is equal to 'Hello'? $(result2)"
say "[1,2,3] is equal to [1,2,3,4]? $(result3)\n"

# Not equal (!=)
# The notequal instruction can be used to compare not only numbers, but also strings and arrays.
notequal num1 num2 result1
notequal "Hello" "Hello" result2
notequal [1,2,3] [1,2,3,4] result3
say "$(num1) is not equal to $(num2)? $(result1)"
say "'Hello' is not equal to 'Hello'? $(result2)"
say "[1,2,3] is not equal to [1,2,3,4]? $(result3)"
