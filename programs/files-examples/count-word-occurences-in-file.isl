var wordToFind "WATERMELONS"
var occurences 0

var filePath "../resources/txt/sentence.txt"
readfilelines filePath fileLines

foreach line fileLines
    split line " " splittedLine
    foreach word splittedLine
        upper word wordUpper
        equal wordUpper wordToFind match
        if match
            increment occurences
        endif
    endforeach    
endforeach

say "$(occurences) occurences of $(wordToFind) were found."
