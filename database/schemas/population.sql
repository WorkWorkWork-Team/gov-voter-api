CREATE TABLE population
(
    CitizenID INTEGER,
    LazerID INTEGER,
    Name VARCHAR(255),
    Lastname VARCHAR(255),
    Birthday DATE,
    Nationality VARCHAR(255),
    ProvinceID INTEGER,
    CONSTRAINT fk_province FOREIGN KEY (ProvinceID) REFERENCES province(ProvinceID),
    PRIMARY KEY (CitizenID)
);