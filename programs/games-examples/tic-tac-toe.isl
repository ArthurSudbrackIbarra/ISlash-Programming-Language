# Board rows.
set row1 ["0","1","2"]
set row2 ["3","4","5"]
set row3 ["6","7","8"]

# Variables that control the game turns.
set turn 1

# While true:
while 1
    # Printing the board.
    say ""
    say row1
    say "-------"
    say row2
    say "-------"
    say row3
    say ""

    # Odd turn number = Player 1.
    # Even turn number = Player 2.
    mod turn 2 isPlayer1
    if isPlayer1
        set playerName "Player 1"
        set symbol "X"
    else
        set playerName "Player 2"
        set symbol "O"
    endif

    # Finding out the correct row.
    input position "=> $(playerName), choose a position by typing a number: "
    lessequal position 2 inRow1
    lessequal position 5 inRow2
    lessequal position 8 inRow3

    # Replacing the position number with an 'X' or 'O'.
    if inRow1
        # Verifying if the position is free.
        accessindex row1 position element
        notequal element "X" notX
        notequal element "O" notO
        and notX notO isFree
        if isFree
            setindex row1 position symbol
        else
            say "\nThe position is not free."
            decrement turn
        endif
    else
        if inRow2
            # Verifying if the position is free.
            sub position 3 position
            accessindex row2 position element
            notequal element "X" notX
            notequal element "O" notO
            and notX notO isFree
            if isFree
                setindex row2 position symbol
            else
                say "\nThe position is not free."
                decrement turn
            endif
        else
            # Verifying if the position is free.
            sub position 6 position
            accessindex row3 position element
            notequal element "X" notX
            notequal element "O" notO
            and notX notO isFree
            if isFree
                setindex row3 position symbol
            else
                say "\nThe position is not free."
                decrement turn
            endif
        endif
    endif

    #
    # Checking if the game has ended.
    #

    # Horizontal 1.
    equal row1 ["X","X","X"] condition1
    equal row1 ["O","O","O"] condition2
    or condition1 condition2 victory
    if victory
        say "\nVictory!"
        break
    endif

    # Horizontal 2.
    equal row2 ["X","X","X"] condition1
    equal row2 ["O","O","O"] condition2
    or condition1 condition2 victory
    if victory
        say "\nVictory!"
        break
    endif

    # Horizontal 3.
    equal row3 ["X","X","X"] or1
    equal row3 ["O","O","O"] or2
    or or1 or2 victory
    if victory
        say "\nVictory!"
        break
    endif

    # Verticals.
    rangearray 3 indexes
    foreach index indexes
        accessindex row1 index element1
        accessindex row2 index element2
        accessindex row3 index element3
        equal element1 "X" eq1
        equal element2 "X" eq2
        equal element3 "X" eq3
        and eq1 eq2 victory1
        and victory1 eq3 victory1
        equal element1 "O" eq1
        equal element2 "O" eq2
        equal element3 "O" eq3
        and eq1 eq2 victory2
        and victory2 eq3 victory2
        or victory1 victory2 victory
        if victory
            say "\nVictory!"
            break
        endif
    endforeach

    # Diagonal 1 - Right Left.
    accessindex row1 0 element1
    accessindex row2 1 element2
    accessindex row3 2 element3
    equal element1 "X" eq1
    equal element2 "X" eq2
    equal element3 "X" eq3
    and eq1 eq2 victory1
    and victory1 eq3 victory1
    equal element1 "O" eq1
    equal element2 "O" eq2
    equal element3 "O" eq3
    and eq1 eq2 victory2
    and victory2 eq3 victory2
    or victory1 victory2 victory
    if victory
        say "\nVictory!"
        break
    endif

    # Diagonal 2 - Left Right.
    accessindex row1 2 element1
    accessindex row2 1 element2
    accessindex row3 0 element3
    equal element1 "X" eq1
    equal element2 "X" eq2
    equal element3 "X" eq3
    and eq1 eq2 victory1
    and victory1 eq3 victory1
    equal element1 "O" eq1
    equal element2 "O" eq2
    equal element3 "O" eq3
    and eq1 eq2 victory2
    and victory2 eq3 victory2
    or victory1 victory2 victory
    if victory
        say "\nVictory!"
        break
    endif

    # Draw.
    greaterequal turn 9 draw
    if draw
        say "\nIt's a draw!"
        break    
    endif

    # Setup for the next turn.
    set inRow1 0
    set inRow2 0
    set inRow3 0
    increment turn
endwhile

# Printing the board one last time.
say ""
say row1
say "-------"
say row2
say "-------"
say row3
say ""
