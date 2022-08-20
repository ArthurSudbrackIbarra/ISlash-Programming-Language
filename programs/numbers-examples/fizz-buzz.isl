input n "Enter a value for n: "

var index 1

lessequal index n indexIsLessThanOrEqualN

while indexIsLessThanOrEqualN
    mod index 3 remainer3
    mod index 5 remainer5

    not remainer3 divisibleBy3
    not remainer5 divisibleBy5

    var output ""

    if divisibleBy3
        concat output "Fizz"
    endif

    if divisibleBy5
        concat output "Buzz"
    endif

    say "$(index): $(output)"
    say "----------------"

    increment index
    lessequal index n indexIsLessThanOrEqualN
endwhile
