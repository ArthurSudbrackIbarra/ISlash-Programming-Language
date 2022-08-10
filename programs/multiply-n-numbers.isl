say "---------------------------------------------"

declare numbers []number

input limit "How many numbers do you want to multiply? "
input multiplier "Multiply by which number? "

say "---------------------------------------------"

while limit
    input number "Type a number: "
    append numbers number
    decrement limit
endwhile

say "---------------------------------------------"

foreach element numbers
    mult element multiplier result
    say "$(element) x $(multiplier) = $(result)"
endforeach
