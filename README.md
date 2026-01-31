<h1 align="center">
  <img
    src="https://github.com/user-attachments/assets/f95690a5-3b63-4761-89eb-26ee4ac2f728"
    alt="tuidoo"
    height="200"
  />
</h1>

<p align="center">
  <img src="https://img.shields.io/badge/language-Go-blue" />
  <img src="https://img.shields.io/badge/license-GPLv3-blue.svg" />
  <img src="https://img.shields.io/badge/interface-TUI-green" />
  <img src="https://img.shields.io/badge/platform-terminal-black" />
  <img src="https://img.shields.io/badge/style-retro--terminal-brightgreen" />
</p>

<p align="center">
  <em>A retro TUI todo manager</em>
</p>

---

## 📟 About

**tuidoo** is an early-stage terminal-based todo manager written in Go.  
It focuses on **keyboard-first workflows**, **fast startup**, and a **retro terminal aesthetic** inspired by classic computing.

The goal is to keep task management simple, fast, and enjoyable — without leaving the terminal.

> ⚠️ **Status:** Pre-alpha. Expect breaking changes.

---

## ✨ Planned Features

- Create, edit, complete, and delete todos  
- Multiple lists / contexts  
- Persistent local storage  
- Fully keyboard-driven navigation  
- Retro terminal UI using `tview`  
- Cross-platform support (macOS, Linux, Windows)

---

## 🗺️ Roadmap

### Phase 1 — Core Functionality
- [ ] Project skeleton and command entry point  
- [ ] Basic TUI layout (list + detail view)  
- [ ] In-memory todo model  
- [ ] Keyboard navigation  
- [ ] Quit / confirm handling  

### Phase 2 — Persistence
- [ ] Local storage (file-based)  
- [ ] Load/save on startup/exit  
- [ ] Data migration strategy  

### Phase 3 — UX & Polish
- [ ] Retro color themes  
- [ ] Config file support  
- [ ] Help / keybinding overlay  
- [ ] Error handling and recovery  

### Phase 4 — Distribution
- [ ] Versioned releases  
- [ ] Prebuilt binaries  
- [ ] Documentation polish  

---

## 🚀 Getting Started (Development)

### Prerequisites

- Go **1.22+**  
- A terminal with ANSI color support

### Clone & build

```bash
git clone https://github.com/YOURUSER/tuidoo.git
cd tuidoo
go build ./...
```

### Run locally:

```bash
./tuidoo
```

## 🛠️ Tech Stack

- Language: Go
- TUI Framework: tview
- License: GPLv3

🤝 Contributing
This project is open to contributions — see CONTRIBUTING.md for details.

📄 License
Licensed under the GNU General Public License v3.0.
See the LICENSE file for full text.
