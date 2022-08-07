# The ISlash Programming Language

**ISlash** is a dynamically typed, interpreted programming language with **no real use** that I created for fun, which runs on top of [Golang](https://go.dev/). The language somewhat resembles Assembly but it is simpler to understand and more high-level.

My main goal when creating ISlash was learning Golang (Go), as I had never used that language before. The name 'ISlash' is a pun with my last name 'Ibarra', because 'barra' means 'slash' in portuguese, which is my native language.
 
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

|  Instruction |        Description       |
|:------------:|:------------------------:|
|    DECLARE   |    Declares variables.   |
|      ADD     |        + operator.       |
|      SUB     |        - operator.       |
|     MULT     |        * operator.       |
|      DIV     |        / operator.       |
|      MOD     |        % operator.       |
|    GREATER   |        > operator.       |
| GREATEREQUAL |       >= operator.       |
|    LESSER    |        < operator.       |
|  LESSEREQUAL |       <= operator.       |
|      NOT     |  NOT operator. (TBD...)  |
|      AND     |  AND operator. (TBD...)  |
|      OR      |   OR operator. (TBD...)  |
|      IF      |      If statements.      |
|     ELSE     |     Else statements.     |
|     ENDIF    |   Closes if statements.  |
|     EQUAL    |   Compares 2 variables.  |
|    CONCAT    |   Concatenates strings.  |
|      SAY     |     Prints to screen.    |
|     WHILE    |     While statements.    |
|   ENDWHILE   | Closes while statements. |

## Language Features

Below, ISlash language features will be explained:

### String interpolation

ISlash allows the interpolations of Strings using the `$()` operator:

```
declare name "Arthur"
declare age 20
say "My name is $(name) and I am $(age) years old."
```

# Example Programs

Example programs using the ISlash language can be found inside the [programs folder](https://github.com/ArthurSudbrackIbarra/ISlash-Programming-Language/tree/main/programs).

## How to Use

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
docker compose up
```

5. Enter inside the Docker container that you started:

```sh
docker exec -it islash-container /bin/bash
```

6. Run the ISlash programs you wish with:

```sh
islash <PATH_TO_MY_PROGRAM>

# Example:
islash even-or-odd.islash
```

**NOTE**: All files inside the 'programs' directory are shared between your host machine and the Docker ISlash container using a **bind mount volume**, so you can modify the .islash files or create new ones in your host machine and then run them from inside the container. 
