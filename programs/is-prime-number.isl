# Modify 'number' to the value you wish. 31 in this example.
declare number 31

# Variable to store how many dividers the number has.
declare dividersCount 1

# Variable to store the dividers that will be tested.
declare divider 2

# This will be our counter.
# counter = number
declare counter number

# While counter > 0
while counter
    # remainer = number % divider
    mod number divider remainer
    # remainerIsZero = !remainer
    not remainer remainerIsZero
    if remainerIsZero
        # dividersCount = dividersCount + 1
        increment dividersCount
    endif
    # divider = divider + 1
    increment divider
    # counter = counter - 1
    decrement counter
endwhile

# isPrime = (dividersCount == 2) ?
equal dividersCount 2 isPrime

if isPrime
    say "The number $(number) is prime."
else
    say "The number $(number) is not prime."
endif
