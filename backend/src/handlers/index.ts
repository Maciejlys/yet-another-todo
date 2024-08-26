import { Router } from "express";
import { handleHealthCheck } from "./health.handler";
import { deleteTodo, getTodos, patchTodo, postTodos } from "./todos.handler";
import { getDetail, getDetails } from "./details.handler";
import Routes from "./routes";

const router = Router();

// get
router.get(Routes.Health, handleHealthCheck);
router.get(Routes.Todos, getTodos);
router.get(Routes.Details, getDetails);
router.get(Routes.Details + "/:id", getDetail);

// post
router.post(Routes.Todos, postTodos);

// delete
router.delete(Routes.Todos + "/:id", deleteTodo);

// patch
router.patch(Routes.Todos + "/:id", patchTodo);

export default router;
