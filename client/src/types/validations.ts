import { z } from "zod";

export const NewTeamDetailsFormValidator = z.object({
  businessName: z.string(),
  addressLineOne: z.string(),
  addressLineTwo: z.string().optional(),
  city: z.string(),
  state: z.string().optional(),
  country: z.string(),
  zip: z.string().optional(),
  billingEmail: z.string(),
});

export type NewTeamDetailsFormSchema = z.infer<typeof NewTeamDetailsFormValidator>;

export const NewTeamSubscriptionFormValidator = z.object({
  teamName: z.string(),
  numberOfSeats: z.number(), 
});

export type NewTeamSubscriptionFormSchema = z.infer<typeof NewTeamSubscriptionFormValidator>;
