<h1 align="center">The ISlash Programming Language</h1>

**ISlash** is a dynamically typed, interpreted programming language with **no real use** that I created for fun, which runs on top of [Golang](https://go.dev/). The language somewhat resembles Assembly but it is simpler to understand and more high-level.

My main goal when creating ISlash was learning Golang (Go), as I had never used that language before. The name 'ISlash' is a pun with my last name 'Ibarra', because 'barra' means 'slash' in portuguese, which is my native language.

## Table of Contents

* [Data Types](#data-types)
* [Instructions](#instructions)
* [Language Features](#language-features)
    * [Comments](#comments)
    * [String Interpolation](#string-interpolation)
    * [New Lines in Strings](#new-lines-in-strings)
* [Example Programs](#example-programs)
* [Download ISlash](#download-islash)
* [Uninstall ISlash](#uninstall-islash)
 
## Data Types

| Data Type |                       Description                      |
|:---------:|:------------------------------------------------------:|
|   string  |  Strings are declared with double quotes. Ex: "Hello!" |
|   number  | Numbers may or may not have decimal places. Ex: 1, 2.3 |

In ISlash, although there is not a boolean data type, numbers can be used to represent boolean values:

| Boolean Value | Numbers Range |
|---------------|---------------|
| true          | numbers ≥ 1   |
| false         | numbers < 0   |

## Instructions

Instructions are **not case sensitive**.

|    Instruction   |          Description          |
|:----------------:|:-----------------------------:|
|      DECLARE     |      Declares variables.      |
|        ADD       |          + operator.          |
|        SUB       |          - operator.          |
|       MULT       |          * operator.          |
|        DIV       |          / operator.          |
|        MOD       |          % operator.          |
|     INCREMENT    |     ++ operator (Adds 1).     |
|     DECREMENT    |   -- operator (Subtracts 1).  |
|    GREATERTHAN   |          > operator.          |
| GREATERTHANEQUAL |          >= operator.         |
|     LESSTHAN     |          < operator.          |
|   LESSTHANEQUAL  |          <= operator.         |
|        NOT       |          NOT operator         |
|        AND       |         AND operator.         |
|        OR        |          OR operator.         |
|        IF        |         If statements.        |
|       ELSE       |        Else statements.       |
|       ENDIF      |     Closes if statements.     |
|       EQUAL      |          == operator.         |
|     NOTEQUAL     |          != operator.         |
|      CONCAT      |     Concatenates strings.     |
|      LENGTH      |  Gets the length of a string. |
|      GETCHAR     | Gets the nth char of a string |
|        SAY       |       Prints to screen.       |
|       INPUT      |        Gets user input.       |
|       WHILE      |       While statements.       |
|     ENDWHILE     |    Closes while statements.   |

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

## Example Programs

Example programs using the ISlash language can be found inside the [programs folder](https://github.com/ArthurSudbrackIbarra/ISlash-Programming-Language/tree/main/programs).

In the example below, we are checking whether or not a string inputed by the user is a palindrome:

![Example Program](https://user-images.githubusercontent.com/69170322/183494031-69d39191-e8e0-4b40-82a5-e17d762e7462.png)

## Download ISlash

To test the ISlash language, follow the steps being described below:

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

**NOTE**: All files inside the 'programs' directory are shared between your host machine and the Docker ISlash container using a **bind mount volume**, so you can modify the .islash files or create new ones in your host machine and then run them from inside the container. 

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