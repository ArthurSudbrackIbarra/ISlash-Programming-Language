input str1 "Enter a string: "
input str2 "Enter another string: "

length str1 str1Len
length str2 str2Len

lessthan str2Len str1Len str2LenIsShorter

if str2LenIsShorter
    set limit str2Len
else
    set limit str1Len
endif

set counter 0
notequal counter limit shouldContinue

set contains 1

while shouldContinue
    getchar str1 counter str1Char
    getchar str2 counter str2Char
    notequal str1Char str2Char charsAreNotEqual
    if charsAreNotEqual
        decrement contains
    endif
    increment counter
    notequal counter limit shouldContinue
endwhile

if contains
    say "The string '$(str1)' contains the string '$(str2)'"
else
    say "The string '$(str1)' DOES NOT contains the string '$(str2)'"
endif
