-- CREATE DATABASE
CREATE DATABASE freesnow_db;

-- CREATE TABLES
CREATE TABLE ski_resort (
    id SERIAL PRIMARY KEY,
    resort_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timezone VARCHAR(50),
    version INT
);

CREATE TABLE snow_report (
    id SERIAL PRIMARY KEY,
    ski_resort_id INT,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    _24_hr INT,
    _72_hr INT,
    week INT
);

CREATE TABLE trail (
    id INT,
    ski_resort_id INT,
    name VARCHAR(255) NOT NULL,
    difficulty INT,
    status INT,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE lift (
    id INT,
    ski_resort_id INT,
    name VARCHAR(255) NOT NULL,
    status INT,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE week_forecast(
    id
)

CREATE TABLE day_forecast(

)

-- CREATE RELATIONSHIPS
