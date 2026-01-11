import { z } from "zod";

export const loginSchema = z.object({
  email: z
    .email({ error: "Please enter a valid email." })
    .min(1, { error: "This field is required." }),
  password: z
    .string({ error: "Please enter your password." })
    .min(1, { error: "This field is required." }),
});
