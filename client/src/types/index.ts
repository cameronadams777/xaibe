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
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
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
