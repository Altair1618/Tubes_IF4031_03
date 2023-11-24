import { Elysia } from "elysia";
import { auth, googleAuth } from "./configs/lucia";
import { OAuthRequestError } from "@lucia-auth/oauth";
import Routes from "./routes";

const app = new Elysia();

// register routes
Routes(app);

app.listen(3000);

console.log(
	`🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);
