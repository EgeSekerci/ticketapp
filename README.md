<h1 align="center">Welcome to Ticket IT 👋</h1>
<p>
<a href="https://tip.golang.org/doc/go1.22" target="_blank">
    <img alt="Go 1.22.5" src="https://custom-icon-badges.demolab.com/badge/Go-1.22.5-cyan.svg?logo=go" />
  </a>
  <a href="https://htmx.org/docs/" target="_blank">
    <img alt="HTMX 2.0.2" src="https://custom-icon-badges.demolab.com/badge/HTMX-2.0.2-blue.svg?logo=htmx" />
  </a>
  <a href="https://github.com/EgeSekerci/ticketapp/blob/main/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> Minimal and fast ticketing service to quickly let your IT department know the issues you are experiencing.

## Prerequisites

- [Go](https://go.dev/dl/) (v1.22 or newer)
- [PostgreSQL](https://www.postgresql.org/)
- [TailwindCSS Standalone CLI](https://tailwindcss.com/blog/standalone-cli)

## Install

```sh
git clone https://github.com/EgeSekerci/ticketapp.git

cd ticketapp/

go mod tidy
```
## Database Setup

Make sure PostgreSQL is running, then:

```bash
# Create the database
CREATE DATABASE ticketapp;

# Run the schema to create tables
psql -U your_username -d ticketapp -f db/schema.sql
```

## Usage
```sh
go run .
```

## Run tests

```sh
cd test/
go test .
```

## Author

👤 **Yiğit Ege Şekerci**

* Github: [@EgeSekerci](https://github.com/EgeSekerci)
* LinkedIn: [@yegesekerci](https://linkedin.com/in/yegesekerci)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/EgeSekerci/ticketapp/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2025 [Ege Şekerci](https://github.com/EgeSekerci).<br />
This project is [MIT](https://github.com/EgeSekerci/ticketapp/blob/main/LICENSE) licensed.
