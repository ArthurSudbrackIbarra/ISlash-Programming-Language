set continue 1
set numbers []number

while continue
    input number "Type a number or type 0 to stop: "
    append numbers number
    notequal number 0 continue
endwhile

set sum 0

foreach element numbers
    add element sum sum
endforeach

say "The sum of the numbers that you typed is: $(sum)."
