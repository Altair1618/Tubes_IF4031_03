import { Context } from "elysia";
import googleSignInCallbackService from "../../../services/auth/google/signInCallback.service";
import { httpResponse } from "../../../utils/httpResponse";

const googleSignInCallbackController = async (context: Context) => {
	const serviceResponse = await googleSignInCallbackService({ context });
	return httpResponse(serviceResponse);
};

export default googleSignInCallbackController;
