-- Table: users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    last_login TIMESTAMP NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: roles
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role VARCHAR(155) NOT NULL
);

-- Data: roles
INSERT INTO roles (role)
VALUES
    ('ADMIN'),
    ('USER'),
    ('MODERATOR');


CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Table: company_categories
CREATE TABLE company_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NULL
);

-- Data: company_categories
INSERT INTO company_categories (name, description)
VALUES
    ('Technology', 'Products and services related to technology'),
    ('Health', 'Healthcare products, services, and wellness'),
    ('Education', 'Educational materials, courses, and resources'),
    ('Fashion', 'Apparel, clothing, and fashion accessories'),
    ('Food & Beverage', 'Food and drink products, restaurants, and catering services'),
    ('Finance', 'Banking, investment, and financial services'),
    ('Entertainment', 'Movies, music, gaming, and other forms of entertainment'),
    ('Sports', 'Sports equipment, activities, and events'),
    ('Automotive', 'Vehicles, car parts, and automotive services'),
    ('Home & Garden', 'Furniture, home decor, and gardening products');


-- Table: companies
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cnpj VARCHAR(20) UNIQUE NOT NULL,
    address VARCHAR(255) NULL,
    email VARCHAR(255) NULL,
    phone VARCHAR(20) NULL,
    category_id INT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_company_category FOREIGN KEY (category_id) REFERENCES company_categories(id) ON DELETE SET NULL
);

-- Table: complaints
CREATE TABLE complaints (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    company_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_complaint_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_complaint_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);

-- Table: responses
CREATE TABLE responses (
    id SERIAL PRIMARY KEY,
    complaint_id INT NOT NULL,
    user_id INT NULL,
    company_id INT NULL,
    content TEXT NOT NULL,
    company_response BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_response_complaint FOREIGN KEY (complaint_id) REFERENCES complaints(id) ON DELETE CASCADE,
    CONSTRAINT fk_response_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_response_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL
);

-- Table: categories
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NULL
);

-- Table: complaint_categories
CREATE TABLE complaint_categories (
    complaint_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (complaint_id, category_id),
    CONSTRAINT fk_complaint_categories_complaint FOREIGN KEY (complaint_id) REFERENCES complaints(id) ON DELETE CASCADE,
    CONSTRAINT fk_complaint_categories_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- Table: ratings
CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    complaint_id INT NOT NULL,
    user_id INT NOT NULL,
    rating_value INT NOT NULL CHECK (rating_value BETWEEN 1 AND 5),
    comment TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_rating_complaint FOREIGN KEY (complaint_id) REFERENCES complaints(id) ON DELETE CASCADE,
    CONSTRAINT fk_rating_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
