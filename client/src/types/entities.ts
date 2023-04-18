import { z } from "zod";

export const InviteStatus = {
  "PENDING": 0,
  "ACCEPTED": 1,
  "REJECTED": 2
};

export const UserSchema = z.object({
  id: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullish(),
  firstName: z.string(),
  lastName: z.string(),
  email: z.string(),
  stripeId: z.string(),
  isAdmin: z.boolean(),
  isVerified: z.boolean(),
  resetPasswordCode: z.string(),
  resetPasswordExpiry: z.string(),
  applications: z.any(),
  teams: z.any(),
});

export type User = z.infer<typeof UserSchema>;

export const TeamSchema = z.object({
  id: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullish(),
  name: z.string(),
  users: z.lazy(() => z.array(UserSchema)),
  managers: z.lazy(() => z.array(UserSchema)),
  applications: z.any(),
});

export type Team = z.infer<typeof TeamSchema>;

export const ApplicationSchema = z.object({
  id: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullish(),
  name: z.string(),
  teamId: z.string().nullish(),
  team: z.lazy(() => TeamSchema).nullish(),
  userId: z.string().nullish(),
  user: z.lazy(() => UserSchema).nullish(),
  uniqueId: z.string(),
  alertSchemaId: z.string().nullish(),
  alertSchema: z.lazy(() => AlertSchemaSchema),
  serviceTokens: z.lazy(() => z.array(ServiceTokenSchema)),
});

export type Application = z.infer<typeof ApplicationSchema>;

export const AlertSchemaSchema = z.object({
  id: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullish(),
  applicationId: z.string(),
  title: z.string(),
  description: z.string(),
  link: z.string(),
});

export type AlertSchema = z.infer<typeof AlertSchemaSchema>;

export const ServiceTokenSchema = z.object({
  id: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullish(),
  token: z.string(),
  applicationId: z.string(),
  application: z.any(),
  expiresAt: z.string(),
});

export type ServiceToken = z.infer<typeof ServiceTokenSchema>;

export const TeamInviteSchema = z.object({
  id: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullish(),
  senderId: z.string(),
  sender: z.lazy(() => UserSchema),
  teamId: z.string(),
  team: z.lazy(() => TeamSchema),
  email: z.string(),
  status: z.number(),
});

export type TeamInvite = z.infer<typeof TeamInviteSchema>;

export interface IAlert extends Omit<AlertSchema, "id"> {
  link: string; 
}
