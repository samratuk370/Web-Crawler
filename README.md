# Web Crawler Dashboard

A full-stack app to analyze website URLs by crawling and extracting page metadata.

## ✨ Features

- Add URLs and initiate analysis
- Crawl HTML to extract:
  - HTML version
  - Page title
  - Heading structure
  - Internal vs external links
  - Broken links (4xx/5xx)
  - Login form detection
- Real-time crawl status updates
- Dashboard view with filters, sorting, and pagination
- Detailed report view with charts

---

## 🧱 Tech Stack

| Layer      | Tech                          |
|------------|-------------------------------|
| Frontend   | React, TypeScript, Tailwind   |
| Backend    | Go (Gin), MySQL               |
| Crawling   | net/http, goquery             |
| Charts     | Recharts                      |
| DevOps     | Docker, Docker Compose        |

---

## 🚀 Run with Docker

```bash
git clone https://github.com/your-username/web-crawler.git
cd web-crawler
docker-compose up --build
```

- Frontend: http://localhost:4173
- Backend API: http://localhost:8080/api
- DB: MySQL at port 3306

---

## 🧪 Tests

- Frontend includes basic happy-path tests (`react-testing-library`)
- Run locally:

```bash
cd frontend
npm install
npm run test
```

---

## 🧩 API Endpoints

| Method | Endpoint            | Description         |
|--------|---------------------|---------------------|
| GET    | `/api/urls`         | List URLs           |
| POST   | `/api/urls`         | Add new URL         |
| POST   | `/api/urls/:id/start` | Start crawl        |
| GET    | `/api/urls/:id`     | URL details         |

> All requests must include:
>
> `Authorization: Bearer secret-token`

---

## 🧠 Author

Built with ❤️ for Sykell test task  
[GitHub Profile](https://github.com/samratuk370)
