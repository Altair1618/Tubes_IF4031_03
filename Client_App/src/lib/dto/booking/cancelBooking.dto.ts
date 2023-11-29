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

export const purchaseSchema = z.object({
  payment_url: z

  .string({ required_error: 'Payment URL is required' })
  .trim()
  .min(1, { message: 'Payment URL is required' })
})