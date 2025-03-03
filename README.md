# Day3-Training-ARC-13523105

This project is a **User and Exam Scores Management System** built with:

- **Backend:** Go Fiber (REST API)
- **Frontend:** HTML, CSS, JavaScript (Vanilla JS)
- **Database:** JSON file storage


## **Table of Contents**

- [Project Overview](#project-overview)
- [Features](#features)
- [Backend Setup](#backend-setup)
- [Frontend Setup](#frontend-setup)
- [How to Use](#how-to-use)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## **Project Overview**

This project provides a user-friendly interface to manage:

- Users
- Exam Scores
- Student Courses

The backend is built using Fiber (Go) to provide a REST API, while the frontend is a lightweight HTML + JavaScript application for interacting with the API.

## **Features**

- User Management (Add, View, Update, Delete)
- Exam Score Management (Add, View, Update, Delete)
- Student Course Mapping
- Frontend UI with Dynamic Updates
- JSON-based Data Storage

## **Backend Setup**

### **Prerequisites**

Ensure you have **Go installed** (v1.18+ recommended).

### **Clone the Repository**

```sh
git clone https://github.com/fathurwithyou/Day3-Training-ARC-13523105.git
cd backend
```

### **Install Dependencies**

```sh
go mod tidy
```

### **Run the API Server**

```sh
go run cmd/Day3-Training-ARC-13523105/main.go
```

Or run `run.bat` for Windows. The backend will start at `http://localhost:3000`.

### **Backend API Routes**

| Method   | Endpoint           | Description                           |
| -------- | ------------------ | ------------------------------------- |
| `GET`    | `/users`           | Get all users                         |
| `POST`   | `/users`           | Add a new user                        |
| `PUT`    | `/users/{id}`      | Update user details                   |
| `DELETE` | `/users/{id}`      | Remove a user                         |
| `GET`    | `/examscores/{id}` | Get exam scores for a user            |
| `POST`   | `/examscores`      | Add a new exam score                  |
| `PUT`    | `/examscores/{id}` | Update an exam score                  |
| `DELETE` | `/examscores/{id}` | Remove an exam score                  |
| `GET`    | `/studentcourses`  | Get users with their enrolled courses |


## **Frontend Setup**

### **Open the Frontend Folder**

```sh
cd frontend
```

### **Open in a Browser**

Open the `index.html` file in your browser, or use **Live Server** if using VS Code.

### **Frontend Features**

- Displays users in a table.
- Clicking a user opens detailed view with courses & exam scores.
- Add User button to create new users.
- Delete User option in the details panel.
- Notifications for success and error messages.

## **How to Use**

### **Viewing Users**

- Open the frontend in a browser.
- The users list will load automatically.
- Click a user row to view detailed information.

### **Adding a User**

- Click the **"Add User"** button.
- Enter the **Name, NIM, and Email**.
- Click **Submit** to add the user.

### **Viewing Exam Scores**

- Click a user to see their exam scores and enrolled courses.

### **Deleting a User**

- Click a user and then click **"Delete User"**.
- Confirm the deletion in the prompt.

---

## **Project Structure**

```
📁 Day3-Training-ARC-13523105
├── 📁 backend
│   ├── 📁 cmd
│   │   ├── 📁 Day3-Training-ARC-13523105
│   │   │   ├── main.go          # Entry point for the API
│   ├── 📁 internal
│   │   ├── 📁 handlers          # API handlers
│   │   ├── 📁 models            # Data models (User, ExamScore, Course)
│   │   ├── 📁 server            # Route setup
│   ├── 📁 data                  # JSON file storage
│   ├── 📄 go.mod                # Go dependencies
│
├── 📁 frontend
│   ├── 📄 index.html            # Main UI
│   ├── 📁 css                   # Styling for UI
│   ├── 📁 js                    # logic
│   
├── 📄 README.md                 # Documentation
```


## **License**

This project is licensed under the **MIT License**.
