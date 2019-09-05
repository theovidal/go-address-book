# Address book

**A simple learning project written in Golang**

[![License Badge](https://img.shields.io/github/license/exybore/go-address-book)](#-license)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/7305d25d08c14de2a780cb8c1cbca0c2)](https://www.codacy.com/app/exybore/go-address-book?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=exybore/go-address-book&amp;utm_campaign=Badge_Grade)

With this address book you can do simple stuff like listing, creating, updating and deleting contacts from the book. It can be done through the command-line interface (CLI), and soon through a graphical user interface (GUI).

All the data is stored in a single .yml file. In the future, it'll be encrypted.

**Disclaimer : this projet isn't really aimed to be used as an every day address book. It's just a small project that I made in order to learn the language and how to write clean code.**

- [⌨ Usage](#-usage)
- [💻 Contribute](#-contribute)
- [📜 Credits](#-credits)
- [🔐 License](#-license)

## ⌨ Usage

Before continuing, make sure you have Go 1.12 or earlier installed on your system.

Clone the address book on your computer :

```bash
git clone https://github.com/exybore/go-address-book.git  # HTTP
git clone git@github.com:exybore/go-address-book          # SSH
```

Edit the configuration through the `config.toml` file. Every option and possible choice is explained, and the syntax is very user-friendly.

If you want to quickly test the application, use `go run main.go`. Otherwise, use `go build` to compile into an executable file.

## 💻 Contribute

Thanks for being interested in the development of this small open-source project ! Please consider these simple but required steps :

- [Fork the repository](https://github.com/exybore/go-address-book/fork)
- Clone it locally using the method shown above
- Make all the changes you want. Please use [this standard format] for commit messages.
- Push the repository and open a new pull request in which you explain why you made these changes

## 📜 Credits

- Maintainer : [Exybore](https://github.com/exybore/go-address-book)

## 🔐 License

MIT License

Copyright (c) 2019 Exybore

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
