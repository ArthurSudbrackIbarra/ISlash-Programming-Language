# CSV file path.
var csvPath "../resources/csv/names.csv"

# Reading files and storing content as string[] in csvContent.
readfilelines csvPath csvLines

# Storing the CSV header line in 'headerLine'.
removefirst csvLines headerLine

# Extracting the csv header fields.
split headerLine "," headerFields

# Creating an array from 0 to length(headerLine) - 1.
length headerFields headerLength
rangearray headerLength headerIndexes

# Creating an array from 1 to length(csvLines) - 1.
length csvLines dataLength
rangearray dataLength dataIndexes
removefirst dataIndexes _

# Foreach loop.
foreach dIndex dataIndexes
    get csvLines dIndex dataLine
    split dataLine "," dataFields
    foreach hIndex headerIndexes
        get headerFields hIndex fieldName
        get dataFields hIndex fieldValue
        say "$(fieldName): $(fieldValue)"
    endforeach
    say "======================"
endforeach
