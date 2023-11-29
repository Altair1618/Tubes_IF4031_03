import Elysia from "elysia";
import getBookingsController from "../controllers/booking/getBookings.controller";
import cancelBookingController from "../controllers/booking/cancelBooking.controller";
import requestBookingController from "../controllers/booking/requestBooking.controller";
import dequeueBookingController from "../controllers/test/dequeue.controller";

export const bookingRouteV1 = new Elysia({ prefix: "/bookings" })
	.use(getBookingsController)
	.use(cancelBookingController)
	.use(requestBookingController)
	.use(dequeueBookingController)