import dataBase from "../db/db";

export class Service {
  constructor(protected db: typeof dataBase) { }
}
