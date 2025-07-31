    // "id": 1,
    // "name": "Andi Wijaya",
    // "email": "andi.wijaya@example.com",
    // "phone": "081234567890",
    // "created_at": "2025-07-31T10:15:00Z"
import { z } from "zod"

export const UserSchema = z.object({
      id: z.number(),
      name: z.string(),
      email: z.string(),
      phone: z.string(),
      created_at: z.string(),
})
