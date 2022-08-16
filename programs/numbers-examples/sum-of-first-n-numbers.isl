# Modify 'n' to the value you wish.
var n 20

# This will be our counter.
# counter = n
var counter n

# Variable to store the sum of the numbers.
var sum 0

# While counter > 0
while counter
    # sum = sum + counter
    add sum counter sum
    # counter = counter - 1
    decrement counter
endwhile

say "The sum of the first $(n) integer numbers is: $(sum)."
