import { Router } from "express";
import { handleHealthCheck } from "./health.handler";
import { deleteTodo, getTodos, patchDone, postTodos } from "./todos.handler";
import Routes from "./routes";

const router = Router();

// get
router.get(Routes.Health, handleHealthCheck);
router.get(Routes.Todos, getTodos);

// post
router.post(Routes.Todos, postTodos);

// deleted
router.delete(Routes.Todos, deleteTodo);

// patch
router.patch(Routes.Todos, patchDone);

export default router;
