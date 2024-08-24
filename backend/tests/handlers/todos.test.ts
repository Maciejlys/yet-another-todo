import { test, describe } from "vitest";
import { agent } from "supertest";
import server from "../../src/index";

describe("Todo handler", () => {
  test.skip("should return 200 OK", async () => {
    await agent(server).get("/api/todos").expect([]);
  });

  test.skip("should return 201 OK", async () => {
    await agent(server)
      .post("/api/todos")
      .send({ task: "tested", done: "false" })
      .expect(201);
  });
});
