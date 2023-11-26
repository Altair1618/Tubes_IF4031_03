import Elysia from "elysia";
import jwt from "jsonwebtoken";
import { auth } from "../configs/lucia";

const parseJWTMiddleware = new Elysia({ name: "authMiddleware" }).derive(
	async ({ request: { headers } }) => {
		const authorizationHeader = headers.get("Authorization");
		if (!authorizationHeader) {
			return {
				auth: {
					success: false,
					message: "Unauthorized",
				},
			};
		}

		const authorizationHeaderArray = authorizationHeader.split(" ");
		if (authorizationHeaderArray.length !== 2) {
			return {
				auth: {
					success: false,
					message: "Unauthorized",
				},
			};
		}

		const tokenType = authorizationHeaderArray[0];
		const token = authorizationHeaderArray[1];

		if (tokenType !== "Bearer") {
			return {
				auth: {
					success: false,
					message: "Unauthorized",
				},
			};
		}

		try {
			const decodedToken = jwt.verify(
				token,
				process.env.RSA_PUBLIC_KEY ?? "",
			) as jwt.JwtPayload;

			await auth.validateSession(decodedToken.sessionId);

			return {
				auth: {
					success: true,
					data: { ...decodedToken },
				},
			};
		} catch (e) {
			console.log(e);
			return {
				auth: {
					success: false,
					message: "Unauthorized",
				},
			};
		}
	},
);

export default parseJWTMiddleware;
