import { z } from "zod";

export const updateProfileSchema = z.object({
  name: z.
    string({ required_error: 'Name is required' }).
    trim().
    min(1, { message: 'Name is required' }).
    max(50, { message: 'Name is too long, maximum 50 chars.' })
})
