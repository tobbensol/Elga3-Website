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
   npx tailwindcss -i ./src/UI/static/css/input.css -o ./src/UI/static/css/output.css --watch --config .\configs\tailwind.config.js
   ```
4. Host a PostgreSQL database and make a .env file with the connection string

   ```cmd
   echo 'connection_str="your_DB_connection_str_here' > .env
   ```

5. Run the Go server:
   ```cmd
   (cd src && go run ./main/main.go)
   ```
6. Open your web browser and visit http://localhost:8000
