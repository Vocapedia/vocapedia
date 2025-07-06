# Firstly vibe coding then debugging

## Installation

This project uses Docker Compose for backend services and pnpm for the frontend. The application consists of a Go backend with PostgreSQL, Redis, and Meilisearch services, and a Vue.js frontend.

### Prerequisites

*   [Docker](https://docs.docker.com/get-docker/)
*   [Docker Compose](https://docs.docker.com/compose/install/)
*   [Node.js](https://nodejs.org/) (v18 or higher)
*   [pnpm](https://pnpm.io/) (for the frontend). If you don't have pnpm, you can install it with `npm install -g pnpm`.
*   [Go](https://golang.org/) (v1.23 or higher) - for local development
*   [Make](https://www.gnu.org/software/make/) - for running makefile commands

### Quick Start (Separate Frontend & Backend)

1.  Clone the repository:
    ```bash
    git clone https://github.com/Vocapedia/vocapedia
    cd vocapedia
    ```

2.  Start the backend services (API + databases):
    ```bash
    make docker-up
    ```

3.  Start the frontend development server:
    ```bash
    cd frontend
    pnpm install
    pnpm run dev
    ```

4.  The application will be available at:
    - **Frontend**: [http://localhost:5173](http://localhost:5173)
    - **Backend API**: [http://localhost:3000](http://localhost:3000)
    - PostgreSQL: [`localhost:5432`](http://localhost:5432)
    - Redis: [`localhost:6379`](http://localhost:6379)
    - Meilisearch: [http://localhost:7700](http://localhost:7700)

### Development Setup

#### Option 1: Hot Reload Development (Recommended)

For active development with hot reload:

1.  Start development backend services:
    ```bash
    make hot
    ```

This will start:
- PostgreSQL database
- Redis cache  
- Meilisearch search engine
- Go backend with hot reload (using Air)

2.  In a separate terminal, start the frontend development server:
    ```bash
    cd frontend
    make dev
    # or alternatively: pnpm run dev
    ```

3.  Access the application:
    - **Frontend**: [http://localhost:5173](http://localhost:5173)
    - **Backend API**: [http://localhost:3000](http://localhost:3000)

#### Option 2: Manual Development Setup

1.  Start only the infrastructure services:
    ```bash
    docker-compose -f docker-compose.hot.yaml up vocapedia-psql vocapedia-redis meilisearch
    ```

2.  Run the backend locally:
    ```bash
    cd backend
    make local
    ```

3.  Run the frontend locally:
    ```bash
    cd frontend
    pnpm install
    pnpm run dev
    ```

### Testing Environment

To run tests with isolated test databases:

1.  Start test services:
    ```bash
    make test
    ```

This provides:
- Test PostgreSQL: `localhost:5433`
- Test Redis: `localhost:6380`
- Test Meilisearch: `localhost:7701`
- MailHog (email testing): [http://localhost:8025](http://localhost:8025)

### Production Deployment

For production deployment, you need to run both frontend and backend separately:

#### Backend Production

1.  Start the backend services:
    ```bash
    docker-compose up -d
    ```

#### Frontend Production

1.  Build the frontend for production:
    ```bash
    cd frontend
    pnpm install
    pnpm run build
    ```

2.  Serve the built frontend using a web server (nginx, apache, etc.) or:
    ```bash
    cd frontend
    pnpm run preview
    ```

**Note**: The frontend and backend run as separate services. Make sure to configure your frontend to point to the correct backend API URL in production.

### Available Services

The application stack includes:

- **Backend (Go)**: REST API server with gRPC support (runs on port 3000)
- **Frontend (Vue.js)**: Progressive web application (runs on port 5173 in development)
- **PostgreSQL**: Primary database with PGroonga extension for full-text search
- **Redis**: Caching and session management
- **Meilisearch**: Fast search engine for content search

**Architecture**: The frontend and backend are decoupled services that communicate via REST API.

### Database Configuration

Default database credentials:
- **Username**: `vocapedia`
- **Password**: `vocapedia`
- **Database**: `vocapedia`

For testing:
- **Username**: `vocapedia_test`
- **Password**: `vocapedia_test`
- **Database**: `vocapedia_test`

## Project Structure

The project is composed of two main components:

*   **Backend:** Written in Go, it handles the application's core logic and API.
*   **Frontend:** Written in Vue.js, it provides the user interface for interacting with the application.

## Dependencies

The project relies on the following external services:

*   **PostgreSQL:** Used as the primary database for storing application data.
*   **Redis:** Used for caching and session management.
*   **Meilisearch:** Used as the search engine for providing fast and relevant search results.