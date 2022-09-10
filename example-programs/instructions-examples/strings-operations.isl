# Strings operations.

# Let's define our string:
var text "This is my awesome text!"

# LENGTH
# Usage: length MY_STRING VARIABLE_TO_STORE_VALUE
length text textLength
say "The length of the string is: $(textLength)\n"

# CHARAT
# Usage: charat MY_STRING INDEX VARIABLE_TO_STORE_VALUE
charat text 0 charAt0
say "The character at position 0 is: $(charAt0)\n"

# UPPER
# Usage: upper MY_STRING VARIABLE_TO_STORE_VALUE
upper text textToUpper
say "Text to uppercase: $(textToUpper)\n"

# LOWER
# Usage: lower MY_STRING VARIABLE_TO_STORE_VALUE
lower text textToLower
say "Text to lowercase: $(textToLower)\n"

# SPLIT
# Usage: split MY_STRING PATTERN VARIABLE_TO_STORE_VALUE
# Let's split the string using spaces " " as the pattern to get an array with the text words:
split text " " words
# firstWord = words[0]
get words 0 firstWord
say "First word of the text: $(firstWord)\n"

# REPLACE
# Usage: split MY_STRING OLD_PATTERN NEW_PATTERN VARIABLE_TO_STORE_VALUE
# Let's replace 'awesome' with 'cool'.
replace text "awesome" "cool" newText
say "Text after replace instruction: $(newText)\n"

# CONTAINS
# Usage: contains MY_STRING PATTERN_TO_FIND VARIABLE_TO_STORE_VALUE
# If the string contains the pattern, VARIABLE_TO_STORE_VALUE will be 1, else 0.
contains text "my" containsMy
contains text "rocket" containsRocket
say "The text contains the word 'my': $(containsMy)"
say "The text contains the word 'rocket': $(containsRocket)"
