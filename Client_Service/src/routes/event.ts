import Elysia from "elysia";
import getEventsController from "../controllers/event/getEvents.controller";

export const eventRouteV1 = new Elysia({ prefix: "/event" })
    .use(getEventsController)