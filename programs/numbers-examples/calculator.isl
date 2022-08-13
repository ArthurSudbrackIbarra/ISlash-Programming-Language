set continue 1

while continue
    # Menu.
    say "\n=== CALCULATOR ==="
    say "[a] - Add"
    say "[s] - Subtract"
    say "[m] - Multiply"
    say "[d] - Divide"
    say "\n[Other] - Exit\n"

    # Get user option.
    input option "Choose an option: "

    # 
    set isAdd 0
    set isSub 0
    set isMult 0
    set isDiv 0

    # Add
    equal option "a" isAdd
    if isAdd
        input num1 "Type the first number: "
        input num2 "Type the second number: "
        add num1 num2 result
        say "\nThe result is $(result)."
        increment continue
    endif

    # Subtract
    equal option "s" isSub
    if isSub
        input num1 "Type the first number: "
        input num2 "Type the second number: "
        sub num1 num2 result
        say "\nThe result is $(result)."
        increment continue
    endif

    # Multiply
    equal option "m" isMult
    if isMult
        input num1 "Type the first number: "
        input num2 "Type the second number: "
        mult num1 num2 result
        say "\nThe result is $(result)."
        increment continue
    endif

    # Divide
    equal option "d" isDiv
    if isDiv
        input num1 "Type the first number: "
        input num2 "Type the second number: "
        div num1 num2 result
        say "\nThe result is $(result)."
        increment continue
    endif

    # Setup for the next loop iteration.
    decrement continue
endwhile
