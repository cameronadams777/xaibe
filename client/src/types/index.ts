export interface IUser {
  id: number;
  createAt: Date;
  updatedAt: Date;
  deletedAt: Date;
  firstName: string;
  lastName: string;
  email: string;
  isAdmin: boolean;
  isVerified: boolean;
  applications: IApplication[];
}

export interface ITeam {
  id: number;
  created_at: Date;
  updated_at: Date;
  deleted_at: Date;
  name: string;
  applications: IApplication[];
}

export interface IApplication {
  id: number;
  name: string;
  teamId?: number;
  userId?: number;
  uniqueId: string;
}
