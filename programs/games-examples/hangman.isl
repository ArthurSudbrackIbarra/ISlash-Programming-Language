# List of possible words.
var words ["Play","Melon","Whale","Subject","Magic"]

# The user can choose a custom word too.
input customWord "Type a custom word or type 0 to use a predefined word: "
notequal customWord 0 isCustom
if isCustom
    var words [customWord]
endif

# maxRandomRange = length(words) - 1
length words wordsArrayLength
sub wordsArrayLength 1 maxRandomRange

# randomNumber = random from 0 to maxRandomRange (inclusive).
random 0 maxRandomRange randomNumber

# randomWord = uppercase(words[randomNumber])
get words randomNumber randomWord
upper randomWord randomWord

# secretWord = randomWord
var secretWord randomWord

# Replacing the letters with '_'.
var letters ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]
foreach letter letters
    replace secretWord letter "_" secretWord
endforeach

# Printing the ocult word.
say "\n$(secretWord)"

# wordIndexes = [0, 1, 2, ... length(randomWord) - 1]
length randomWord randomWordLength
rangearray randomWordLength wordIndexes

# Variable to keep track of our lives.
var lifes 5

while lifes
    # Getting user input and turning it into uppercase.
    input guessedLetter "\nLifes: $(lifes) | Letters: $(randomWordLength) | Type a letter: "
    upper guessedLetter guessedLetter
    # Checking if the word contains the letter.
    contains randomWord guessedLetter containsLetter
    if containsLetter
        var newSecretWord ""
        foreach index wordIndexes
            # currentLetter = randomWord[index]
            charat randomWord index currentLetter
            # letterInSecretWord = secretWord[index]
            charat secretWord index letterInSecretWord
            # Checking if the current letter is equal to the guessed letter.
            equal currentLetter guessedLetter match
            if match
                # If the letter is equal, concatenate the letter in 'newSecretWord'.
                concat newSecretWord currentLetter
            else
                # If the letter is NOT equal, concatenate 'letterInSecretWord' in 'newSecretWord'.
                concat newSecretWord letterInSecretWord
            endif
        endforeach
        # secretWord = newSecretWord
        var secretWord newSecretWord
        # Printing the current state of the word.
        say "\n$(secretWord)"
        # Checking victory.
        equal secretWord randomWord victory
        if victory
            say "\nYou win!"
            exit 0
        endif
    else
        # Printing the current state of the word.
        say "\n$(secretWord)"
        # Decrementing user lifes.
        decrement lifes
    endif
endwhile

# Game over if no more lifes.
say "\nGame over, you lost!"
say "The answer was $(randomWord)."
