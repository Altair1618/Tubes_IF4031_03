import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../middlewares/authMiddleware";
import deleteProfileService from "../../services/profile/deleteProfile.service";
import { httpResponse } from "../../utils/httpResponse";

const deleteProfileController = new Elysia().use(parseJWTMiddleware).delete(
	"/",
	async ({ auth: { data } }) => {
		const serviceResponse = await deleteProfileService({
			userId: data?.userId,
		});
		return httpResponse(serviceResponse);
	},
	{
		beforeHandle: [
			({ auth: { success, message } }) => {
				return authMiddleware(success, message);
			},
		],
	},
);

export default deleteProfileController;
