# Contributing to tuidoo

Thanks for your interest in contributing to **tuidoo**!  
This project is licensed under the **GNU General Public License v3.0 (GPLv3)**, and all contributions must comply with its terms.

---

## 📜 License & Contribution Terms

By contributing to this project, you agree that:

- Your contributions are licensed under **GPLv3**
- You have the right to submit the code you are contributing
- You are not submitting proprietary or confidential code
- You understand that GPLv3 requires derivative works to also be GPLv3

If you do not agree with these terms, please do not submit contributions.

---

## 🛠️ How to Contribute

### 1. Fork the repository

Create your own fork on GitHub.

### 2. Create a feature branch

```bash
git checkout -b feature/my-feature
```

### 3. Make your changes

- Keep changes focused and minimal
- Follow idiomatic Go practices
- Prefer clarity over cleverness

### 4. Test your changes

Ensure the application builds and runs correctly.

```bash
go build ./...
go test ./...
```
### 5. Commit your work

- Write clear, descriptive commit messages, for example:
  - Add basic todo list navigation
  - Fix crash when deleting last item

### 6. Open a Pull Request

- Describe what you changed and why
- Reference issues where applicable
- Be open to feedback

## 🧹 Code Style & Guidelines

- Follow standard Go formatting (```gofmt```)
- Keep functions small and readable
- Avoid unnecessary abstractions
- Prefer explicit behavior over magic

## 🐛 Reporting Issues

- When reporting bugs, please include:
- What you expected to happen
- What actually happened
- Steps to reproduce
- Your OS and Go version

## 💡 Feature Requests

- Feature ideas are welcome, but keep in mind:
- tuidoo aims to stay simple and fast
- Features should fit a terminal-first workflow
- Large changes should be discussed before implementation

## ❤️ Final Note
- tuidoo is an early-stage project.
- Constructive feedback, patience, and collaboration are appreciated.
- Thanks for helping make tuidoo better!
