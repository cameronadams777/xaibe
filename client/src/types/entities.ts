import { z } from "zod";

export const InviteStatus = {
  "PENDING": 0,
  "ACCEPTED": 1,
  "REJECTED": 2
};

export const UserSchema = z.object({
  id: z.string(),
  created_at: z.date(),
  updated_at: z.date(),
  deleted_at: z.date(),
  first_name: z.string(),
  last_name: z.string(),
  email: z.string(),
  stripe_id: z.string(),
  is_admin: z.boolean(),
  is_verified: z.boolean(),
  reset_password_code: z.string(),
  reset_password_expiry: z.date(),
  applications: z.any(),
  teams: z.any(),
});

export type User = z.infer<typeof UserSchema>;

export const TeamSchema = z.object({
  id: z.string(),
  created_at: z.date(),
  updated_at: z.date(),
  deleted_at: z.date(),
  name: z.string(),
  users: z.lazy(() => z.array(UserSchema)),
  managers: z.lazy(() => z.array(UserSchema)),
  applications: z.any(),
});

export type Team = z.infer<typeof TeamSchema>;

export const ApplicationSchema = z.object({
  id: z.string(),
  created_at: z.date(),
  updated_at: z.date(),
  deleted_at: z.date(),
  name: z.string(),
  team_id: z.string().optional(),
  team: z.lazy(() => TeamSchema).optional(),
  user_id: z.string().optional(),
  user: z.lazy(() => UserSchema).optional(),
  unique_id: z.string(),
  alert_schema_id: z.string().optional(),
  alert_schema: z.lazy(() => AlertSchemaSchema),
  service_tokens: z.lazy(() => z.array(ServiceTokenSchema)),
});

export type Application = z.infer<typeof ApplicationSchema>;

export const AlertSchemaSchema = z.object({
  id: z.string(),
  created_at: z.date(),
  updated_at: z.date(),
  deleted_at: z.date(),
  application_id: z.string(),
  title: z.string(),
  description: z.string(),
  link: z.string(),
});

export type AlertSchema = z.infer<typeof AlertSchemaSchema>;

export const ServiceTokenSchema = z.object({
  id: z.number(),
  created_at: z.date(),
  updated_at: z.date(),
  deleted_at: z.date(),
  token: z.string(),
  application_id: z.string(),
  application: z.any(),
  expires_at: z.date(),
});

export type ServiceToken = z.infer<typeof ServiceTokenSchema>;

export const TeamInviteSchema = z.object({
  id: z.number(),
  created_at: z.date(),
  updated_at: z.date(),
  deleted_at: z.date(),
  sender_id: z.string(),
  sender: z.lazy(() => UserSchema),
  team_id: z.string(),
  team: z.lazy(() => TeamSchema),
  email: z.string(),
  status: z.number(),
});

export type TeamInvite = z.infer<typeof TeamInviteSchema>;

export interface IAlert extends Omit<AlertSchema, "id"> {
  link: string; 
}
