import db from "../db/db";
import { Service } from "./service";

class DetailsService extends Service {
  async getAll() {
    return this.db.prepare("SELECT * FROM details").all();
  }

  async get(id: number) {
    const sql = `SELECT * FROM details WHERE todo_id == (?)`;
    return this.db.prepare(sql).get(id);
  }
}

export default new DetailsService(db);
