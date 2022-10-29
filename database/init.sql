CREATE DATABASE IF NOT EXISTS devDB;
USE devDB

CREATE TABLE IF NOT EXISTS Province
(
    ProvinceID INTEGER NOT NULL AUTO_INCREMENT,
    Name VARCHAR(255),
    PRIMARY KEY (ProvinceID)
);

CREATE TABLE IF NOT EXISTS District
(
    DistrictID INTEGER NOT NULL AUTO_INCREMENT,
    Name VARCHAR(255),
    ProvinceID INTEGER,
    CONSTRAINT fk_province FOREIGN KEY (ProvinceID) REFERENCES Province(ProvinceID),
    PRIMARY KEY (DistrictID)
);

CREATE TABLE IF NOT EXISTS Population
(
    CitizenID BIGINT,
    LazerID VARCHAR(255),
    Name VARCHAR(255),
    Lastname VARCHAR(255),
    Birthday DATE,
    Nationality VARCHAR(255),
    DistrictID INTEGER,
    CONSTRAINT fk_district FOREIGN KEY (DistrictID) REFERENCES District(DistrictID),
    PRIMARY KEY (CitizenID)
);

CREATE TABLE IF NOT EXISTS ApplyVote
(
    ID INTEGER NOT NULL AUTO_INCREMENT,
    CitizenID BIGINT,
    CONSTRAINT fk_citizen_apply_vote FOREIGN KEY (CitizenID) REFERENCES Population(CitizenID),
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Candidate
(
    ID INTEGER NOT NULL AUTO_INCREMENT,
    CitizenID BIGINT,
    CONSTRAINT fk_citizen_candidate FOREIGN KEY (CitizenID) REFERENCES Population(CitizenID),
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Mp
(
    ID INTEGER NOT NULL AUTO_INCREMENT,
    CitizenID BIGINT,
    CONSTRAINT fk_citizen_mp FOREIGN KEY (CitizenID) REFERENCES Population(CitizenID),
    PRIMARY KEY (ID)
);