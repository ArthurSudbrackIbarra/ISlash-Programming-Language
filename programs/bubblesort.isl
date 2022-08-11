# Declaring the unsorted array:
set array [10,3,5,1,7,9,0,2,13,99,30,23]

length array iLimit

length array yLimit
decrement yLimit

set i 1
set y 0

less i iLimit iContinue
less y yLimit yContinue

while iContinue
    while yContinue
        accessindex array y element1
        add y 1 yPlusOne
        accessindex array yPlusOne element2
        greater element1 element2 firstIsGreater
        if firstIsGreater
            swap array y yPlusOne
        endif
        say array
        increment y
        less y yLimit yContinue
    endwhile
    set y 0
    set yContinue 1
    increment i
    less i iLimit iContinue
endwhile

say array
