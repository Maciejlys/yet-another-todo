import { Todo } from "./types";
import Database, { Database as DB } from "better-sqlite3";

class DatabaseFacade {
  private db: DB;

  constructor() {
    this.db = new Database("./todos.db", { verbose: console.log });
    this.init();
  }

  private init() {
    const sql = `CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        todo TEXT NOT NULL,
        done INTEGER NOT NULL
    )`;
    this.db.prepare(sql).run();
  }

  async getAll() {
    return this.db.prepare("SELECT * FROM todos").all();
  }

  async addTodo(todo: Todo) {
    const sql = `INSERT INTO todos(todo, done) VALUES (@task, @done)`;
    return this.db.prepare(sql).run(todo);
  }

  async deleteTodo(id: number) {
    const sql = `DELETE FROM todos WHERE id == (?)`;
    return this.db.prepare(sql).run(id);
  }

  async editDone(set: { done: number; id: number }) {
    const sql = `UPDATE todos SET done = (@done) WHERE id == (@id)`;
    return this.db.prepare(sql).run(set);
  }

  async editTodo(set: { todo: string; id: number }) {
    const sql = `UPDATE todos SET todo = (@todo) WHERE id == (@id)`;
    return this.db.prepare(sql).run(set);
  }
}

export default new DatabaseFacade();
