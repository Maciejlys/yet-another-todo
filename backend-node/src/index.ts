import express from "express";
import cors from "cors";
import bodyParser from "body-parser";
import router from "./handlers";

const server = express();

server.use(bodyParser.urlencoded({ extended: true }));
server.use(bodyParser.json());
server.use(
  cors({
    origin: "*", // @TODO: Change this
  }),
);

server.use("/api", router);
server.use((_, res) => {
  return res.sendStatus(404);
});

const start = () => {
  return server.listen(3000, () => {
    console.log(`Server is running on http://localhost:${3000}`);
  });
};

if (process.env.NODE_ENV !== "test") {
  console.log(`Staring server...`);
  start();
}

export default server;
export { start };
