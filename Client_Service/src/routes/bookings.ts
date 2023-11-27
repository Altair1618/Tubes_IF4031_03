import Elysia from "elysia";
import getBookingGroupController from "../controllers/bookings/getBookingGroup.controller";
import getBookingsController from "../controllers/bookings/getBookings.controller";

export const profileRouteV1 = new Elysia({ prefix: "/bookings" })
	.use(getBookingsController)
	.use(getBookingGroupController)