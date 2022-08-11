# Modify 'n' to the value you wish. 7! in this example.
set n 7

# This will be our counter.
# counter = n
set counter n

# Variable to store the multiplication of the numbers.
set result 1

# While counter > 0
while counter
    # result = result * counter
    mult result counter result
    # counter = counter - 1
    decrement counter
endwhile

say "The factorial of $(n) is: $(result)."
