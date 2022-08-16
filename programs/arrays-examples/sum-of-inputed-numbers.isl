var continue 1
var numbers []number

while continue
    input number "Type a number or type 0 to stop: "
    append numbers number
    notequal number 0 continue
endwhile

var sum 0

foreach element numbers
    add element sum sum
endforeach

say "The sum of the numbers that you typed is: $(sum)."
