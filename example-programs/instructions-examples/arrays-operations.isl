# Arrays operations.

# Let's define our array:
var array [1,2,3,4,5]

# LENGTH
# Usage: length ARRAY VARIABLE_TO_STORE_RESULT
length array arrayLength
say "The length of $(array) is: $(arrayLength)\n"

# GET
# Usage: get ARRAY INDEX VARIABLE_TO_STORE_RESULT
get array 0 element
say "The element at position 0 is: $(element)\n"

# SETINDEX
# Usage: setindex ARRAY INDEX ELEMENT
setindex array 0 10
say "The element at position 0 is now 10: $(array)\n"

# APPEND
# Usage: append ARRAY ELEMENT
append array 80
say "After appending 80 to the array: $(array)\n"

# PREPEND
# Usage: prepend ARRAY ELEMENT
prepend array 40
say "After prepending 40 to the array: $(array)\n"

# REMOVEFIRST
# Usage: removefirst ARRAY VARIABLE_TO_STORE_RESULT
removefirst array element
say "After removing first: $(array), element removed: $(element)\n"

# REMOVELAST
# Usage: removelast ARRAY VARIABLE_TO_STORE_RESULT
removelast array element
say "After removing last: $(array), element removed: $(element)\n"

# SWAP
# Usage: swap ARRAY INDEX_1 INDEX_2
swap array 0 1
say "After swapping positions 0 and 1: $(array)\n"

# CONTAINS
# Usage: contains ARRAY ELEMENT VARIABLE_TO_STORE_RESULT
contains array 2 contains1
contains array 999 contains999
# If the array contains the element, VARIABLE_TO_STORE_RESULT will be 1, else 0.
say "The array contains 2? $(contains1)"
say "The array contains 999? $(contains999)"
