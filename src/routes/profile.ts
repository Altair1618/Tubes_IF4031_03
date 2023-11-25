import Elysia from "elysia";
import getProfileController from "../controllers/profile/getProfile.controller";
import authMiddleware from "../middlewares/authMiddleware";

export const profileRouteV1 = new Elysia({ prefix: "/profile" }).use(
	getProfileController,
);
