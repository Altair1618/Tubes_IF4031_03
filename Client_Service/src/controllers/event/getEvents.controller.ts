import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import getEventsService from "../../services/event/getEvents.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";

const getEventsController = new Elysia().use(parseJWTMiddleware).get(
    "/",
    async ({ auth: { data }, query }) => {
        const serviceResponse = await getEventsService({
            query: query.query as string,
            page: parseInt(query.page as string),
            jwt: data?.token as string
        });

        return httpResponse(serviceResponse);
    }, 
    {
        beforeHandle: [
            ({ auth: { success, message } }) => {
                return authMiddleware(success, message);
            },
        ],
    },
);

export default getEventsController;