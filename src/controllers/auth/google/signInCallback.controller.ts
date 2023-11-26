import Elysia, { Context } from "elysia";
import googleSignInCallbackService from "../../../services/auth/google/signInCallback.service";
import { httpResponse } from "../../../utils/httpResponse";

const googleSignInCallbackController = new Elysia().get(
	"/google/callback",
	async (context: Context) => {
		const serviceResponse = await googleSignInCallbackService({ context });
		return httpResponse(serviceResponse);
	},
);

export default googleSignInCallbackController;
