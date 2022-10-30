export * from "./button";

export interface IUser {
  ID: number;
  CreateAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date;
  FirstName: string;
  LastName: string;
  Email: string;
  IsAdmin: boolean;
  IsVerified: boolean;
  Teams: ITeam[];
  Applications: IApplication[];
}

export interface ITeam {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date;
  Name: string;
  Applications: IApplication[];
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
}
