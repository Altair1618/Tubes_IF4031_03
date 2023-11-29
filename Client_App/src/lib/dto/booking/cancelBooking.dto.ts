import { z } from "zod";

export const cancelBookingSchema = z.object({
  id: z.
    string({ required_error: 'Booking ID is required' })
    .trim()
    .min(1, { message: 'Booking ID is required' })
})
