# 🎨 ASCII-Art

> A Go program that converts text into beautiful ASCII graphical banners.

---

## 📖 Description

**ASCII-Art** is a Go application that receives a string as a command-line argument and prints it in a graphical ASCII format.

The program reads predefined banner templates and renders each character as an 8-line ASCII representation.

Each character:

* Has a height of **8 lines**
* Is generated using a selected **banner style**
* Supports letters, numbers, spaces, special characters, and `\n`

---

## 🗂 Project Structure

```
ascii-art/
│
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── funcs/
│   ├── helper.go
│   ├── printer.go
│   └── splitendl.go
│
├── main.go
└── go.mod
```

---

## 🚀 Usage

### ▶ Run with Default Banner (`standard.txt`)

```
go run main.go "any_text" | cat -e
```

### Example

```
go run main.go "Hello" | cat -e
```

---

### 🎭 Run with Specific Banner Style

```
go run main.go "text" "banner_name" | cat -e
```

### Available Banner Styles

* `standard.txt`
* `shadow.txt`
* `thinkertoy.txt`

### Example

```
go run main.go "Hello" "shadow.txt" | cat -e
```

---

## ✨ Features

✔ Supports:

* Uppercase and lowercase letters
* Numbers
* Spaces
* Special characters
* `\n` for new lines

✔ Handles multiple lines correctly

✔ Uses only **standard Go packages**

✔ Clean project structure

---

## 🖥 Example with New Line

```
go run main.go "Hello\nWorld" | cat -e
```

---

## ⚙ Requirements

* Go installed
* Only standard Go packages are used
* Banner files must be inside the `banners/` directory

---

## 📌 Notes

* If no banner is specified, `standard.txt` is used by default.
* Input must be passed as a command-line argument.
* The program reads banner files from the `banners/` folder.
