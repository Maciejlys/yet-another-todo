import { RequestHandler, Request, Response } from "express";

export const handleHealthCheck: RequestHandler = async (
  _: Request,
  res: Response,
) => {
  return res.send("OK").status(200);
};
