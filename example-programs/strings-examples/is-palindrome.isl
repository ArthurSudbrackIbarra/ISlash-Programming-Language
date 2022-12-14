input str "Enter a string: "

# index1 = 0
var index1 0

# index2 = length(str) - 1
length str index2
decrement index2

# limit = index2 / 2
div index2 2 limit

# isPalidrome = true
var isPalidrome 1

# while limit > 0
while limit
    # char1 = str[index1]
    # char2 = str[index2]
    charat str index1 char1
    charat str index2 char2

    # charsAreNotEqual = (char1 != char2)
    notequal char1 char2 charsAreNotEqual

    if charsAreNotEqual
        # isPalidrome will be 0, which means false.
        decrement isPalidrome
    endif

    increment index1
    decrement index2
    decrement limit
endwhile

if isPalidrome
    say "The string '$(str)' is a palindrome!"
else
    say "The string '$(str)' is NOT a palindrome!"
endif
