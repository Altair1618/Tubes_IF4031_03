import { z } from "zod";

export const processPaymentSchema = z.object({
	paymentUrl: z.string(),
	jwtToken: z.string(),
});

export type ProcessPaymentServicePayload = z.infer<typeof processPaymentSchema>;
