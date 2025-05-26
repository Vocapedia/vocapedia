### DONT FORGET
- [ ] google extension

## Installation

This project uses Docker Compose for backend services and pnpm for the frontend.

### Prerequisites

*   [Docker](https://docs.docker.com/get-docker/)
*   [Docker Compose](https://docs.docker.com/compose/install/)
*   [Node.js](https://nodejs.org/) (for the frontend)
*   [pnpm](https://pnpm.io/) (for the frontend). If you don't have pnpm, you can install it with `npm install -g pnpm`.

### Running the Backend Services

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```
2.  Start the backend services (Go application, PostgreSQL, Redis, Meilisearch) using Docker Compose:
    ```bash
    docker-compose up -d
    ```
3.  The backend API will be accessible at [http://localhost:3000](http://localhost:3000).

### Running the Frontend

1.  Navigate to the frontend directory:
    ```bash
    cd frontend
    ```
2.  Install dependencies:
    ```bash
    pnpm install
    ```
3.  Start the frontend development server:
    ```bash
    pnpm run dev
    ```
4.  The frontend will be accessible at [http://localhost:5173](http://localhost:5173).

## Project Structure

The project is composed of two main components:

*   **Backend:** Written in Go, it handles the application's core logic and API.
*   **Frontend:** Written in Vue.js, it provides the user interface for interacting with the application.

## Dependencies

The project relies on the following external services:

*   **PostgreSQL:** Used as the primary database for storing application data.
*   **Redis:** Used for caching and session management.
*   **Meilisearch:** Used as the search engine for providing fast and relevant search results.