# ELGA3 WEBSITE!

A fun little project where i am finally trying out the full tech stack, from DB to web-server to front-end.

technologies used are:

- GO for web-server
- HTMX for front-end
- Tailwind CSS for styling
- PostgreSQL server
- Gorm as ORM

## Requirements:

- [GO](https://go.dev/dl/)
- [Node.js for npm](https://nodejs.org/en/download)
- [postgreSQL server](https://www.postgresql.org/download/)

## Installation

1. Clone this repository:
   ```cmd
   git clone git@github.com:tobbensol/elga_3_website.git
   cd ELGA_3_WEBSITE
   ```
2. Install Node.js dependencies:
   ```cmd
   npm install
   ```
3. Build Tailwind CSS:
   ```cmd
   (cd configs && npx tailwindcss -i ../src/static/css/input.css -o ../src/static/css/output.css --watch)
   ```
4. Host a PostgreSQL database and make a .env file with the connection string

   ```cmd
   echo 'connection_str="your_DB_connection_str_here' > .env
   ```

5. Change directory to ./src

   ```cmd
   cd src
   ```

6. Run the Go server:
   ```cmd
   go run main.go
   ```
7. Open your web browser and visit http://localhost:8000
