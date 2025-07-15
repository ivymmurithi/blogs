# Blog API (Go + PostgreSQL)

The backend for the blog app, built with Go (no frameworks) and PostgreSQL.

## Setup

1. Setup DB_URL and port in your .env file. Check .env.example file inside the blogs-backend directory to see how the .env file should look like.
2.Create table in psql
   ```bash
     CREATE TABLE blogs (
      id SERIAL PRIMARY KEY,
      title TEXT NOT NULL,
      content TEXT NOT NULL,
      slug TEXT UNIQUE NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      );
    ```
3.  Add dummy blogs to the database.
    ```bash
     psql -U youruser -d blogs -f seed.sql
     ```
4. Run the server:

```bash
go run main.go db.go handlers.go models.go
```

#Frontend

# 🎨 Blog Frontend (Angular + Tailwind CSS + Material)

Frontend for the blog app, built with Angular and styled using Tailwind CSS and Angular Material.


## Setup

1. Use Node version 18:

```bash
nvm use 18
npm install
ng serve
```

<img width="1907" height="865" alt="Screenshot from 2025-07-16 00-01-28" src="https://github.com/user-attachments/assets/413ce8ee-cf2f-4ba5-bd4c-5efa64af0ab9" />
<img width="1907" height="865" alt="Screenshot from 2025-07-16 00-01-40" src="https://github.com/user-attachments/assets/b1fd6ed5-e1c3-4450-8fb2-ff7b855ecf6e" />

