# Arrays must have the same size.
set names ["Maria","Arold","Bianca","Maria"]
set surnames ["Johnson","Smith","Willians","Sydney"]
set grades [8.5,9,7,10]

input inputedName "Enter a name: "

length names namesLength
rangearray namesLength indexes

set found 0

foreach index indexes
    get names index name
    equal name inputedName namesMatch
    if namesMatch
        get grades index grade
        get surnames index surname
        say "Name: $(name), Surname: $(surname), Grade: $(grade)"
        increment found
    endif
endforeach

if found
    say "$(found) record(s) were found for the name $(inputedName)."
else
    say "No records were found for the name $(inputedName)."
endif
