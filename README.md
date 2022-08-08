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

|    Instruction   |         Description        |
|:----------------:|:--------------------------:|
|      DECLARE     |     Declares variables.    |
|        ADD       |         + operator.        |
|        SUB       |         - operator.        |
|       MULT       |         * operator.        |
|        DIV       |         / operator.        |
|        MOD       |         % operator.        |
|     INCREMENT    |    ++ operator (Adds 1).   |
|     DECREMENT    | -- operator (Subtracts 1). |
|    GREATERTHAN   |         > operator.        |
| GREATERTHANEQUAL |        >= operator.        |
|     LESSTHAN     |         < operator.        |
|   LESSTHANEQUAL  |        <= operator.        |
|        NOT       |        NOT operator        |
|        AND       |        AND operator.       |
|        OR        |        OR operator.        |
|        IF        |       If statements.       |
|       ELSE       |      Else statements.      |
|       ENDIF      |    Closes if statements.   |
|       EQUAL      |    Compares 2 variables.   |
|      CONCAT      |    Concatenates strings.   |
|        SAY       |      Prints to screen.     |
|       INPUT      |      Gets user input.      |
|       WHILE      |      While statements.     |
|     ENDWHILE     |  Closes while statements.  |

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

# Example Programs

Example programs using the ISlash language can be found inside the [programs folder](https://github.com/ArthurSudbrackIbarra/ISlash-Programming-Language/tree/main/programs).

## How to Use

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
islash even-or-odd.islash
```

![Running Programs](https://user-images.githubusercontent.com/69170322/183312708-8dfb28cc-7c13-418f-9f92-bb55927b13ab.png)

**NOTE**: All files inside the 'programs' directory are shared between your host machine and the Docker ISlash container using a **bind mount volume**, so you can modify the .islash files or create new ones in your host machine and then run them from inside the container. 

## How to Delete

To completely delete all the resources that islash created in your machine, use the following commands:

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