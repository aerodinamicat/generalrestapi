--* Sentencias DDL - Definiciónes de datos

/** Creación de la base de datos 'generalrestapi'
* No son necesarias las instrucciones de creación, y uso, de base de datos en PostgreSQL,
* al contrario que MySQL, ya que la propia conexión las compone y ejecuta con sus propios
* parámetros.
* No obstante las dejo aquí abajo escritas porque nunca se sabe si llegarán
* a ser útiles en algún momento.
 */
/**
DROP DATABASE IF EXISTS generalrestapi;
CREATE DATABASE IF NOT EXISTS generalrestapi;

USE generalrestapi;
 */

/** Modelo 'persons' (Personas)
*
* Propósito: Representa al modelo 'Person' definido en los modelos de Go
 */
DROP TABLE IF EXISTS persons;
CREATE TABLE IF NOT EXISTS persons(
    id VARCHAR(32) PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    second_name VARCHAR(32),
    first_surname VARCHAR(32) NOT NULL,
    second_surname VARCHAR(32),
    date_of_birth TIMESTAMP NOT NULL,
    gender VARCHAR(32) NOT NULL,
    alias VARCHAR(32),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by_user_id VARCHAR(32) NOT NULL,
    modified_at TIMESTAMP NOT NULL DEFAULT NOW(),
    modified_by_user_id VARCHAR(32) NOT NULL
);
