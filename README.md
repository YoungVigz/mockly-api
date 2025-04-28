# Mockly API 🚀

Mockly API enables the generation and management of mock data using user-provided JSON schemas. The API processes these schemas using the **Mockly** CLI (see [Mockly.js GitHub](https://github.com/YoungVigz/mockly-cli)) and delivers generated data along with real-time notifications via WebSockets.


## Key Functionalities ✨

- **CRUD for Data Schemas**:  
  Create, read, update, and delete data schemas used for data generation. This allows users to manage their schemas easily.

- **Data Generation via CLI**:  
  Trigger data generation by processing a submitted schema. The Mockly.js CLI tool is used for generating the data.
- **Real-time Notifications**:  
  Receive live updates on the generation progress, errors, and operation status through WebSocket connections. This provides a smooth user experience as they monitor data generation in real time.

- **JWT Authentication**:  
  Secure endpoints for data generation and schema management using JWT tokens for user authentication.


## Requirements 🛠️

- **Go** (1.24.1)
- **Node.js** (for the CLI tool, Mockly.js)
- **Docker**


## Installation & Running (dev) 🚀

```bash
    git clone https://github.com/YoungVigz/mockly-api.git
    cd mockly-api
    go mod tidy
    docker compose build
    docker compose up --watch
```

You can change provided env in docker-compose.yml file

# 📖 Documentation

After running api localy you can check [swagger docs](http://localhost:8080/api/docs/index.html) 


# Author
[@Gabriel Gałęza](https://github.com/YoungVigz)
