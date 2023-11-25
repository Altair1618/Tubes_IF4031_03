import Elysia from "elysia";
import { authRouteV1 } from "./auth";
import { profileRouteV1 } from "./profile";

const Routes = (app: Elysia) => {
	app.group("/api", (app) =>
		app.group("/v1", (app) => app.use(profileRouteV1).use(authRouteV1)),
	);
};

export default Routes;
