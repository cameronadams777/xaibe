import { warn } from "console";
import { defineStore } from "pinia";
import { Team } from "src/types";

type RemoveUserConfirmationProps = {
  isOpen: boolean;
  userId: string | undefined;
  teamId: string | undefined;
}

export const emptyRemoveUserConfirmationProps: RemoveUserConfirmationProps = {
  isOpen: false,
  userId: undefined,
  teamId: undefined,
};

type AddUserToTeamProps = {
  isOpen: boolean;
  teamId: string | undefined;
}

export const emptyAddUserToTeamProps: AddUserToTeamProps = {
  isOpen: false,
  teamId: undefined,
};

type TeamSubscriptionDetailsProps = {
  isOpen: boolean;
  team: Team | undefined;
}

export const emptyTeamSubscriptionDetailsProps: TeamSubscriptionDetailsProps = {
  isOpen: false,
  team: undefined,
}

interface IModalStoreState {
  addUserToTeamProps: AddUserToTeamProps;
  teamSubscriptionDetailsProps: TeamSubscriptionDetailsProps;
  isNewElementModalShown: boolean;
  isDeleteApplicationConfirmationModalShown: boolean;
  isDeleteTeamConfirmationModalShown: boolean;
  removeUserConfirmationProps: RemoveUserConfirmationProps;
}

export const useModalStore = defineStore("modals", {
  state: (): IModalStoreState => {
    return {
      addUserToTeamProps: emptyAddUserToTeamProps,
      teamSubscriptionDetailsProps: emptyTeamSubscriptionDetailsProps,
      isNewElementModalShown: false,
      isDeleteApplicationConfirmationModalShown: false,
      isDeleteTeamConfirmationModalShown: false,
      removeUserConfirmationProps: emptyRemoveUserConfirmationProps,
    };
  },
  actions: {
    setAddUserToTeamProps(props: AddUserToTeamProps): void {
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
    setRemoveUserConfirmationProps(props: RemoveUserConfirmationProps): void {
      this.removeUserConfirmationProps = props;
    },
    setTeamSubscriptionDetailsProps(props: TeamSubscriptionDetailsProps): void {
      this.teamSubscriptionDetailsProps = props;
    }
  },
});
