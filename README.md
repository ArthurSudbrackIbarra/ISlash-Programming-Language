<h1 align="center">The ISlash Programming Language</h1>

<p align="center">
    <img src="https://raw.githubusercontent.com/ArthurSudbrackIbarra/ISlash-VSCode-Language-Support/main/icons/islash.svg" width=180px/>
</p>

**ISlash** is a dynamically typed, interpreted programming language with **no real use** that I created for fun, which runs on top of [Golang](https://go.dev/). The language somewhat resembles Assembly but it is simpler to understand and more high-level.

My main goal when creating ISlash was learning Golang (Go), as I had never used that language before. The name 'ISlash' is a pun with my last name 'Ibarra', because 'barra' means 'slash' in portuguese, which is my native language.

## Table of Contents

* [Example Programs](#example-programs)
* [Language Support in VSCode](#language-support-in-vscode)
* [Data Types](#data-types)
* [Instructions](#instructions)
* [Language Features](#language-features)
    * [Comments](#comments)
    * [String Interpolation](#string-interpolation)
    * [New Lines in Strings](#new-lines-in-strings)
* [Things to Notice](#things-to-notice)
* [Try ISlash](#try-islash)
* [Uninstall ISlash](#uninstall-islash)
* [Possible Future Features](#possible-future-features)

## Example Programs

Example programs using the ISlash language can be found inside the [programs folder](https://github.com/ArthurSudbrackIbarra/ISlash-Programming-Language/tree/main/programs).

In the example below, we are multiplying numbers inputed by the user.

![Example Program](https://user-images.githubusercontent.com/69170322/184281568-9251885e-06d5-4089-a115-b09a803836cc.png)

## Language Support in VSCode

Download the [ISlash Language Support](https://marketplace.visualstudio.com/items?itemName=ArthurSudbrackIbarra.islash-language-support) Visual Studio Code extension to get syntax highlighting, code snippets, hovering tips and language icon.

## Data Types

|   Type    | Description                                                                                                  |
|:---------:|--------------------------------------------------------------------------------------------------------------|
|   string  | Strings are declared with double quotes. Ex: "Hello!"                                                        |
|   number  | Numbers may or may not have decimal places. Ex: 1, 2.3                                                       |
|   array   | Arrays are declared with square brackets. Ex: [1,2,3], ["Hi","Hello"]. Do **not** put spaces between commas. |

In ISlash, although there is not a boolean data type, numbers can be used to represent boolean values:

| Boolean Value | Numbers Range |
|:-------------:|:-------------:|
|      true     |  numbers ≥ 1  |
|     false     |  numbers < 0  |

## Instructions

Instructions are **not case sensitive**.

| Instruction   | Description                                                                 |
|:-------------:|-----------------------------------------------------------------------------|
| VAR           | Sets/declares variables.                                                    |
| ADD           | + operator.                                                                 |
| SUB           | - operator.                                                                 |
| MULT          | * operator.                                                                 |
| DIV           | / operator.                                                                 |
| MOD           | % operator.                                                                 |
| POWER         | ^ operator.                                                                 |
| ROOT          | Square roots, cubic roots...                                                |
| INCREMENT     | ++ operator (Adds 1).                                                       |
| DECREMENT     | -- operator (Subtracts 1).                                                  |
| RANDOM        | Generates a random integer value within a range.                            |
| GREATER       | > operator.                                                                 |
| GREATEREQUAL  | >= operator.                                                                |
| LESS          | < operator.                                                                 |
| LESSEQUAL     | <= operator.                                                                |
| NOT           | NOT operator                                                                |
| AND           | AND operator.                                                               |
| OR            | OR operator.                                                                |
| IF            | If statements.                                                              |
| ELSEIF        | Else if statements.                                                         |
| ELSE          | Else statements.                                                            |
| ENDIF         | Closes if blocks.                                                           |
| EQUAL         | == operator.                                                                |
| NOTEQUAL      | != operator.                                                                |
| CONCAT        | Concatenates strings.                                                       |
| LENGTH        | Gets the length of a string or an array.                                    |
| UPPER         | Turn strings into uppercase.                                                |
| LOWER         | Turn strings into lowercase.                                                |
| SPLIT         | Splits a string using a pattern, produces a string[] variable.              |
| REPLACE       | Replaces a pattern in a string by another pattern.                          |
| CHARAT        | Gets the nth char of a string.                                              |
| SAY           | Prints to screen.                                                           |
| INPUT         | Gets user input.                                                            |
| WHILE         | While statements.                                                           |
| BREAK         | Exits out of while blocks.                                                  |
| ENDWHILE      | Closes while blocks.                                                        |
| FOREACH       | Use to iterate over arrays.                                                 |
| ENDFOREACH    | Closes foreach blocks.                                                      |
| APPEND        | Appends an element to an array.                                             |
| PREPEND       | Preppends an element to an array.                                           |
| REMOVEFIRST   | Removes the first element of an array.                                      |
| REMOVELAST    | Removes the last element of an array.                                       |
| SWAP          | Swaps arrays positions.                                                     |
| GET           | Gets the nth element of an array.                                           |
| SETINDEX      | Changes the element at an index.                                            |
| CONTAINS      | Checks if a string contains a character or if an array contains an element. |
| READFILE      | Reads a file, produces a string variable.                                   |
| READFILELINES | Reads a file line by line, produces a string[] variable.                    |
| WRITEFILE     | Writes to a file, overrides previous content.                               |
| EXIT          | Exits the program with a status code.                                       |

## Language Features

Below, some ISlash language features will be explained:

### Comments

Comments can be made using the `#` character at the beginning of lines:

```
# This is a comment!
say "Cool!"
```

### String interpolation

ISlash allows the interpolation of Strings using the `$()` symbol:

```
var name "Arthur"
var age 20
say "My name is $(name) and I am $(age) years old."
```

### New Lines in Strings

To represent new lines, use the `\n` symbol:

```
say "Hi!\nThis is in a new line!"
```

## Things to Notice

Here are some important things to notice about the ISlash language:

### 1. Variables

1. All variables are global and can be accessed from anywhere, as there are no scopes. If you create a variable inside an *if* block, you **will** be able to reference it outside of the *if* block. The only **exception** to this rule are *foreach* element variables, which only exists inside the *foreach* block.

Example:

```
foreach element [1,2,3]
    # element exists here!
endforeach
# element does not exist here!
```

2. The **VAR** command will create new variables if they don't exist yet. If they do, the previous value will be replaced by the new one.

### 2. Arrays

1. Even though ISlash is a dynamically typed language, arrays **cannot** contain values of different types. Because of that, when creating an **empty array**, you must specify if it will contain numbers or strings.

Example:

```
var array [1,2,3] # OK!
var array ["Hello", "Bye"] # OK!
var array []number # OK!
var array []string # OK!

var array [] # NOT OK!
```

## Try ISlash

To try the ISlash language, follow the steps being described below:

1. Download [Docker](https://www.docker.com/products/docker-desktop/).

2. Clone this repository.

```sh
git clone https://github.com/ArthurSudbrackIbarra/ISlash-Programming-Language.git
```

3. Go to the repository directory.

```sh
cd ISlash-Programming-Language
```

4. Start the Docker container:

```sh
docker compose up -d
```

5. Enter inside the Docker container that you started:

```sh
docker exec -it islash-container /bin/bash
```

6. Run the ISlash programs you wish with:

```sh
islash <PATH_TO_MY_PROGRAM>

# Example:
islash myProgram.isl
```

![Running Programs](https://user-images.githubusercontent.com/69170322/183551455-e2b7d46f-7115-4a69-a03a-eafb5b67a323.png)

**NOTE**: All files inside the 'programs' directory are shared between your host machine and the Docker ISlash container using a **bind mount volume**, so you can modify the .isl files or create new ones in your host machine and then run them from inside the container. 

## Uninstall ISlash

To completely delete all the resources that ISlash created in your machine, use the following commands:

1. Stop the container.

```sh
docker compose down
```

2. Delete the container.

```sh
docker rm islash-container
```

3. Delete the container image.

```sh
docker rmi islash/islash-programming-language:v1
```

## Possible Future Features

These are some of the features that I might implement in the future:

- Functions
- Swich/Case
