import Elysia from "elysia";
import getProfileService from "../../services/profile/getProfile.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";

const getProfileController = new Elysia().use(parseJWTMiddleware).get(
	"/",
	async ({ auth: { data } }) => {
		const serviceResponse = await getProfileService({
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
export default getProfileController;
