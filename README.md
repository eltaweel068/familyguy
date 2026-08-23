<div align="center">

# 📺 Family Guy CLI & TUI

```text
███████╗ █████╗ ███╗   ███╗██╗██╗  ██╗   ██╗   ██████╗ ██╗   ██╗██╗   ██╗
██╔════╝██╔══██╗████╗ ████║██║██║  ╚██╗ ██╔╝  ██╔════╝ ██║   ██║╚██╗ ██╔╝
█████╗  ███████║██╔████╔██║██║██║   ╚████╔╝   ██║  ███╗██║   ██║ ╚████╔╝ 
██╔══╝  ██╔══██║██║╚██╔╝██║██║██║    ╚██╔╝    ██║   ██║██║   ██║  ╚██╔╝  
██║     ██║  ██║██║ ╚═╝ ██║██║███████╗██║     ╚██████╔╝╚██████╔╝   ██║   
╚═╝     ╚═╝  ╚═╝╚═╝     ╚═╝╚═╝╚══════╝╚═╝      ╚═════╝  ╚═════╝    ╚═╝   
```

**Interactive dual-pane TUI & command-line media streamer for Family Guy.**

[![Go Version](https://img.shields.io/github/go-mod/go-version/eltaweel068/familyguy?style=flat-square&logo=go&logoColor=white&color=00ADD8)](https://golang.org)
[![Release](https://img.shields.io/github/v/release/eltaweel068/familyguy?style=flat-square&color=3FB950&logo=github)](https://github.com/eltaweel068/familyguy/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-8A2BE2?style=flat-square)](https://github.com/eltaweel068/familyguy/releases)

<br/>

[Overview](#overview) • [Preview](#preview) • [Features](#features) • [Installation](#installation) • [Usage](#usage) • [Keybindings](#keyboard-shortcuts) • [Architecture](#project-architecture) • [Roadmap](#roadmap)

</div>

---

## Overview

**Family Guy CLI & TUI** is an interactive terminal application and command-line media navigator built with **Go**, **Bubble Tea**, and **Lip Gloss**.

It functions both as a full visual, keyboard-driven dual-pane TUI for exploring **23+ seasons** and as a fast, scriptable CLI tool for instant episode playback (`familyguy s5 e4`). Streams seamlessly via your **Default Web Browser** or **VLC Media Player**.

---

## Preview

*( demo recording soon)*

> **Architecture Note:** The dual-pane TUI layout was prototyped via vector wireframing ([designUI.svg](designUI.svg)) before implementation with Charm Bubble Tea & Lip Gloss.

### Dual-Pane Episode Explorer

```text
Seasons:          │ Episodes (season :5):

  Season 01       │ ▸01 │ Stewie Loves Lois
  Season 02       │  02 │ Mother Tucker
  Season 03       │  03 │ Hell Comes to Quahog
  Season 04       │  04 │ Saving Private Brian
• Season 05       │  05 │ Whistle While Your Wife Works
  Season 06       │  06 │ Prick Up Your Ears
  Season 07       │  07 │ Chick Cancer
  Season 08       │  08 │ Barely Legal
  ...             │  ...

[↑/↓: Navigate • →/Enter: Open Season • ←: Back • Esc / Ctrl+C: Quit]
```

### Stream Player Selection

```text
? Select the player:

▸ Option 1: Default Web Browser (Recommended)
  Option 2: VLC Player (May fail if Cloudflare Challenge is active)

[↑/↓: Navigate • Enter: Select • Esc: Cancel]
```

> **Note:** The default web browser is recommended for seamless playback without direct stream interruption.

---

## Features

- **Interactive Dual-Pane TUI**: Smooth, keyboard-driven interface powered by Charm's [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styled with [Lip Gloss](https://github.com/charmbracelet/lipgloss).
- **Direct CLI Playback**: Bypass the TUI and start playback instantly (`familyguy s05e04`, `familyguy 5 4`, `familyguy S5 E4`).
- **Multi-Player Engine**: Native launch support for Web Browsers and VLC Media Player (Linux, macOS, and Windows detection).
- **Cloud Database Sync**: Automatically loads the latest episode catalog from remote JSON over HTTPS with timeout safeguards.
- **Single Static Binary**: Zero runtime dependencies—compiles down to a fast, standalone executable.
- **Fully Tested**: Table-driven unit tests validating query parsing, case insensitivity, and leading-zero normalization.
- **Automated Releases**: Cross-platform binaries automatically built and published via **GoReleaser** with Homebrew and Scoop support.

---

## Installation

### [Homebrew](https://brew.sh) (macOS & Linux)

```bash
brew install eltaweel068/joy/familyguy
```

*Update anytime with:*
```bash
brew upgrade familyguy
```

---

### [Scoop](https://scoop.sh) (Windows)

```powershell
scoop bucket add toy https://github.com/eltaweel068/toy
scoop install familyguy
```

*Update anytime with:*
```powershell
scoop update familyguy
```

---

### Pre-built Binaries

Download pre-compiled binaries directly from the **[GitHub Releases](https://github.com/eltaweel068/familyguy/releases/latest)** page:

| OS | Architecture | Binary Archive |
| :--- | :--- | :--- |
| Linux | `x86_64` (amd64) | `familyguy_Linux_x86_64.tar.gz` |
| Linux | `arm64` | `familyguy_Linux_arm64.tar.gz` |
| macOS | `Apple Silicon` (arm64) | `familyguy_Darwin_arm64.tar.gz` |
| macOS | `Intel` (x86_64) | `familyguy_Darwin_x86_64.tar.gz` |
| Windows | `64-bit` (x86_64) | `familyguy_Windows_x86_64.zip` |
| Windows | `arm64` | `familyguy_Windows_arm64.zip` |

---

### From Source (Go 1.22+)

**Automated install:**
```bash
go install github.com/eltaweel068/familyguy/cmd/familyguy@latest
```
*(Ensure `$GOPATH/bin` or `~/go/bin` is in your `$PATH`)*

**Manual build:**
```bash
# Clone the repository
git clone https://github.com/eltaweel068/familyguy.git
cd familyguy

# Compile binary
go build -o familyguy ./cmd/familyguy

# Verify build
./familyguy -v
```

---

## Prerequisites

- **Web Browser**: Any modern browser (Default recommended).
- **VLC Media Player** *(Optional)*:
  - Debian / Ubuntu: `sudo apt install vlc`
  - Fedora / RHEL: `sudo dnf install vlc`
  - Arch Linux: `sudo pacman -S vlc`
  - macOS: `brew install --cask vlc`
  - Windows: Download from [videolan.org](https://www.videolan.org/vlc/)

---

## Usage

### Interactive Explorer

Launch the dual-pane interactive browser:

```bash
familyguy
```

### Direct Episode Launch

Pass the season and episode number to stream immediately:

```bash
familyguy s5 e4       # Standard format
familyguy 5 4         # Short notation
familyguy S05 E04     # Uppercase with leading zeros
familyguy s22 e10     # Recent seasons
```

#### Query Format Support:
| Command | Parsed Season | Parsed Episode | Status |
| :--- | :---: | :---: | :---: |
| `familyguy s5 e4` | Season 5 | Episode 4 | Supported |
| `familyguy 5 4` | Season 5 | Episode 4 | Supported |
| `familyguy S05 E04` | Season 5 | Episode 4 | Supported |
| `familyguy s22 e1` | Season 22 | Episode 1 | Supported |

### Version Flag

```bash
familyguy -v
# Output: familyguy v1.0.0 (commit: abc1234, built at: 2026-08-20T12:00:00Z)
```

---

## Keyboard Shortcuts

| Key | Context | Action |
| :--- | :--- | :--- |
| <kbd>↑</kbd> / <kbd>↓</kbd> | Global | Navigate up / down in lists |
| <kbd>→</kbd> / <kbd>Enter</kbd> | Seasons Pane | Focus and open episode list for selected season |
| <kbd>←</kbd> | Episodes Pane | Return focus to seasons list |
| <kbd>Enter</kbd> | Episodes Pane | Select episode & open player selection |
| <kbd>Enter</kbd> | Player Prompt | Launch stream in chosen video player |
| <kbd>Esc</kbd> / <kbd>Ctrl+C</kbd> | Global | Cancel prompt / Quit application |

---

## Project Architecture

```text
familyguy/
├── cmd/
│   └── familyguy/
│       └── main.go              # Entry point, CLI argument routing & runtime execution
├── internal/
│   ├── episode/
│   │   ├── episode.go          # Episode domain models & query normalization algorithm
│   │   └── episode_test.go     # Table-driven unit test suite
│   ├── player/
│   │   ├── player.go           # Cross-platform media launcher (VLC & Browser)
│   │   └── playerMenu.go       # Bubble Tea player selection interactive prompt
│   ├── store/
│   │   └── store.go            # Remote HTTPS JSON database loader
│   └── ui/
│       └── model.go            # Bubble Tea dual-pane TUI model, update & view logic
├── data/
│   └── episodes.json           # Offline reference dataset
├── scraping/
│   └── familyguy_scraping.py   # Python scraper for episode extraction
├── .github/
│   └── workflows/
│       ├── ci.yml              # CI automated test & build check workflow
│       └── release.yml         # CD automated GoReleaser release workflow
├── .goreleaser.yml             # GoReleaser cross-compilation manifest
├── designUI.svg                # Vector UI wireframe
├── go.mod / go.sum             # Go module specifications
├── LICENSE                     # MIT License
└── README.md                   # Project documentation & overview
```

---

## Testing

Run the full unit test suite:

```bash
go test -v ./...
```

Run test coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

---

## Roadmap

- [x] Web scraping pipeline for all 23+ seasons
- [x] Direct CLI command parser (`sX eY`)
- [x] Cross-platform playback launcher (VLC / Browser)
- [x] Dual-pane Bubble Tea interactive TUI
- [x] Remote JSON database synchronization
- [x] Table-driven unit test suite
- [x] GitHub Actions CI pipeline
- [x] Automated multi-architecture CD releases via GoReleaser
- [x] Homebrew tap & Scoop bucket packaging
- [ ] Refactor `internal/ui/model.go` into modular Bubble Tea sub-models and remove legacy commented code
- [ ] Custom theme engine and advanced CSS-like layouts using Charm Lip Gloss
- [ ] Fuzzy title search mode (`familyguy --search "road to"`, `-s`)
- [ ] Random episode picker (`familyguy --random`, `-r`)
- [ ] IMDb ratings & episode descriptions in TUI
- [ ] SSH-hosted TUI using [Wish](https://github.com/charmbracelet/wish) (`ssh familyguy.tv`)

---

## Contributing

Contributions, issues, and feature requests are welcome!

1. **Fork** the repository
2. **Create** your feature branch: `git checkout -b feature/amazing-feature`
3. **Commit** your changes: `git commit -m "feat: add amazing feature"`
4. **Push** to the branch: `git push origin feature/amazing-feature`
5. **Open** a Pull Request

---

## License

Distributed under the **MIT License**. See [`LICENSE`](LICENSE) for more information.

---

## Author

**Mohamed Eltaweel (Meko)**

- GitHub: [@eltaweel068](https://github.com/eltaweel068)

<div align="center">
  <sub>Built with Go • If you find this project fun or useful, please consider starring the repository.</sub>
</div>
