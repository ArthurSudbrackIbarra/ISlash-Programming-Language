# Modify 'number' to the value you wish. 10 in this example.
var number 10

# remainer = number % 2
mod number 2 remainer

if remainer
    # Odd
    say "The number $(number) is odd."
else
    # Even
    say "The number $(number) is even."
endif
