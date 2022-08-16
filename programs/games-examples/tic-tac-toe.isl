# Board rows.
set row1 ["0","1","2"]
set row2 ["3","4","5"]
set row3 ["6","7","8"]

# Variable that controls the game turns.
set turn 1

# [ Constants ]

set PLAYER_1_NAME "Player 1"
set PLAYER_1_SYMBOL "X"

set PLAYER_2_NAME "Player 2"
set PLAYER_2_SYMBOL "O"

set VICTORY_MESSAGE "\nVictory!"
set DRAW_MESSAGE "\nIt's a draw!"

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
        set playerName PLAYER_1_NAME
        set symbol PLAYER_1_SYMBOL
    else
        set playerName PLAYER_1_NAME
        set symbol PLAYER_2_SYMBOL
    endif

    # Finding out the correct row.
    input position "=> $(playerName), choose a position by typing a number: "
    lessequal position 2 inRow1
    lessequal position 5 inRow2
    lessequal position 8 inRow3

    # Replacing the position number with the player symbol.
    if inRow1
        # Verifying if the position is free.
        accessindex row1 position element
        notequal element PLAYER_1_SYMBOL notX
        notequal element PLAYER_2_SYMBOL notO
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
            notequal element PLAYER_1_SYMBOL notX
            notequal element PLAYER_2_SYMBOL notO
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
            notequal element PLAYER_1_SYMBOL notX
            notequal element PLAYER_2_SYMBOL notO
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
    equal row1 [symbol,symbol,symbol] victory
    if victory
        say VICTORY_MESSAGE
        break
    endif

    # Horizontal 2.
    equal row2 [symbol,symbol,symbol] victory
    if victory
        say VICTORY_MESSAGE
        break
    endif

    # Horizontal 3.
    equal row3 [symbol,symbol,symbol] victory
    if victory
        say VICTORY_MESSAGE
        break
    endif

    # Verticals.
    rangearray 3 indexes
    foreach index indexes
        accessindex row1 index element1
        accessindex row2 index element2
        accessindex row3 index element3
        equal element1 symbol eq1
        equal element2 symbol eq2
        equal element3 symbol eq3
        and eq1 eq2 victory
        and victory eq3 victory
        if victory
            say VICTORY_MESSAGE
            break
        endif
    endforeach

    # Diagonal 1 - Right Left.
    accessindex row1 0 element1
    accessindex row2 1 element2
    accessindex row3 2 element3
    equal element1 symbol eq1
    equal element2 symbol eq2
    equal element3 symbol eq3
    and eq1 eq2 victory
    and victory eq3 victory
    if victory
        say VICTORY_MESSAGE
        break
    endif

    # Diagonal 2 - Left Right.
    accessindex row1 2 element1
    accessindex row2 1 element2
    accessindex row3 0 element3
    equal element1 symbol eq1
    equal element2 symbol eq2
    equal element3 symbol eq3
    and eq1 eq2 victory
    and victory eq3 victory
    if victory
        say VICTORY_MESSAGE
        break
    endif

    # Draw.
    greaterequal turn 9 draw
    if draw
        say DRAW_MESSAGE
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
