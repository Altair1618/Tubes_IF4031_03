import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import getEventTicketsService from "../../services/event/getEventTickets.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";

const getEventTicketsController = new Elysia().use(parseJWTMiddleware).get(
    "/:id/tickets",
    async ({ auth: { data }, params }) => {
        const serviceResponse = await getEventTicketsService({
            id: params.id as string,
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

export default getEventTicketsController;
