# Condition blocks.

# Usage:
# IF CONDITION
#   Your logic goes here...
# ELSEIF (OPTIONAL)
#   Your logic goes here...
# ELSE (OPTIONAL)
#   Your logic goes here...
# ENDIF

# The 'CONDITION' argument can be a number, a string or an array.
#
#   Numbers:
#       <= 0 FALSE
#       > 1 TRUE
#
#   Strings and Arrays:
#       length 0 FALSE
#       length > 0 TRUE

# Let's define some variables:
var zero 0
var one 1
var five 5

var horse "Horse"
var emptyStr ""

# Will enter the if block.
if one
    say "(IF 1) This will be printed.\n"
endif
# Will enter the if block.
if five
    say "(IF 2) This will be printed.\n"
endif
# Will NOT enter the if block.
if zero
    say "(IF 3) This will NOT be printed.\n"
endif

# Will enter the if block.
if horse
    say "(IF 4) This will be printed.\n"
endif
# Will NOT enter the if block.
if emptyStr
    say "(IF 5) This will NOT be printed.\n"
endif

# Else if and else.
if 0
    say "(IF 6) This will NOT be printed.\n"
elseif 1
    say "(ELSEIF 7) This will be printed.\n"
endif

# Else if and else.
if 0
    say "(IF 8) This will NOT be printed.\n"
elseif 0
    say "(ELSEIF 9) This will be printed.\n"
else
    say "(ELSE) This will be printed."
endif
