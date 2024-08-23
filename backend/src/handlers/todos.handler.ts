import { RequestHandler, Request, Response } from "express";
import todoService from "../services/todo.service";

export const getTodos: RequestHandler = async (_: Request, res: Response) => {
  try {
    const todos = await todoService.getAll();
    return res.json(todos);
  } catch (error) {
    console.log("Error getting todos");
    res.json([]);
  }
};

export const postTodos: RequestHandler = async (
  req: Request,
  res: Response,
) => {
  try {
    await todoService.add(req.body);
    return res.send("Created").status(201);
  } catch (error) {
    return res.send(error).status(503);
  }
};

export const deleteTodo: RequestHandler = async (
  req: Request,
  res: Response,
) => {
  try {
    await todoService.deleteId(req.body.id);
    return res.send("Deleted").status(204);
  } catch (error) {
    return res.send(error).status(503);
  }
};

export const patchDone: RequestHandler = async (
  req: Request,
  res: Response,
) => {
  try {
    if (!req.body.id) throw new Error("No id provided");
    if (req.body.done) {
      await todoService.editDone(req.body);
    }
    if (req.body.todo) {
      await todoService.editTodo(req.body);
    }
    return res.send("Changed").status(200);
  } catch (error) {
    if (error instanceof Error) {
      return res.send(error.message).status(500);
    }
    return res.send("Error").status(503);
  }
};
