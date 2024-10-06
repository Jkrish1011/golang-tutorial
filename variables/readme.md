# Basic Variables

bool: a boolean value, either true or false

string: a sequence of characters

int: a signed integer

float64: a floating-point number

byte: exactly what it sounds like: 8 bits of data

also

int8, int16, int32, int64
uint8, uint 16, uint 32, uint64, uintptr

byte // alias for uint8

rune // alias for int32 and represents a unicode code point

float32, float64

complex64, complex128

# Declaring a variable the sad way

### var mySkillIssues int

or

### mySkillIssues = 42

The first line, var mySkillIssues int, defaults the mySkillIssues variable to its zero value, 0. On the next line, 42 overwrites the zero value.
