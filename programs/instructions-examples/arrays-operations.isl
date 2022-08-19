# Arrays operations.

# Let's define our array:
var array [1,2,3,4,5]

# Length
# Usage: LENGTH ARRAY VARIABLE_TO_STORE_RESULT
length array arrayLength
say "The length of $(array) is: $(arrayLength)\n"

# Get
# Usage: GET ARRAY INDEX VARIABLE_TO_STORE_RESULT
get array 0 element
say "The element at position 0 is: $(element)\n"

# Setindex
# Usage: SETINDEX ARRAY INDEX ELEMENT
setindex array 0 10
say "The element at position 0 is now 10: $(array)\n"

# Append
# Usage: APPEND ARRAY ELEMENT
append array 80
say "After appending 80 to the array: $(array)\n"

# Prepend
# Usage: PREPEND ARRAY ELEMENT
prepend array 40
say "After prepending 40 to the array: $(array)\n"

# Removefirst
# Usage: REMOVEFIRST ARRAY VARIABLE_TO_STORE_RESULT
removefirst array element
say "After removing first: $(array), element removed: $(element)\n"

# Removelast
# Usage: REMOVELAST ARRAY VARIABLE_TO_STORE_RESULT
removelast array element
say "After removing last: $(array), element removed: $(element)\n"

# Contains
# Usage: CONTAINS ARRAY ELEMENT VARIABLE_TO_STORE_RESULT
contains array 2 contains1
contains array 999 contains999
# If the array contains the element, VARIABLE_TO_STORE_RESULT will be 1, else 0.
say "The array contains 2? $(contains1)"
say "The array contains 999? $(contains999)"
