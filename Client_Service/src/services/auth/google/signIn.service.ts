import { googleAuth } from "../../../configs/lucia";
import { GoogleSignInServicePayload } from "../../../types/auth";
import { ServiceResponse } from "../../../types/common";

const googleSignInService = async ({
	google_oauth_state,
}: GoogleSignInServicePayload): Promise<ServiceResponse> => {
	const [url, state] = await googleAuth.getAuthorizationUrl();
	google_oauth_state.set({
		value: state,
		httpOnly: true,
		secure: process.env.NODE_EV === "production",
		path: "/",
		maxAge: 60 * 60 * 1000, // 1 hour
	});

	return {
		headers: {
			Location: url.toString(),
		},
		code: 302,
		message: "Success",
	};
};
export default googleSignInService;
