
# 💪 ToFuerte

Welcome to **ToFuerte**! This project is a healthy lifestyle organizer designed to help you organize your routines, meals, and exercises at home. In the future, a gym option will be implemented, where you will be able to complete your daily routine interactively. It will be available for web, Android, and iOS.

## Technologies Used

- ![React Native](https://img.shields.io/badge/React%20Native-20232A?style=for-the-badge&logo=react&logoColor=61DAFB) **React Native** – Framework for building native mobile apps using React.
- ![Expo](https://img.shields.io/badge/Expo-1C1E24?style=for-the-badge&logo=expo&logoColor=white) **Expo** – A platform and set of tools for building React Native apps faster.
- ![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) **Golang** – Used to build the backend of the application.
- ![MariaDB](https://img.shields.io/badge/MariaDB-003545?style=for-the-badge&logo=mariadb&logoColor=white) **MariaDB** – The database management system.

---

## 🚀 Setup Instructions

### 1. **Database Setup**

The database structure can be found in the `db` directory of this project. To load the database into your local MariaDB instance, use:

```bash
mysql -u username -p database_name < path_to_db_file.sql
```

### 2. **Backend Setup**

1. Create a `.env` file with the database credentials:

    ```bash
    usuariodb=user_of_db
    passdb=password_of_db
    ```

2. The backend will run on  `8080` by default. To start the Go server, run:

    ```bash
    go run main/main.go
    ```

### 3. **Frontend Setup (React Native)**

1. You must have expo-CLI installed:

    ```bash
    npm install -g expo-cli
    ```

2. Navigate to the project directory, install the  dependencies, and run:

    ```bash
    npm install
    expo start
    ```

---

## 🌟 Features

- Organize your home workout routines.
- Plan your daily meals to track your diet.
- Simple, responsive interface for easy navigation.
- Available for web, Android, and iOS.

### Upcoming Features

- **Gym Mode**: Complete your daily workout routine interactively at the gym.
- **Mobile integration**: More personalized tracking with functions available on both Android and iOS.

---

## 📫 Contributing

If you'd like to contribute to **ToFuerte**, feel free to open issues or submit pull requests. Any feedback and suggestions are always welcome!

---

