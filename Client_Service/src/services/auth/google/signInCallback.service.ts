import { OAuthRequestError } from "@lucia-auth/oauth";
import { auth, googleAuth } from "../../../configs/lucia";
import { ServiceResponse } from "../../../types/common";
import jwt from "jsonwebtoken";
import { GoogleSignInCallbackServicePayload } from "../../../types/auth";

const googleSignInCallbackService = async ({
	context,
}: GoogleSignInCallbackServicePayload): Promise<ServiceResponse> => {
	const storedState = context.cookie.google_outh_state.value;
	const state = context.query.state;
	const code = context.query.code;

	if (
		!storedState ||
		!state ||
		storedState !== state ||
		typeof code !== "string"
	) {
		return {
			code: 400,
			message: "Invalid oauth credentials",
		};
	}

	try {
		const { getExistingUser, googleUser, createUser } =
			await googleAuth.validateCallback(code);

		const getUser = async () => {
			const existingUser = await getExistingUser();
			if (existingUser) return existingUser;
			const user = await createUser({
				attributes: {
					email: googleUser.email as string,
					name: googleUser.name,
					picture: googleUser.picture,
				},
			});
			return user;
		};

		const user = await getUser();
		const session = await auth.createSession({
			userId: user.userId,
			attributes: {},
		});
		const authRequest = auth.handleRequest(context);
		authRequest.setSession(session);

		const accessToken = jwt.sign(
			{
				userId: user.userId,
				sessionId: session.sessionId,
			},
			process.env.RSA_PRIVATE_KEY ?? "",
			{ algorithm: "RS256" },
		);
		return {
			headers: {
				Location: `${process.env.CLIENT_BASE_URL}/auth/google?token=${accessToken}`,
			},
			code: 302,
			message: "Success",
		};
	} catch (e) {
		console.log(e);
		if (e instanceof OAuthRequestError) {
			return {
				code: 400,
				message: "Invalid oauth credentials",
			};
		}
		return {
			code: 500,
			message: "Something went wrong, please try again later",
		};
	}
};

export default googleSignInCallbackService;
