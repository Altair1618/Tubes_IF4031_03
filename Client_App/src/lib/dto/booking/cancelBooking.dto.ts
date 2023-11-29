import { z } from "zod";

export const cancelBookingSchema = z.object({
  id: z
    .string({ required_error: 'Booking ID is required' })
    .trim()
    .min(1, { message: 'Booking ID is required' })
})


export const movePageSchema = z.object({
  page: z

    .number({ required_error: 'Page is required' })
    .min(1, { message: 'Page 1 is minimum' })
})