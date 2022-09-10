# Loop blocks.

# WHILE 
# Usage:
# while CONDITION
#   Your logic goes here...
# endwhile

# The 'CONDITION' argument must be a number or a number variable.
#
#   numbers <= 0 FALSE
#   number > 1 TRUE

# Let's declare a variable:
var x 10

# 10 iterations will happen:
# Once x reaches 0, we will not enter the while block anymore.
while x
    say "X value: $(x)"
    decrement x
endwhile

# It is possible to simulate a while true using 'while 1'.
# It is also possible to exit out from while blocks using the 'break' instruction.
while 1
    say "\nThis will only be printed 1 time.\n"
    break
endwhile

# FOREACH
# Foreach is used to iterate over array elements.
# Usage:
# foreach ELEMENT ARRAY
#   Your logic goes here...
# endforeach

# Let's declare a variable:
var array [1,2,3]

foreach element array
    # Element will be 1, then 2, then 3.
    say "Element: $(element)"
endforeach

# Notice: The 'break' instruction does not work with foreach blocks.
# Notice: If you have nested foreach blocks, give different variable names to the elements.
# Example:
# foreach ELEMENT_1 MY_ARRAY
#   foreach ELEMENT_2 MY_OTHER_ARRAY
#       ...
#   endforeach
# endforeach
