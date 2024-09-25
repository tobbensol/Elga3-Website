# ELGA3 WEBSITE!

A fun little project where i am finally trying out the full tech stack, from DB to web-server to front-end.

technologies used are:

- Go: Backend web server
- HTMX: Front-end interactivity
- Tailwind CSS: Styling framework
- Templ: Templating engine for HTML generation
- PostgreSQL: Database
- Gorm: ORM (Object Relational Mapping) for database interaction

## Requirements:
- [Go](https://go.dev/dl/) (for building and running the backend)
- [Node.js](https://nodejs.org/en/download) (for npm to build Tailwind CSS)
- [PostgreSQL](https://www.postgresql.org/download) (database server)

## Installation
1. Clone this repository:
   ```bash
   git clone git@github.com:tobbensol/elga_3_website.git
   cd ELGA_3_WEBSITE
   ```
2. Install Node.js dependencies for tailwind:
   ```bash
   npm install
   ```
3. Enter the working directory
   ```bash
   cd src
   ```

4. Initialize tailwindCSS and Templ LSP:
   ```bash
   cmd /c "npx tailwindcss -i ./UI/static/css/input.css -o ./UI/static/css/output.css --watch --config ..\tailwind.config.js && templ lsp"
   ```

4. Set up and host a PostgreSQL database on your local machine or a remote server.

5. Obtain an OAuth 2.0 Client ID and Secret from Discord for authentication.

6. Put the PostgreSQL connection string,  discord Client ID and Secret into a .env file
   ```cmd
   CONNECTION_STR="your_postgres_connection_string_here"
   CLIENT_ID="your_discord_client_id_here"
   CLIENT_SECRET="your_discord_client_secret_here"
   ```

7. Run the Go server:
   ```bash
   (go generate && go run .)
   ```
8. Open your web browser and visit http://localhost:8000
