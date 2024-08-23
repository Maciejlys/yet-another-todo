import db from "../db";

const sql = `CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        todo TEXT NOT NULL,
        done INTEGER NOT NULL
    )`;

db.prepare(sql).run();
db.close();
