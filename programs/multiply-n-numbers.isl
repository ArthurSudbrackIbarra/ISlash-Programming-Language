# Declaring empty number array.
set numbers []number

# Getting user input.
input limit "How many numbers do you want to multiply? "
input multiplier "Multiply by which number? "

# While loop.
while limit
    input number "Type a number: "
    append numbers number
    decrement limit
endwhile

# Foreach loop.
foreach element numbers
    mult element multiplier result
    say "$(element) x $(multiplier) = $(result)"
endforeach
