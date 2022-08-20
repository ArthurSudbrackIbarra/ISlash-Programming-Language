# Reading-Writing to Files.

# Reading from Files:
# Usage: INSTRUCTION FILE_PATH VARIABLE_TO_STORE_VALUE

# Let's define a variable with the file path:
var filePath "../resources/txt/sentence.txt"

# Reading a file and storing all the content in a variable:
readfile filePath fileContent
say "=== The file content is: ===\n"
say fileContent

# Reading a file and storing each line in an array.
readfilelines filePath fileLines
# firstLine = fileLines[0]
get fileLines 0 firstLine
say "\n=== The first line is: ===\n"
say firstLine

# Writing to Files:
# Usage: writefile FILE_PATH CONTENT_TO_WRITE

# Let's define a variable with the file path:
var filePath "my-file.txt"

writefile filePath "This will be written to the file!"
say "\n=== 'my-file.txt' has been created or modified by now. ==="
