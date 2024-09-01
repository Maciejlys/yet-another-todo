import { describe, it, expect, beforeEach, vi, afterEach } from "vitest";
import DetailsService from "../../src/services/details.service";
import db from "../../src/db/db";

describe("DetailsService", () => {
  beforeEach(() => {
    db.prepare = vi.fn().mockReturnValue({ run: vi.fn() });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  it("should get all details", async () => {
    const mockDetails = [{ todo_id: 1, description: "test" }];
    db.prepare = vi.fn().mockReturnValue({ all: () => mockDetails });

    const details = await DetailsService.getAll();

    expect(details).toEqual(mockDetails);
    expect(db.prepare).toHaveBeenCalledWith("SELECT * FROM details");
  });

  it("should get detail", async () => {
    const mockDetail = { todo_id: 1, description: "test" };
    db.prepare = vi.fn().mockReturnValue({ get: () => mockDetail });

    const detail = await DetailsService.get(1);

    expect(detail).toEqual(mockDetail);
    expect(db.prepare).toHaveBeenCalledWith(
      "SELECT * FROM details WHERE todo_id == (?)",
    );
  });
});
