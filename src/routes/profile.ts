import Elysia from "elysia";
import getProfileController from "../controllers/profile/getProfile.controller";
import updateProfileController from "../controllers/profile/updateProfile.controller";
import deleteProfileController from "../controllers/profile/deleteProfile.controller";

export const profileRouteV1 = new Elysia({ prefix: "/profile" })
	.use(getProfileController)
	.use(updateProfileController)
	.use(deleteProfileController);
