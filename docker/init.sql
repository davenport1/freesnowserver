-- Check if the database exists before attempting to create it
SELECT datname FROM pg_database WHERE datname = 'freesnow_db';

-- If the database does not exist, create it
DO $$BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'freesnow_db') THEN
    CREATE DATABASE freesnow_db;
END IF;
END$$;

-- SWITCH TO DATABASE
\c freesnow_db

SET TIMEZONE TO 'UTC';

-- CREATE TABLES
CREATE TABLE IF NOT EXISTS ski_resort (
    id SERIAL PRIMARY KEY,
    resort_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    location geometry,
    timezone VARCHAR(50),
    version INT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS snow_report (
    id SERIAL PRIMARY KEY,
    ski_resort_id INT REFERENCES ski_resort(id),
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    _24_hr INT,
    _72_hr INT,
    week INT
);

CREATE TABLE IF NOT EXISTS trail (
    id SERIAL PRIMARY KEY,
    ski_resort_id INT REFERENCES ski_resort(id),
    name VARCHAR(255) NOT NULL,
    difficulty INT,
    status INT,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lift (
    id SERIAL PRIMARY KEY,
    ski_resort_id INT REFERENCES ski_resort(id),
    name VARCHAR(255) NOT NULL,
    status INT,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS weather_forecast_day(
    id SERIAL PRIMARY KEY,
    ski_resort_id INT REFERENCES ski_resort(id),
    temperature_high INT,
    temperature_low INT,
    wind_direction INT,
    wind_speed_mph INT,
    overcast_level INT,
    humidity_percentage INT,
    temperature_feels_like INT,
    temperature_wind_chill INT,
    sunrise TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    sunset TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- CREATE RELATIONSHIPS
