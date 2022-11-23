USE devDB

INSERT INTO Province
VALUES
    (1, 'Bangkok'),
    (2, 'Ayutthaya');

INSERT INTO District
VALUES
    (1, 'Phayathai', 1),
    (2, 'Bangbo', 1);

INSERT INTO Population
VALUES
    ("1234567891234", "1234AB", 'Somsri', 'MairuMairu', '2015-12-17', 'Thai', 1),
    ("1234567891235", "CCAADD", 'Smorn', 'YeomPoen', '2015-12-02', 'Thai', 2);

INSERT INTO ApplyVoteMp
VALUES
    (1, "1234567891234");

INSERT INTO ApplyVoteParty
VALUES
    (1, "1234567891234");

INSERT INTO Candidate
VALUES
    (1, "1234567891234"),
    (2, "1234567891235");

INSERT INTO Mp
VALUES
    (1, "1234567891235");