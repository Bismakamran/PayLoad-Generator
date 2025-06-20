# 🔐 Payload Encoder & Obfuscator (Go)

This is a standalone **Go-based utility** that performs advanced encoding and obfuscation techniques commonly used in web payloads to test WAFs, filters, and input validation.

---

## ✨ Features

### 🔸 Encoders
Convert any input string into:
- ✅ **Base64**
- ✅ **URL Encoding**
- ✅ **Hex Encoding** (`\x41\x42`)
- ✅ **Unicode Escaping** (`\u0041\u0042`)

### 🔸 Obfuscators
Modify payloads to evade filters using:
- 🔀 **Random spacing** between characters
- 💬 **HTML comment injection**
- 🆎 **Case flipping** (e.g. `<ScRiPt>`)
- 🧙 **Homoglyph replacement** (Unicode look-alikes)

---

## 🧪 Example Output


Output:
Base64: PHNjcmlwdD5hbGVydCgxKTwvc2NyaXB0Pg==
URL Encoded: %3Cscript%3Ealert%281%29%3C%2Fscript%3E
Hex Encoded: \x3c\x73\x63\x72\x69\x70\x74...
Unicode: \u003C\u0073\u0063\u0072...

Random Spacing: < s c r i p t > a l e r t ( 1 )
With Comments: <script><!-- -->alert(1)</script><!-- -->
Case Flipped: <ScRiPt>ALert(1)</ScRipT>
Homoglyphs: <ѕсrіpt>аⅼеrt(1)</ѕсrіpt>



---

## 📁 Folder Structure

payload-generator-go/
│
├── main.go # Test runner for encoder/obfuscator
├── go.mod # Go module config
│
└── utils/
├── encoder.go # Encoding functions
└── obfuscator.go # Obfuscation functions



---

## 🚀 Getting Started

### 1. Clone the Repo or Copy Files
Make sure `main.go`, `utils/encoder.go`, and `utils/obfuscator.go` are in place.

### 2. Initialize the Module
```bash
go mod init payload-generator-go

## Run the Tool

go run main.go
