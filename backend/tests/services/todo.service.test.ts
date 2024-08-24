import { describe, it, expect, beforeEach, vi } from "vitest";
import TodoService from "../../src/services/todo.service";
import db from "../../src/db/db";

describe("TodoService", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("should get all todos", async () => {
    const mockTodos = [{ id: 1, todo: "Test", done: 0 }];
    db.prepare = vi.fn().mockReturnValue({ all: () => mockTodos });

    const todos = await TodoService.getAll();

    expect(todos).toEqual(mockTodos);
    expect(db.prepare).toHaveBeenCalledWith("SELECT * FROM todos");
  });

  it("should add a todo", async () => {
    const mockTodo = { task: "New Task", done: 0 };
    db.prepare = vi.fn().mockReturnValue({ run: vi.fn() });

    await TodoService.add(mockTodo);

    expect(db.prepare).toHaveBeenCalledWith(
      "INSERT INTO todos(todo, done) VALUES (@task, @done)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith(mockTodo);
  });

  it("should delete a todo by id", async () => {
    db.prepare = vi.fn().mockReturnValue({ run: vi.fn() });

    await TodoService.deleteId(1);

    expect(db.prepare).toHaveBeenCalledWith(
      "DELETE FROM todos WHERE id == (?)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith(1);
  });

  it("should update todo status", async () => {
    const mockSet = { done: 1, id: 1 };
    db.prepare = vi.fn().mockReturnValue({ run: vi.fn() });

    await TodoService.editDone(mockSet);

    expect(db.prepare).toHaveBeenCalledWith(
      "UPDATE todos SET done = (@done) WHERE id == (@id)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith(mockSet);
  });

  it("should update todo text", async () => {
    const mockSet = { todo: "Updated Task", id: 1 };
    db.prepare = vi.fn().mockReturnValue({ run: vi.fn() });

    await TodoService.editTodo(mockSet);

    expect(db.prepare).toHaveBeenCalledWith(
      "UPDATE todos SET todo = (@todo) WHERE id == (@id)",
    );
    expect(db.prepare("").run).toHaveBeenCalledWith(mockSet);
  });
});
