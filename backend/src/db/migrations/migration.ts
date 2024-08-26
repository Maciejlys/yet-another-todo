import db from "../db";

const todos = `CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL,
        done INTEGER NOT NULL
    )`;

const details = `CREATE TABLE IF NOT EXISTS details (
        todo_id INTEGER NOT NULL REFERENCES todos(id) ON DELETE CASCADE,
        description TEXT NOT NULL
    )`;

db.prepare(todos).run();
db.prepare(details).run();
db.close();
