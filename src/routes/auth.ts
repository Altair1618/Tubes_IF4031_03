import Elysia from "elysia";
import googleSignInController from "../controllers/auth/google/signIn.controller";
import googleSignInCallbackController from "../controllers/auth/google/signInCallback.controller";
import signOutController from "../controllers/auth/signOut.controller";

export const authRouteV1 = new Elysia({ prefix: "/auth" })
	.use(googleSignInController)
	.use(googleSignInCallbackController)
	.use(signOutController);
