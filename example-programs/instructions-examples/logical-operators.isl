# Logical Operators.

# Remember:
# numbers less than or equal to 0 are interpreted as FALSE.
# numbers greater than 1 are interpreted as TRUE.

# AND
# Usage: AND X Y VARIABLE_TO_STORE_RESULT
# 0 and 1 = 0
and 0 1 result
say "0 and 1: $(result)"
# 1 and 1 = 1
and 1 1 result
say "1 and 1: $(result)"
# 1 and (any number >= 1) = 1
and 1 5 result
say "1 and 5: $(result)\n"

# OR
# Usage: OR X Y VARIABLE_TO_STORE_RESULT
# 0 or 1 = 1
or 0 1 result
say "0 or 1: $(result)"
# 1 or 1 = 1
and 0 0 result
say "0 or 0: $(result)"
# -1 or -2 = 0
and -1 -2 result
say "-1 and -2: $(result)\n"

# NOT
# Usage: NOT X VARIABLE_TO_STORE_RESULT
# not 0 = 1
not 0 result
say "not 0: $(result)"
# not 1 = 0
not 1 result
say "not 1: $(result)"
# not -10 = 1
not -10 result
say "not -10: $(result)"
