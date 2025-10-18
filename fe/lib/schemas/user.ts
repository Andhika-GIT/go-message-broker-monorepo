import { z } from "zod"

export const UserSchema = z.object({
      id: z.number(),
      name: z.string(),
      email: z.string(),
      phone_number: z.string(),
})

export type User = z.infer<typeof UserSchema>
