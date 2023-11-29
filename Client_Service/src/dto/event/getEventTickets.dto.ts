import { z } from "zod"

export const getEventTicketsRequestSchema = z.object({
    id: z.string({
        required_error: "Id is required",
        invalid_type_error: "Id is not in a valid type",
    }).uuid(),
    jwt: z.string({
        required_error: "JWT Token is required",
        invalid_type_error: "JWT Token is not in a valid type",
    }).trim().min(1, {
        message: "JWT Token is required",
    }),
})

export type GetEventTicketsRequest = z.infer<typeof getEventTicketsRequestSchema>