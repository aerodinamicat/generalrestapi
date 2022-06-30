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

/** Modelo 'logbook_records' (Registros del Cuaderno de Bitácora)
*
* Propósito: Representa al modelo 'LogbookRecord' definido en los modelos de Go.
 */
DROP TABLE IF EXISTS logbook_records;
CREATE TABLE IF NOT EXISTS logbook_records(
    reference_id VARCHAR(32) NOT NULL,
    recorded_at TIMESTAMP NOT NULL DEFAULT NOW(),
    recorded_by_user_id VARCHAR(32) NOT NULL
);

/** Modelo 'persons' (Personas)
*
* Propósito: Representa al modelo 'Person' definido en los modelos de Go.
 */
DROP TABLE IF EXISTS persons;
CREATE TYPE gender_type AS ENUM('0','1');
CREATE TABLE IF NOT EXISTS persons(
    id VARCHAR(32) PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    second_name VARCHAR(32),
    first_surname VARCHAR(32) NOT NULL,
    second_surname VARCHAR(32),
    date_of_birth TIMESTAMP NOT NULL,
    gender gender_type NOT NULL,
    alias VARCHAR(32)
);

/** Modelo 'users' (Usuarios)
*
* Propósito: Representa al modelo 'User' definido en los modelos de Go.
 */
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(32) PRIMARY KEY,
    email VARCHAR(128) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    password_hint VARCHAR(255) NOT NULL,
    password_recovery_question VARCHAR(255) NOT NULL,
    password_recovery_answer VARCHAR(255) NOT NULL,
    person_id VARCHAR(32) NOT NULL,
    FOREIGN KEY(person_id) REFERENCES persons(id)
);

/** Modelo 'companies' (Empresas)
*
* Propósito: Representa al modelo 'Company' definido en los modelos de Go.
 */
DROP TABLE IF EXISTS companies;
CREATE TABLE IF NOT EXISTS companies(
    id VARCHAR(32) PRIMARY KEY,
    business_name VARCHAR(255) NOT NULL,
    trade_name VARCHAR(255)
);

/** Modelo 'accounting_transactions' (Asientos contables)
*
* Propósito: Representa al modelo 'AccountingTransaction' definido en los modelos de Go.
 */
DROP TABLE IF EXISTS accounting_transactions;
CREATE TABLE IF NOT EXISTS accounting_transactions(
    id VARCHAR(32) PRIMARY KEY,
    post_journal_position VARCHAR(32) NOT NULL,
    date_of_registry TIMESTAMP NOT NULL DEFAULT NOW()
    company_id VARCHAR(32) NOT NULL,
    FOREIGN KEY(company_id) REFERENCES companies(id)
);

/** Modelo 'accounting_records' (Movimientos contables)
*
* Propósito: Representa al modelo 'AccountingRecord' definido en los modelos de Go.
 */
DROP TABLE IF EXISTS accounting_records;
CREATE TYPE accounting_record_nature AS ENUM('0','1');
CREATE TABLE IF NOT EXISTS accounting_records(
    id VARCHAR(32) PRIMARY KEY,
    account_number VARCHAR(4) NOT NULL,
    subaccount_number VARCHAR(32) NOT NULL,
    concept VARCHAR(255) NOT NULL,
    nature accounting_record_nature NOT NULL,
    amount NUMERIC DEFAULT O,
    accounting_transaction_id VARCHAR(32) NOT NULL,
    FOREIGN KEY(accounting_transaction_id) REFERENCES accounting_transactions(id)
);
