import Database, { Database as DB } from "better-sqlite3";

class DatabaseFacade {
  private db: DB;

  constructor() {
    this.db = new Database("./todos.db", { verbose: console.log });
  }

  getDb() {
    return this.db;
  }
}

export default new DatabaseFacade().getDb();
