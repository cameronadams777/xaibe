import { NewTeamDetailsFormSchema } from "src/types";
import { rawPost } from "./http";

export const createNewStripeCustomer = async (formValues: NewTeamDetailsFormSchema): Promise<boolean> => {
  const response = await rawPost({ 
    url: "/api/create-customer", 
    body: {
      ...formValues
    }
  });
  return response.ok;
}

export const createNewTeamSubscription = async (formValues: NewTeamDetailsFormSchema): Promise<boolean> => {
  const response = await rawPost({ 
    url: "/api/create-new-team-subscription", 
    body: {
      ...formValues
    }
  });
  return response.ok;
}
