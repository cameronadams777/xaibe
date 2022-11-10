import { defineStore } from "pinia";

export const emptyRemoveUserConfirmationProps: IRemoveUserConfirmationProps = {
  isOpen: false,
  userId: undefined,
  teamId: undefined,
};

interface IRemoveUserConfirmationProps {
  isOpen: boolean;
  userId: number | undefined;
  teamId: number | undefined;
}

export const emptyAddUserToTeamProps: IAddUserToTeamProps = {
  isOpen: false,
  teamId: undefined,
};

interface IAddUserToTeamProps {
  isOpen: boolean;
  teamId: number | undefined;
}

interface IModalStoreState {
  addUserToTeamProps: IAddUserToTeamProps;
  isNewElementModalShown: boolean;
  isDeleteApplicationConfirmationModalShown: boolean;
  isDeleteTeamConfirmationModalShown: boolean;
  removeUserConfirmationProps: IRemoveUserConfirmationProps;
}

export const useModalStore = defineStore("modals", {
  state: (): IModalStoreState => {
    return {
      addUserToTeamProps: emptyAddUserToTeamProps,
      isNewElementModalShown: false,
      isDeleteApplicationConfirmationModalShown: false,
      isDeleteTeamConfirmationModalShown: false,
      removeUserConfirmationProps: emptyRemoveUserConfirmationProps,
    };
  },
  actions: {
    setAddUserToTeamProps(props: IAddUserToTeamProps): void {
      this.addUserToTeamProps = props;
    },
    setIsNewElementModalShown(newValue: boolean): void {
      this.isNewElementModalShown = newValue;
    },
    setIsDeleteApplicationConfirmationModalShown(newValue: boolean): void {
      this.isDeleteApplicationConfirmationModalShown = newValue;
    },
    setIsDeleteTeamConfirmationModalShown(newValue: boolean): void {
      this.isDeleteTeamConfirmationModalShown = newValue;
    },
    setRemoveUserConfirmationProps(props: IRemoveUserConfirmationProps): void {
      this.removeUserConfirmationProps = props;
    },
  },
});
