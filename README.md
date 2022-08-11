<h1 align="center">The ISlash Programming Language</h1>

<p align="center">
    <img src="https://raw.githubusercontent.com/ArthurSudbrackIbarra/ISlash-VSCode-Language-Support/main/icons/islash.svg" width=180px/>
</p>

**ISlash** is a dynamically typed, interpreted programming language with **no real use** that I created for fun, which runs on top of [Golang](https://go.dev/). The language somewhat resembles Assembly but it is simpler to understand and more high-level.

My main goal when creating ISlash was learning Golang (Go), as I had never used that language before. The name 'ISlash' is a pun with my last name 'Ibarra', because 'barra' means 'slash' in portuguese, which is my native language.

## Table of Contents

* [Example Programs](#example-programs)
* [Data Types](#data-types)
* [Instructions](#instructions)
* [Language Features](#language-features)
    * [Comments](#comments)
    * [String Interpolation](#string-interpolation)
    * [New Lines in Strings](#new-lines-in-strings)
* [Things to Notice](#things-to-notice)
* [Try ISlash](#try-islash)
* [Language Support in VSCode](#language-support-in-vscode)
* [Uninstall ISlash](#uninstall-islash)
* [Known Issues](#known-issues)

## Example Programs

Example programs using the ISlash language can be found inside the [programs folder](https://github.com/ArthurSudbrackIbarra/ISlash-Programming-Language/tree/main/programs).

In the example below, we are calculating the sum of *X* numbers inputed by the user.

![Example Program](https://user-images.githubusercontent.com/69170322/184058971-f64d1b1f-2f5c-4ce1-89d5-8e4bdc8f3f83.png)

## Data Types

| Data Type | Description                                                                                                  |
|-----------|--------------------------------------------------------------------------------------------------------------|
| string    | Strings are declared with double quotes. Ex: "Hello!"                                                        |
| number    | Numbers may or may not have decimal places. Ex: 1, 2.3                                                       |
| array     | Arrays are declared with square brackets. Ex: [1,2,3], ["Hi","Hello"]. Do **not** put spaces between commas. |

In ISlash, although there is not a boolean data type, numbers can be used to represent boolean values:

| Boolean Value | Numbers Range |
|---------------|---------------|
| true          | numbers ≥ 1   |
| false         | numbers < 0   |

## Instructions

Instructions are **not case sensitive**.

|  Instruction |                Description               |
|:------------:|:----------------------------------------:|
|      SET     |         Sets/declares variables.         |
|      ADD     |                + operator.               |
|      SUB     |                - operator.               |
|     MULT     |                * operator.               |
|      DIV     |                / operator.               |
|      MOD     |                % operator.               |
|   INCREMENT  |           ++ operator (Adds 1).          |
|   DECREMENT  |        -- operator (Subtracts 1).        |
|    GREATER   |                > operator.               |
| GREATEREQUAL |               >= operator.               |
|     LESS     |                < operator.               |
|   LESSEQUAL  |               <= operator.               |
|      NOT     |               NOT operator               |
|      AND     |               AND operator.              |
|      OR      |               OR operator.               |
|      IF      |              If statements.              |
|     ELSE     |             Else statements.             |
|     ENDIF    |             Closes if blocks.            |
|     EQUAL    |               == operator.               |
|   NOTEQUAL   |               != operator.               |
|    CONCAT    |           Concatenates strings.          |
|    LENGTH    | Gets the length of a string or an array. |
|    GETCHAR   |      Gets the nth char of a string.      |
|      SAY     |             Prints to screen.            |
|     INPUT    |             Gets user input.             |
|     WHILE    |             While statements.            |
|   ENDWHILE   |           Closes while blocks.           |
|    FOREACH   |        Use to iterate over arrays.       |
|  ENDFOREACH  |          Closes foreach blocks.          |
|     BREAK    |           (Not implemented yet)          |
|    APPEND    |      Appends an element to an array.     |
|  ACCESSINDEX |     Gets the nth element of an array.    |

## Language Features

Below, ISlash language features will be explained:

### Comments

Comments can be made using the `#` character at the beginning of lines:

```
# This is a comment!
say "Cool!"
```

### String interpolation

ISlash allows the interpolation of Strings using the `$()` symbol:

```
declare name "Arthur"
declare age 20
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

2. The **SET** command will create new variables if they don't exist yet. If they do, the previous value will be overrided.

### 2. Arrays

1. Even though ISlash is a dynamically typed language, arrays **cannot** contain values of different types. Because of that, when creating an **empty array**, you must specify if it will contain numbers or strings.

Example:

```
set array [1,2,3] # OK!
set array ["Hello", "Bye"] # OK!
set array []number # OK!
set array []string # OK!

set array [] # NOT OK!
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

## Language Support in VSCode

The ISlash language support Visual Studio Code extension will be released soon...

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

## Known Issues

The issues listed below are known by me and **will be fixed soon**:

* It is currently not possible to compare arrays with the 'equal' or 'notequal' instructions.
