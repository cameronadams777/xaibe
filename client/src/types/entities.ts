export enum InviteStatus {
  PENDING = 0,
  ACCEPTED = 1,
  REJECTED = 2,
}

export interface IEntity {
  ID: number;
  CreateAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date;
}

export interface IUser extends IEntity {
  FirstName: string;
  LastName: string;
  Email: string;
  IsAdmin: boolean;
  IsVerified: boolean;
  Teams: ITeam[];
  Applications: IApplication[];
}

export interface ITeam extends IEntity {
  Name: string;
  Users: IUser[];
  Managers: IUser[];
  Applications: IApplication[];
}

export interface ITeamInvite extends IEntity {
  Email: string;
  TeamID: number;
  Team: ITeam;
  SenderID: number;
  Sender: IUser;
  Status: InviteStatus;
}

export interface IApplication {
  ID: number;
  Name: string;
  TeamId?: number;
  UserId?: number;
  UniqueId: string;
  AlertSchema: IAlertSchema;
}

export interface IAlert extends Omit<IAlertSchema, "ID"> {
  Link: string;
}

export interface IAlertSchema {
  ID: number;
  Title: string;
  Description: string;
  Link: string;
}
