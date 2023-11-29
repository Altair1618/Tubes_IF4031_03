import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import { httpResponse } from "../../utils/httpResponse";
import getEventDetailService from "../../services/event/getEventDetail.service";
import authMiddleware from "../../middlewares/authMiddleware";

const getEventDetailController = new Elysia().use(parseJWTMiddleware).get(
    "/:id",
    async ({ auth: { data }, params }) => {
        const serviceResponse = await getEventDetailService({
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

export default getEventDetailController;