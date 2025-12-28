# Go Calculator

A simple command-line and graphical calculator written in Go.

This project demonstrates a basic Go application with a modular structure, unit tests, a graphical interface using Fyne, and a multi-platform build and release pipeline using GitHub Actions.

## Installation

You can download the latest pre-built executables for your operating system from the [GitHub Releases page](https://github.com/joshua-gibson/go-calculator/releases).

---

### Linux

The Linux binaries are built on Ubuntu but are generic and should work on most modern Linux distributions (like Fedora, Debian, Arch, etc.).

1.  Download the `calc-linux-amd64` (CLI) and `gcalc-linux-amd64` (GUI) files.
2.  Open your terminal and navigate to the download directory.
3.  Make the files executable:
    ```bash
    chmod +x calc-linux-amd64 gcalc-linux-amd64
    ```
4.  (Optional) Rename and move the files into your PATH for easy access:
    ```bash
    sudo mv calc-linux-amd64 /usr/local/bin/calc
    sudo mv gcalc-linux-amd64 /usr/local/bin/gcalc
    ```

---

### macOS

1.  Download the `calc-macos-amd64` (CLI) and `gcalc-macos-amd64` (GUI) files.
2.  Open your terminal and navigate to the download directory.
3.  Make the files executable:
    ```bash
    chmod +x calc-macos-amd64 gcalc-macos-amd64
    ```
4.  **Important:** Because the application is not signed by an Apple Developer, you must bypass Gatekeeper the first time you run it.
    *   Right-click (or Ctrl-click) the `gcalc-macos-amd64` file and select "Open".
    *   You will see a warning dialog. Click the "Open" button to run the application. You only need to do this once.

---

### Windows

1.  Download the `calc-windows-amd64.exe` (CLI) and `gcalc-windows-amd64.exe` (GUI) files.
2.  **Important:** Windows Defender SmartScreen may show a warning because the application is not signed.
    *   Click "More info".
    *   Click "Run anyway" to start the calculator.
3.  You can run the executables directly from Command Prompt or PowerShell. To make them accessible from anywhere, move them to a folder that is included in your system's PATH environment variable.

## Usage

### Command-Line Interface (CLI)

Run the `calc` executable with two numbers and an operator (`+`, `-`, `*`, `/`).

```bash
./calc 10 * 5
# Output: 50

./calc 100 / 2.5
# Output: 40
```

### Graphical User Interface (GUI)

Run the `gcalc` executable to open the graphical calculator window.

```bash
./gcalc
```

## Building from Source

If you have Go installed, you can build the calculator yourself.

1.  Clone the repository:
    ```bash
    git clone https://github.com/joshua-gibson/go-calculator.git
    cd go-calculator
    ```
2.  Build the CLI:
    ```bash
    go build -o calc ./cmd/calc
    ```
3.  Build the GUI:
    ```bash
    go build -o gcalc ./cmd/gcalc
    ```
