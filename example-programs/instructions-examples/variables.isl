# Variables.

# How to declare/set variables:
# var VARIABLE_NAME VALUE_TO_STORE

# Numbers:
# Integers:
var one 1
# Decimals:
var half 0.5
say "Integer number one: $(one)"
say "Decimal number 1/2: $(half)\n"

# Strings:
var str "my string"
var emptyStr ""
var strWithInterpolation "Look, this is $(str)."
say "A string: $(str)"
say "An empty string: $(emptyStr)"
say "A string declared using interpolation: $(strWithInterpolation)\n"

# Arrays:
# Number arrays:
var notEmptyNumArray [1,2,3]
var emptyNumArray []number
say "Not empty number array: $(notEmptyNumArray)"
say "Empty number array: $(emptyNumArray)"
# String arrays:
var notEmptyStrArray ["Hello","Bye","Good Morning"]
var emptyStrArray []string
say "Not empty string array: $(notEmptyStrArray)"
say "Empty string array: $(emptyStrArray)\n"

# Notice: Empty arrays must be declared with type specification due to ambiguity.

# The 'var' instruction is also used to override variable values.
var ten 20
var ten 10
say "Ten variable: $(ten)\n"

# The 'VALUE_TO_STORE' argument can also be another variable.
var array [10,20,30]
var anotherArray array
say "Now, 'anotherArray' is a copy of 'array': $(anotherArray)"
