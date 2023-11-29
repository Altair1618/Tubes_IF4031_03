import Elysia from "elysia";
import getEventsController from "../controllers/event/getEvents.controller";
import getEventDetailController from "../controllers/event/getEventDetail.controller";
import getEventTicketsController from "../controllers/event/getEventTickets.controller";

export const eventRouteV1 = new Elysia({ prefix: "/event" })
    .use(getEventsController)
    .use(getEventDetailController)
    .use(getEventTicketsController)
