import { Request, RequestHandler, Response } from "express";
import detailsService from "../services/details.service";

export const getDetails: RequestHandler = async (_: Request, res: Response) => {
  try {
    const details = await detailsService.getAll();
    return res.json(details);
  } catch (error) {
    console.log("Error getting details");
    res.json([]);
  }
};

export const getDetail: RequestHandler = async (
  req: Request,
  res: Response,
) => {
  try {
    const detail = await detailsService.get(+req.params.id);
    return res.json(detail);
  } catch (error) {
    if (error instanceof Error) {
      console.log(error.message);
    }
    res.json([]);
  }
};
