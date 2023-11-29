import Elysia from "elysia";
import processPaymentController from "../controllers/payment/processPayment.controller";

export const paymentRouteV1 = new Elysia({ prefix: "/paymenet" }).use(
	processPaymentController,
);
