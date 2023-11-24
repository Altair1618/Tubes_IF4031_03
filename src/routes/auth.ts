import Elysia from "elysia";
import googleSignInController from "../controllers/auth/google/signIn.controller";
import googleSignInCallbackController from "../controllers/auth/google/signInCallback.controller";

export const authRouteV1 = new Elysia({ prefix: "/auth" })
	.get("/google/callback", googleSignInCallbackController)
	.get("/google", googleSignInController);
