# txtlang: Coding in Natural Language

## Introduction

txtlang is an innovative programming language designed to bridge the gap between natural human language and machine code. It aims to make programming more intuitive and accessible to a wider audience, including those with little to no prior coding experience.

## Key Features

- **Natural Language Syntax**: Write code that reads like everyday language.
- **Intuitive Structure**: Organize your code in a way that makes sense to you.
- **Flexible Typing**: Automatic type inference with optional explicit typing.
- **Cross-Language Compatibility**: Compile txtlang to various programming languages.

## How to Build

To build txtlang, you need to have Go installed on your system. Follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/arielmiki/txtlang.git
   cd txtlang
   ```

2. Build the project using the provided Makefile:
   ```
   make
   ```

This will create an executable named `txtlang` in your current directory.

## How to Use

Once you've built txtlang, you can use it to translate your natural language code to Python or Go. Here's how to use the tool:

1. Create a txtlang file (e.g., `mycode.txt`) with your natural language code.

2. Run the txtlang translator:
   ```
   ./txtlang -i mycode.txt -l <language> -o output.<ext>
   ```
   Replace `<language>` with either `python` or `golang`, and `<ext>` with the appropriate file extension (`.py` for Python or `.go` for Go).

3. The translated code will be saved in the specified output file.

