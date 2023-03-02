import { z } from "zod";

export const NewTeamDetailsFormValidator = z.object({
   numberOfSeats: z.number(),
   businessName: z.string(),
   addressOne: z.string(),
   addressTwo: z.string().optional(),
   city: z.string(),
   state: z.string().optional(),
   country: z.string(),
   zip: z.string().optional(),
   billingEmail: z.string(),
});

export type NewTeamDetailsFormSchema = z.infer<typeof NewTeamDetailsFormValidator>;

