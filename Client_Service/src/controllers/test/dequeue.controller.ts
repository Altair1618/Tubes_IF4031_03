import Elysia from "elysia";
import dequeueBookingService from "../../services/booking/dequeueBooking.service";
import { httpResponse } from "../../utils/httpResponse";

const dequeueBookingController = new Elysia().post(
    "/dequeue",
    async ({ body }: any) => {
        const serviceResponse = await dequeueBookingService({
            ticketId: body?.ticketId,
        });

        return httpResponse(serviceResponse);
    },
)

export default dequeueBookingController;