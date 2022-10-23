CREATE DATABASE devDB;
USE devDB

CREATE TABLE Province
(
    ProvinceID INTEGER,
    ProvinceName VARCHAR(255),
    PRIMARY KEY (ProvinceID)
);

CREATE TABLE District
(
    DistrictID INTEGER,
    ProvinceID INTEGER,
    CONSTRAINT fk_province FOREIGN KEY (ProvinceID) REFERENCES Province(ProvinceID),
    PRIMARY KEY (DistrictID)
);

CREATE TABLE Population
(
    CitizenID INTEGER,
    LazerID INTEGER,
    Name VARCHAR(255),
    Lastname VARCHAR(255),
    Birthday DATE,
    Nationality VARCHAR(255),
    DistrictID INTEGER,
    CONSTRAINT fk_district FOREIGN KEY (DistrictID) REFERENCES District(DistrictID),
    PRIMARY KEY (CitizenID)
);

CREATE TABLE ApplyVote
(
    UUID INTEGER,
    CitizenID INTEGER,
    CONSTRAINT fk_citizen FOREIGN KEY (CitizenID) REFERENCES Population(CitizenID),
    PRIMARY KEY (UUID)
);