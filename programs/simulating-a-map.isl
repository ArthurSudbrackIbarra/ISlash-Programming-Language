# Arrays must have the same size.
set names ["Maria","Arold","Bianca","Maria"]
set surnames ["Johnson","Smith","Willians","Sydney"]
set grades [8.5,9,7,10]

input inputedName "Enter a name: "

set index 0
length names namesLength
less index namesLength continue

set found 0

# While index < length of 'names'.
while continue
    accessindex names index name
    equal name inputedName namesMatch
    if namesMatch
        accessindex grades index grade
        accessindex surnames index surname
        say "Name: $(name), Surname: $(surname), Grade: $(grade)"
        increment found
    endif
    increment index
    less index namesLength continue
endwhile

if found
    say "$(found) record(s) were found for the name $(inputedName)."
else
    say "No records were found for the name $(inputedName)."
endif
