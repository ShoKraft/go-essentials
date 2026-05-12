# Go Essentials

**A curated collection of clean, practical and production-ready Go utilities.**

This repository contains essential code snippets and utilities that every serious Go developer should know in 2026.

---

## Why This Repository Matters

- Go continues to grow in **cloud-native**, **high-performance systems**, **game development**, **scientific computing** and **backend services**.
- Understanding **complex numbers**, **bitwise operations**, **proper type usage**, and **clean formatting** significantly improves code quality and performance.
- These utilities are small, well-commented, and demonstrate modern Go best practices.

---

## Project Structure

| File / Folder                    | Description                                                                    | Importance                                                                                                                                                      |
| -------------------------------- | ------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `complex/complex.go`             | `PrettyComplex(z complex128) string` — beautiful formatting of complex numbers | Essential for scientific computing, DSP, physics simulations and graphics. Clean output (`3 + 4i`, `-2i`, `1.5 - 0.7i`) greatly improves debugging and logging. |
| `bit_tricks/bit_tricks.go`       | Fast bitwise operations (power-of-two modulo, etc.)                            | Critical for high-performance code: hash maps, buffers, allocators, game engines.                                                                               |
| `types/types.go`                 | Overview and best practices for all Go built-in types                          | Helps avoid common mistakes with `int` vs `int64`, `byte`, `rune`, `any`, etc.                                                                                  |
| `constatns/constants.go`         | Human-readable byte formatting + constants (`KB`, `MB`, `GB`...)               | Very useful for logging memory usage, file sizes, and network statistics.                                                                                       |
| `any_processor/any_processor.go` | Universal any processor file                                                   | Used in 90% of production Go services.                                                                                                                          |
| `main.go`                        | Demo of all utilities                                                          | Quick showcase of the library.                                                                                                                                  |
| `go.mod`                         | Go module file                                                                 | Standard project setup.                                                                                                                                         |

---

## Installation & Run

```bash
git clone https://github.com/ShoKraft/go-essentials.git
cd go-essentials
go run main.go




Learning Goals

Master working with complex128
Learn high-performance bitwise tricks
Understand correct type selection in Go
Write clean, professional output


Contributing
Contributions are welcome! You can add new sections like:

Slice utilities
Concurrency helpers
String manipulation
Error handling patterns


Made with ❤️ for the Go community
Last updated: May 2026
```
