import { Context, Cookie } from "elysia";

export interface GoogleSignInServicePayload {
	google_oauth_state: Cookie;
}

export interface GoogleSignInCallbackServicePayload {
	context: Context;
}
