# 🧹 Code Compacter

A CLI tool designed to prepare codebases for both human and LLM consumption by optimizing code readability and structure. Unlike traditional minifiers, Code Compacter focuses on making code more accessible to both humans and LLMs by maintaining clear formatting, meaningful names, and proper structure.

## 🎯 Purpose

When working with Large Language Models (LLMs), the quality of code formatting and structure significantly impacts their ability to understand and process code effectively. Code Compacter helps by:

- Ensuring consistent, readable code formatting
- Preserving meaningful variable and function names
- Maintaining clear code structure and organization
- Removing unnecessary complexity while keeping essential context
- Preparing code that's optimized for both human review and LLM processing

## 🚀 Features

- **Readability Optimization**: Standardizes code formatting for maximum clarity
- **Structure Preservation**: Maintains logical code organization and hierarchy
- **Context Retention**: Keeps important comments and documentation
- **Smart Cleanup**: Removes redundant code while preserving functionality
- **Language-Aware**: Adapts processing based on programming language conventions

## 📦 Code Processing

### What We Keep and Enhance
- Meaningful variable and function names
- Essential comments and documentation
- Clear code structure and indentation
- Important type information
- Logical code organization

### What We Remove or Simplify
- Redundant comments and documentation
- Unused code and dead branches
- Overly complex expressions
- Unnecessary nesting
- Duplicate code patterns

## 📦 What's Included/Excluded

### Included by Default
- Source code files (*.go, *.py, *.js, *.ts, *.java, etc.)
- Project structure and organization
- Essential configuration files (if specified)

### Excluded by Default
- Documentation files (*.md, *.txt)
- Build artifacts (node_modules, dist, build)
- Version control (.git)
- Environment files (.env)
- Test directories (test, tests)
- Documentation directories (docs, documentation)
- Generated files
- Large binary files

## 💻 Usage

```bash
code-compacter /path/to/your/project
```

This will:
1. Process all relevant files in the project
2. Format the code for optimal readability
3. Output the processed code to stdout or a specified file

> **Note**: Initial development focuses on Astro projects to refine the core concepts, with plans to expand to other languages and project types.

## 🛠️ Planned Features

- [ ] Language-specific formatting rules
- [ ] Configurable formatting preferences
- [ ] Code structure analysis
- [ ] Integration with common LLM tools
- [ ] Support for multiple output formats (concatenated, structured, etc.)

## 🚀 Installation

```bash
go install github.com/holz/code-compacter@latest
```

## 🧪 Testing

```bash
go test ./tests
```

## 📝 License

MIT
