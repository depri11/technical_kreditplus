CREATE TABLE Customer (
    id CHAR(36) PRIMARY KEY,  -- UUID format, 36 characters
    nik VARCHAR(20) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NOT NULL,
    place_of_birth VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    salary DECIMAL(15, 2) NOT NULL,
    photo_ktp VARCHAR(255) NOT NULL,  -- Path to the photo or image URL
    photo_selfie VARCHAR(255) NOT NULL -- Path to the photo or image URL
);