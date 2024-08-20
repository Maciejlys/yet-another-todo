import { NextFunction, RequestHandler, Request, Response } from "express";

export const handleHealthCheck: RequestHandler = async (
  _: Request,
  res: Response,
  __: NextFunction,
) => {
  return res.send("OK").status(200);
};
