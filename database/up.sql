DROP TABLE IF EXISTS tutorials;

CREATE TABLE tutorials (
    id VARCHAR(32) PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion VARCHAR(255) NOT NULL,
    estado VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()

);

DROP TABLE IF EXISTS detalles;

CREATE TABLE detalles (
    id VARCHAR(32) PRIMARY KEY,
    autor VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    tutorials_id VARCHAR(32) NOT NULL ,
    FOREIGN KEY (tutorials_id) REFERENCES tutorials(id)

);