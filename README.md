# Tetris Optimizer

This Go program reads a list of tetrominoes from a text file and assembles them into the smallest square possible. Each tetromino is identified by uppercase Latin letters (A, B, C, etc.). The program handles errors gracefully and outputs "ERROR" if there is an issue with the file format or the tetrominoes.

## Features
- **Smallest Square**: Assembles tetrominoes to create the smallest possible square.
- **Tetromino Identification**: Identifies each tetromino using uppercase Latin letters.
- **Error Handling**: Prints "ERROR" for bad file formats or tetromino formats.
- **File Reading**: Reads tetrominoes from a specified text file.

## Instructions

1. **Running the Program**:
   - Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
   - Clone this **tetris-optimizer** repository.
   - Navigate to the directory containing the cloned repository.
   - Run the program using the following command:
     ```console
     go run . sample.txt
     ```
   - You can test the program using different tetrominoes by running:
     ```console
     go run . path/to/your/sample.txt
     ```   

2. **File Format**:
   - The text file should contain tetrominoes separated by a single blank line (More than one blank line is a wrong file format)
   - Each tetromino should be represented as a 4x4 grid of characters (`#` for filled and `.` for empty).
   - Example of a valid file format:

     ```
     #...
     #...
     #...
     #...

     ....
     ....
     ..##
     ..##
     ```

3. **Output**:
   - The program will output the assembled board with tetrominoes labeled with letters (A, B, C, etc.).
   - If it's not possible to form a complete square, the program will leave spaces between tetrominoes.

The following will be the output of the valid file format: 

```
ABB.
ABB.
A...
A...
```

## Contribution
Contributions are welcome! If you have any suggestions, improvements, or ideas, please feel free to open a Pull Request (PR). Here's how you can contribute:

1. Fork the repository.
2. Create a new branch for your feature or fix:
```bash
   git checkout -b your-branch-name
```
Push your changes to your forked repository:

```console
git push origin your-branch-name
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.