# Share-Spreader

Share-Spreader is a powerful and efficient tool designed to help users quickly share and publish content across multiple blogging platforms with ease. Whether you're a blogger, marketer, or content creator, Share-Spreader enables you to expand your reach, save time, and ensure consistency.

## Features

- **Multi-Platform Publishing**: Automatically publish blog posts on platforms like Medium, Dev.to, Hashnode, LinkedIn, and more.
- **Customizable Templates**: Predefine content templates for consistent formatting across platforms.
- **Scheduling**: Schedule posts to go live at specific times on various platforms.
- **User-Friendly Interface**: Simple and clean design for easy navigation and usage.
- **Efficient Backend**: Built with high performance and scalability in mind using microservices architecture.

---

## Tech Stack

- **Golang**: For backend microservices and API integration.
- **gRPC**: For inter-service communication between different services.
- **Plain HTML & CSS**: For a lightweight and responsive frontend.

---

## Getting Started

Follow these steps to get Share-Spreader up and running on your local machine.

### Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/) (Golang programming language)
- [gRPC Tools](https://grpc.io/docs/) (for building and running gRPC services)
- Any modern web browser for frontend testing

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/Share-Spreader.git
   cd Share-Spreader
   ```

2. **Backend Setup**:
   - Navigate to the backend folder:
     ```bash
     cd backend
     ```
   - Install dependencies and build the backend services:
     ```bash
     go mod tidy
     go build
     ```
   - Run the backend services:
     ```bash
     go run main.go
     ```

3. **Frontend Setup**:
   - Navigate to the frontend folder:
     ```bash
     cd frontend
     ```
   - Open `index.html` in your browser to test the frontend interface.

4. **gRPC Service Setup**:
   - Ensure gRPC services are running to enable seamless communication between microservices.

---

## Usage

1. Upload or write your blog content.
2. Choose the platforms you want to publish to (e.g., Medium, Dev.to, LinkedIn).
3. Customize or use pre-built templates for your posts.
4. Schedule the post or publish immediately.
5. Monitor post status and performance analytics (future feature).

---

## Contributing

Contributions are welcome! If you'd like to add new features, fix bugs, or improve the project, follow these steps:

1. **Fork the Project**
2. **Create a Feature Branch**:
   ```bash
   git checkout -b feature/YourFeature
   ```
3. **Commit Your Changes**:
   ```bash
   git commit -m 'Add some feature'
   ```
4. **Push to the Branch**:
   ```bash
   git push origin feature/YourFeature
   ```
5. **Open a Pull Request**

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
