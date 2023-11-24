import { lucia } from "lucia";
import { elysia } from "lucia/middleware";
import { google } from "@lucia-auth/oauth/providers";
import { pg } from "@lucia-auth/adapter-postgresql";
import { pool } from "./drizzle";

export const auth = lucia({
	env: process.env.NODE_ENV === "development" ? "DEV" : "PROD",
	middleware: elysia(),
	adapter: pg(pool, {
		user: "auth_user",
		key: "user_key",
		session: "user_session",
	}),
	getUserAttributes: (data) => {
		return {
			email: data.email,
		};
	},
});

export const googleAuth = google(auth, {
	clientId: process.env.GOOGLE_CLIENT_ID ?? "",
	clientSecret: process.env.GOOGLE_CLIENT_SECRET ?? "",
	redirectUri: process.env.GOOGLE_REDIRECT_URI ?? "",
	scope: ["https://www.googleapis.com/auth/userinfo.email"],
});

export type Auth = typeof auth;
