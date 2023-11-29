import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import requestBookingService from "../../services/booking/requestBooking.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";

const requestBookingController = new Elysia().use(parseJWTMiddleware).post(
    "/",
    async ({ auth: { data }, body }: any) => {
        const serviceResponse = await requestBookingService({
            userId: data?.userId,
            ticketId: body?.ticketId,
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

export default requestBookingController;
