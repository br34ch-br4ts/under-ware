# Ransomware Project - AES-256 Encryption

## 📜 Overview
This project is a ransomware written in **Go** (Golang), designed to encrypt files using **AES-256 encryption**. It demonstrates how to create a program with functionalities such as file encryption, desktop modification, and ransom note generation.

### ⚠️ Disclaimer
This project is intended for **educational purposes only** to help understand the mechanisms of ransomware and enhance cybersecurity measures. Misuse of this code for malicious purposes is illegal and unethical. Use responsibly.

## 🚀 Features
1. **File Encryption with AES-256**
   - Encrypts all files in target directories using the Advanced Encryption Standard (AES) with a 256-bit key.
   - Efficient recursive encryption ensures maximum impact.

2. **Desktop Background Change**
   - Automatically changes the victim's desktop wallpaper to a custom image with instructions or warnings.

3. **Ransom Note Generation**
   - Creates a text file with details on:
     - Instructions to download and install Tor Browser.
     - Navigating to a specific .onion site for payment and decryption software.

4. **Tor Integration**
   - Provides the victim with essential information to access the darknet for further instructions.

## 🔧 Technical Details

### File Encryption
- **Algorithm:** AES-256 in GCM mode (ensures data integrity and confidentiality).
- **Key Handling:** A unique key is generated for each instance, stored securely for decryption upon payment.
- **Targets:** Recursively scans and encrypts files in user-defined directories (e.g., `/home/user` on Linux or `C:\Users` on Windows).
- **Exclusions:** Skips critical system files to avoid rendering the machine unbootable.

### Desktop Background Change
- Replaces the victim's current wallpaper with a custom image provided in the payload.
- Supports common image formats such as `.jpg`, `.png`, and `.bmp`.

### Ransom Note
- Automatically generates a file named `READ_ME.txt` or equivalent in prominent directories (e.g., Desktop, Documents).
- Includes:
  - Steps to download Tor Browser.
  - The onion URL for the ransom page.
  - Instructions for payment and decryption.

### Written in Go
- **Why Go?**
  - Cross-platform compatibility.
  - Strong support for concurrency and networking.
  - Lightweight binaries for easier distribution.

## 📂 Directory Structure
```
├── main.go              # Main program logic
├── encrypt.go           # Encryption functionality
├── wallpaper.go         # Wallpaper modification logic
├── ransom_note.go       # Ransom note generation
├── assets/
│   ├── wallpaper.png   # Custom wallpaper image
│   └── note_template.txt  # Template for ransom note
├── README.md            # Project documentation
```

## 🛠️ Installation & Usage

### Prerequisites
- **Go (Golang):** Install the Go compiler from [golang.org](https://golang.org).
- **Dependencies:**
  - No external dependencies; standard Go libraries are used.

### Build Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/ransomware-project.git
   cd ransomware-project
   ```

2. Build the binary:
   ```bash
   go build -o ransomware
   ```

3. Run the program:
   ```bash
   ./ransomware
   ```

## 🛡️ Legal and Ethical Use
This project must only be used for **research** and **cybersecurity training**. Unauthorized deployment of this program is strictly prohibited and punishable under applicable laws.

---

### 📫 Contact
For questions or collaborations, please reach out to:
- Email: your_email@example.com
- GitHub: [your_username](https://github.com/your_username)

---

**Stay ethical and use this knowledge responsibly!**
