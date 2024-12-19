# Jarvis CLI

Jarvis is a powerful and customizable command-line tool designed to assist with daily workstation activities. This CLI tool is written in Go and offers features like updating applications and more.

## Features

- **Update Applications**: Automatically download and install the latest versions of supported applications like Discord and Chrome.
- **Cross-Platform**: Build and use on Linux systems.
- **Extensible**: Designed for adding more commands and features in the future.

---

## Installation

### Option 1: Install via `go install`

If you have Go installed, you can install Jarvis CLI directly:

```bash
go install github.com/raihankhan/jarvis@latest
```

Ensure your Go binary directory (e.g., `$HOME/go/bin`) is in your `PATH`:

```bash
export PATH=$PATH:$HOME/go/bin
```

Verify the installation:

```bash
jarvis --help
```

### Option 2: Download Precompiled Binary

1. Visit the [Releases](https://github.com/raihankhan/jarvis/releases) page.

2. Download the latest release binary for your platform.

3. Make the binary executable:

   ```bash
   chmod +x jarvis
   ```

4. Move it to a directory in your `PATH`, such as `/usr/local/bin`:

   ```bash
   sudo mv jarvis /usr/local/bin/
   ```

5. Verify installation:

   ```bash
   jarvis --help
   ```

### Option 3: Build from Source

1. Ensure you have Go 1.23 or later installed.

2. Clone this repository:

   ```bash
   git clone https://github.com/raihankhan/jarvis.git
   cd jarvis
   ```

3. Build the binary:

   ```bash
   go build -o jarvis
   ```

4. Move it to a directory in your `PATH`:

   ```bash
   sudo mv jarvis /usr/local/bin/
   ```

5. Verify installation:

   ```bash
   jarvis --help
   ```

---

## Usage

Jarvis CLI offers the following commands:

- **Update Applications**:
    - Update Discord:
      ```bash
      jarvis appupdate discord
      ```
    - Update Chrome:
      ```bash
      jarvis appupdate chrome
      ```

Run `jarvis --help` to see all available commands and options.

---

## Uninstallation

### If Installed via `go install`

1. Remove the binary from the Go binary directory:
   ```bash
   rm $(go env GOPATH)/bin/jarvis
   ```
   Or, if using the default Go binary directory:
   ```bash
   rm $HOME/go/bin/jarvis
   ```

### If Installed Manually

1. Identify the installation path (e.g., `/usr/local/bin`).
2. Remove the binary:
   ```bash
   sudo rm /usr/local/bin/jarvis
   ```

### Verify Removal

Run the following command to ensure `jarvis` is no longer available:

```bash
jarvis --help
```

If it outputs `command not found`, the tool has been successfully uninstalled.

---

## Contributing

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Powered by Go.
- Thanks to the open-source community for inspiration and tools.

