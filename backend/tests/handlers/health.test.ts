import { test, describe } from "vitest";
import { agent } from "supertest";
import server from "../../src/index";

describe("Health handler", () => {
  test("should return 200 OK", async () => {
    await agent(server).get("/api/health").expect(200);
  });
});
