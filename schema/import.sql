-- ------------------------------------------
-- Create the database
-- ------------------------------------------

CREATE DATABASE IF NOT EXISTS markup_scores;

-- ------------------------------------------
-- Create the base schema for the datbase
-- ------------------------------------------

USE markup_scores;

-- ------------------------------------------
-- Create table for 'scores'
-- ------------------------------------------
CREATE TABLE IF NOT EXISTS scores (
    unq_id VARCHAR(25),
    id VARCHAR(25),
    id_date DATE,
    run_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    score INT,
    PRIMARY KEY (unq_id, run_time)
);