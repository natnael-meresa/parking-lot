CREATE TABLE IF NOT EXISTS cars (
    id INT AUTO_INCREMENT PRIMARY KEY,
    model VARCHAR(50) NOT NULL,
    registration VARCHAR(10) UNIQUE NOT NULL,
    mileage INT NOT NULL,
    available ENUM('available', 'rented') DEFAULT 'available'
);