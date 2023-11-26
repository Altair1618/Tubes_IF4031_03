import Elysia from "elysia";
import googleSignInController from "../controllers/auth/google/signIn.controller";
import googleSignInCallbackController from "../controllers/auth/google/signInCallback.controller";
import signOutController from "../controllers/auth/signOut.controller";
import passwordSignUpController from "../controllers/auth/password/signUp.controller";
import passwordSignInController from "../controllers/auth/password/signIn.controller";

export const authRouteV1 = new Elysia({ prefix: "/auth" })
	.use(googleSignInController)
	.use(googleSignInCallbackController)
	.use(passwordSignUpController)
	.use(passwordSignInController)
	.use(signOutController);
