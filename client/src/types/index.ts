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
}
