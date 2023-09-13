# ascii-art

## ASCII Template Reader

This program reads and processes ASCII templates from a text file and provides functionalities to display and manipulate them.

## Usage

To run the program, execute the following command in your terminal:

go run main.go <input_string>


- `<input_string>`: Optional argument specifying the input string.

The program will display the input string followed by the processed ASCII templates.

## Features

1. **Reading ASCII Templates from File**

The program reads ASCII templates from the `standard.txt` file. Each template is separated by a blank line, and each line in a template corresponds to a row of the ASCII art.

2. **Displaying ASCII Templates**

The `ReadStandardTxt` function reads the `standard.txt` file and processes the ASCII templates. It prints each line of each template, excluding the first line if it's empty.

3. **TODO: Character Transformation**

There's a placeholder function named `printMultipleCharacter` that is meant to handle printing multiple characters in a row, similar to the example provided in the comment.

## Notes

- The program currently focuses on reading and displaying ASCII templates. Additional functionalities can be added based on the project's requirements.

- Be cautious when handling newline characters ('\n') within the templates. Different logic might be needed for templates with multiple newlines.

## Example

Suppose the `standard.txt` file contains the following templates:


Running the program with the command:

go run main.go "Hello"


## Future Enhancements

- Add functionality to print multiple characters in a row using `printMultipleCharacter` function.
- Explore ways to handle different newline patterns within templates.

