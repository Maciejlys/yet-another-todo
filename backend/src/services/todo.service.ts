import db from "../db/db";
import { Todo } from "../types";
import { Service } from "./service";

class TodoService extends Service {
  async getAll() {
    return this.db.prepare("SELECT * FROM todos").all();
  }

  async add(todo: Todo) {
    const td = `INSERT INTO todos(task, done) VALUES (@task, @done)`;
    const dt = `INSERT INTO details(todo_id, description) VALUES (@todo_id, @description)`;

    const tran = this.db.transaction(() => {
      const insertTodo = this.db.prepare(td);
      const insertDetails = this.db.prepare(dt);

      const info = insertTodo.run(todo);
      const todoId = info.lastInsertRowid;

      insertDetails.run({ todo_id: todoId, description: todo.description });
    });

    return tran();
  }

  async deleteId(id: number) {
    const sql = `DELETE FROM todos WHERE id == (?)`;
    return this.db.prepare(sql).run(id);
  }

  async editDone(id: number, todo: Partial<Todo>) {
    const sql = `UPDATE todos SET done = (@done) WHERE id == (@id)`;
    return this.db.prepare(sql).run({ id, done: todo.done });
  }

  async editTask(id: number, todo: Partial<Todo>) {
    const sql = `UPDATE todos SET task = (@task) WHERE id == (@id)`;
    return this.db.prepare(sql).run({ id, task: todo.task });
  }
}

export default new TodoService(db);
