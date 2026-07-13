# 🔗 URL Shortener

A simple and efficient URL Shortener built using **Go**, **PostgreSQL**, **HTML**, and **CSS**. The application allows users to shorten long URLs, store them in a PostgreSQL database, redirect using the generated short URL, and track the number of clicks.

---

## 🚀 Features

- ✅ Shorten long URLs
- ✅ Generate unique short codes
- ✅ Redirect to the original URL
- ✅ Store URLs in PostgreSQL
- ✅ Track click count
- ✅ Responsive HTML & CSS interface
- ✅ Environment variables using `.env`

---

## 🛠️ Tech Stack

- Go (Golang)
- PostgreSQL
- Gorilla Mux
- HTML5
- CSS3
- Git
- Docker (Ready)

---

## 📂 Project Structure

```
url-shortener/
│
├── database/
│   └── db.go
│
├── handlers/
│   └── handlers.go
│
├── static/
│   └── style.css
│
├── templates/
│   └── index.html
│
├── utils/
│   └── generator.go
│
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## ⚙️ Installation

### Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/url-shortener.git
```

### Go to the project

```bash
cd url-shortener
```

### Install dependencies

```bash
go mod tidy
```

### Configure environment variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=url_shortener
```

### Run the project

```bash
go run main.go
```

Open:

```
http://localhost:8081
```

---

## 🗄️ Database Schema

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    clicks INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 📸 Screenshots

- Home Page
- Short URL Generated
- PostgreSQL Database

(Add screenshots later.)

---

## 📈 Future Improvements

- Copy to Clipboard button
- URL validation
- QR Code generation
- User authentication
- Custom short URLs

---

## 👩‍💻 Author

**Shaik Farhin**

GitHub: https://github.com/YOUR_USERNAME