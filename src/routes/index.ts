import Elysia from "elysia";
import { authRouteV1 } from "./auth";

const Routes = (app: Elysia) => {
	app.group("/api", (app) => app.group("/v1", (app) => app.use(authRouteV1)));
};

export default Routes;
