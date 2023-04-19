-- Active: 1674940169637@@127.0.0.1@3306@userdDB
CREATE TABLE IF NOT EXISTS DropItUsersDB(
    Secret_Key VARCHAR(255) NOT NULL,
	URole VARCHAR(255) NOT NULL,
	Email VARCHAR (255) NOT NULL,
	Passw VARCHAR (255) NOT NULL,
	First_Name VARCHAR (255) NOT NULL,
	Last_Name VARCHAR (255) NOT NULL,
	LastSearch VARCHAR (255) NULL,
	LastLogin DATE NULL,
	Is_Verified BOOLEAN NOT NULL,
	Creation_Date DATE NOT NULL
) COMMENT '';
INSERT INTO DropItUsersDB(Email, First_Name,Last_Name,PassW, Secret_Key,URole, Creation_Date, Is_Verified ) VALUES ("admin@dropit.com", "Admin","Admin","Admin", "2OdWmGmZucr2JBJllnviBQl5IPw", "Admin", "2023-04-19", FALSE )