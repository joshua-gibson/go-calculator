# Go Calculator

A simple command-line and graphical calculator written in Go.

This project demonstrates a basic Go application with a modular structure, unit tests, a graphical interface using Fyne, and a multi-platform build and release pipeline using GitHub Actions.

## Installation

You can download the latest pre-built executables for Linux, macOS, and Windows from the [GitHub Releases page](https://github.com/joshua-gibson/go-calculator/releases).

---

### Linux

1.  Download the `calc-ubuntu-latest` (CLI) and `gcalc-ubuntu-latest` (GUI) files.
2.  Open your terminal and navigate to the directory where you downloaded the files.
3.  Make the files executable:
    ```bash
    chmod +x calc-ubuntu-latest
    chmod +x gcalc-ubuntu-latest
    ```
4.  (Optional) Rename the files for convenience:
    ```bash
    mv calc-ubuntu-latest calc
    mv gcalc-ubuntu-latest gcalc
    ```
5.  (Optional) Move the files to a directory in your PATH to make them accessible from anywhere:
    ```bash
    sudo mv calc /usr/local/bin/
    sudo mv gcalc /usr/local/bin/
    ```

---

### macOS

1.  Download the `calc-macos-latest` (CLI) and `gcalc-macos-latest` (GUI) files.
2.  Open your terminal and navigate to the directory where you downloaded the files.
3.  Make the files executable:
    ```bash
    chmod +x calc-macos-latest
    chmod +x gcalc-macos-latest
    ```
4.  **For the GUI App (`gcalc-macos-latest`):** Because the application is not signed by an Apple Developer, you may need to bypass Gatekeeper the first time you run it.
    *   Right-click (or Ctrl-click) the `gcalc-macos-latest` file and select "Open".
    *   You will see a warning dialog. Click the "Open" button to run the application. You will only need to do this once.

5.  (Optional) Rename and move the files as described in the Linux instructions.

---

### Windows

1.  Download the `calc-windows-latest.exe` (CLI) and `gcalc-windows-latest.exe` (GUI) files.
2.  **For the GUI App (`gcalc-windows-latest.exe`):** Windows Defender SmartScreen may show a warning because the application is not signed.
    *   Click "More info".
    *   Click "Run anyway" to start the calculator.
3.  You can run the executables directly from the Command Prompt (`cmd.exe`) or PowerShell.
4.  (Optional) Move the `.exe` files to a folder that is included in your system's PATH environment variable to make them accessible from anywhere.

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
