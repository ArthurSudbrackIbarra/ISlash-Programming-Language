input n "Enter the value for n: "

set num1 1
set num2 1

set sequence []number

while n
    append sequence num1
    set aux num1
    set num1 num2
    add aux num2 num2
    decrement n
endwhile

say sequence
