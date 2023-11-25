import Elysia, { t } from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../middlewares/authMiddleware";
import updateProfileService from "../../services/profile/updateProfile.service";
import { httpResponse } from "../../utils/httpResponse";
import {UpdateProfileServicePayload} from "../../dto/profile/updateProfile.dto";

const updateProfileController = new Elysia().use(parseJWTMiddleware).patch(
	"/",
	async ({ body }) => {
		const serviceResponse = await updateProfileService({...(body as UpdateProfileServicePayload)});
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

export default updateProfileController;
