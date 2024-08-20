import { Router } from "express";
import { handleHealthCheck } from "./health.handler";
import Routes from "./routes";

const router = Router();

router.get(Routes.Health, handleHealthCheck);

export default router;
