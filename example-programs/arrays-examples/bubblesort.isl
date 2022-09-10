var array [10,3,5,1,7,9,0,2,13,99,30,23,-17,3,-20]

say "Array before bubblesort: $(array)"

length array iLimit

length array yLimit
decrement yLimit

var i 1
var y 0

less i iLimit iContinue
less y yLimit yContinue

while iContinue
    while yContinue
        get array y element1
        add y 1 yPlusOne
        get array yPlusOne element2
        greater element1 element2 firstIsGreater
        if firstIsGreater
            swap array y yPlusOne
        endif
        increment y
        less y yLimit yContinue
    endwhile
    var y 0
    var yContinue 1
    increment i
    less i iLimit iContinue
endwhile

say "Array after bubblesort: $(array)"
