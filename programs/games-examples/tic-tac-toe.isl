# Board rows.
set row1 ["0","1","2"]
set row2 ["3","4","5"]
set row3 ["6","7","8"]

# Variables that control the game flow.
set continue 1
set turn 1

while continue
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
        # TEM ERRO DE ELSE AQUI, PORQUE O IF PROCURA O PRÓXIMO ELSE AO INVÉS DE IR PRO DA LINHA 56.
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

    # Checking if one of the players has won the game.

    # Row 1 - Horizontal

    # Checking if it's a draw.
    greaterequal turn 9 draw
    if draw
        say "Game over, it's a draw!"
        set continue 0     
    endif


    # Setup for the next turn.
    set inRow1 0
    set inRow2 0
    set inRow3 0
    increment turn
endwhile


