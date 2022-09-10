input limit "Draw a number from 0 to: "
input tries "How many tries do you want? "

random 0 limit drawnNumber
var victory 0

while tries
    input number "Make a guess: "
    equal number drawnNumber correctGuess
    if correctGuess
        var victory 1
        var tries 0
    endif
    decrement tries
endwhile

if victory
    say "You guessed the number!"
else
    say "Too bad, you didn't guess the number: $(drawnNumber)."
endif
