import { z } from "zod"

export const OrderSchema = z.object({
      id: z.number(),
      email: z.string(),
      product_name: z.string(),
      quantity: z.number(),
})

export type Order = z.infer<typeof OrderSchema>