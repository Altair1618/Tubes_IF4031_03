import Elysia from "elysia";
import getBookingController from "../controllers/booking/getBooking.controller";
import getBookingsController from "../controllers/booking/getBookings.controller";
import cancelBookingController from "../controllers/booking/cancelBooking.controller";

export const profileRouteV1 = new Elysia({ prefix: "/bookings" })
	.use(getBookingsController)
	.use(getBookingController)
	.use(cancelBookingController)