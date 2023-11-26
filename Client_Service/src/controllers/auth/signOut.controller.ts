import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../middlewares/authMiddleware";
import signOutService from "../../services/auth/signOut.service";
import { httpResponse } from "../../utils/httpResponse";

const signOutController = new Elysia().use(parseJWTMiddleware).post(
	"/signout",
	async ({ auth: { data } }) => {
		const sessionId = data?.sessionId;
		const serviceResponse = await signOutService(sessionId);
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

export default signOutController;
