import { describe, it, expect, beforeEach, vi, afterEach } from "vitest";
import TodoService from "../../src/services/todo.service";
import db from "../../src/db/db";

describe("TodoService", () => {
  beforeEach(() => {
    db.prepare = vi.fn().mockReturnValue({ run: vi.fn() });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  it("should get all todos", async () => {
    const mockTodos = [{ id: 1, task: "Test", done: 0 }];
    db.prepare = vi.fn().mockReturnValue({ all: () => mockTodos });

    const todos = await TodoService.getAll();

    expect(todos).toEqual(mockTodos);
    expect(db.prepare).toHaveBeenCalledWith("SELECT * FROM todos");
  });

  it("should add a todo", async () => {
    const mockTodo = { task: "New Task", done: 0 };

    await TodoService.add(mockTodo);

    expect(db.prepare).toHaveBeenCalledWith(
      "INSERT INTO todos(task, done) VALUES (@task, @done)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith(mockTodo);
  });

  it("should delete a todo by id", async () => {
    await TodoService.deleteId(1);

    expect(db.prepare).toHaveBeenCalledWith(
      "DELETE FROM todos WHERE id == (?)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith(1);
  });

  it("should update todo status", async () => {
    const id = 1;
    const mockSet = { done: 1 };

    await TodoService.editDone(id, mockSet);

    expect(db.prepare).toHaveBeenCalledWith(
      "UPDATE todos SET done = (@done) WHERE id == (@id)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith({ id, done: mockSet.done });
  });

  it("should update todo text", async () => {
    const id = 1;
    const mockSet = { task: "Updated Task" };

    await TodoService.editTask(id, mockSet);

    expect(db.prepare).toHaveBeenCalledWith(
      "UPDATE todos SET task = (@task) WHERE id == (@id)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith({ id, task: mockSet.task });
  });
});
