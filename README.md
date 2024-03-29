# Chat-Asisten

Code Asisten adalah sebuah aplikasi asisten virtual yang bertujuan untuk memberikan bantuan kepada pengguna dengan menjawab pertanyaan-pertanyaan yang diajukan. Aplikasi ini dirancang untuk memberikan dukungan dan solusi atas pertanyaan yang terkait dengan pemrograman, pengembangan perangkat lunak, bahasa pemrograman, serta topik teknis lainnya.

## Getting Started

### Prerequisites

- Node.js and npm installed on your machine.
- Golang installed on your machine.

### Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/InnakaDylee/Chat-Asisten.git
    cd chat-asisten
    ```

2. **Client Setup (Next.js)**

    ```bash
    cd client
    npm install
    ```

3. **Server Setup (Golang with Echo)**

    ```bash
    cd server
    # Set up your Go environment and dependencies
    go mod download
    ```

### Running the Applications

1. **Start the Next.js Client**

    Inside the `client` directory:

    ```bash
    npm run dev
    ```

    Access the client application via `http://localhost:3000`.

2. **Run the Echo Golang Server**

    Inside the `server` directory:

    ```bash
    go run main.go
    ```

    The server will start at `http://localhost:8080`.

## Folder Structure

- **client**: Contains the Next.js client-side application.
- **server**: Contains the Golang server using Echo framework.

## Additional Notes

- You can customize the functionality, endpoints, and UI elements as per your project requirements.
- Refer to the specific documentation of Next.js and Echo framework for advanced usage and customization.

## Picture
-![Image](./docs/Screenshot%202024-01-02%20092404.png)

## Contributors

- [@InnakaDylee](https://github.com/InnakaDylee) - Innaka Dylee

