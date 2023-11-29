import { z } from "zod"

export const getEventsRequestSchema = z.object({
    query: z.string({
        invalid_type_error: "Query is not in a valid type",
    }).optional(),
    page: z.number({
        required_error: "Page is required",
        invalid_type_error: "Page is not in a valid type",
    }).min(1),
    jwt: z.string({
        required_error: "JWT Token is required",
        invalid_type_error: "JWT Token is not in a valid type",
    }).trim().min(1, {
        message: "JWT Token is required",
    }),
})

export type GetEventsRequest = z.infer<typeof getEventsRequestSchema>