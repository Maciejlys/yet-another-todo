import db from "../db/db";
import { Todo } from "../types";
import { Service } from "./service";

class TodoService extends Service {
  async getAll() {
    return this.db.prepare("SELECT * FROM todos").all();
  }

  async add(todo: Todo) {
    const sql = `INSERT INTO todos(todo, done) VALUES (@task, @done)`;
    return this.db.prepare(sql).run(todo);
  }

  async deleteId(id: number) {
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

export default new TodoService(db);
