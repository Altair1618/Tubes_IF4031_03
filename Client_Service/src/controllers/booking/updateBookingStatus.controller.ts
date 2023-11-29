import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../middlewares/authMiddleware";

const updateBookingStatusController = new Elysia()
	.use(parseJWTMiddleware)
	.patch("/:id", async () => {}, {
		beforeHandle: [
			({ auth: { success, message } }) => {
				return authMiddleware(success, message);
			},
		],
	});

export default updateBookingStatusController;
