import Elysia, { Context } from "elysia";
import googleSignInService from "../../../services/auth/google/signIn.service";
import { httpResponse } from "../../../utils/httpResponse";
import parseJWTMiddleware from "../../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../../middlewares/authMiddleware";

const googleSignInController = new Elysia().get(
	"/google",
	async ({ cookie }) => {
		const google_oauth_state = cookie.google_outh_state ?? "";
		const serviceResponse = await googleSignInService({ google_oauth_state });
		return httpResponse(serviceResponse);
	},
);

export default googleSignInController;
