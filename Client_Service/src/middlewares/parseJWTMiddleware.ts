import Elysia from "elysia";
import jwt, { decode } from "jsonwebtoken";
import { auth } from "../configs/lucia";

interface CustomJwtPayload extends jwt.JwtPayload {
	token: string
}

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

			// create new jwt token for intra service communication
			const newToken = jwt.sign({ exp: Math.floor(Date.now() / 1000) + (60 * 5), sessionId: decodedToken.sessionId, userId: decodedToken.userI, secret: process.env.JWT_TOKEN_SECRET}, process.env.RSA_PRIVATE_KEY as string, { algorithm: "RS256"}) 

			return {
				auth: {
					success: true,
					data: { token: newToken, ...decodedToken } as CustomJwtPayload,
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
