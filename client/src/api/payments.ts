import { NewTeamDetailsFormSchema } from "src/types";
import { rawPost } from "./http";

export const createStripeCustomer = async (formValues: NewTeamDetailsFormSchema): Promise<boolean> => {
  const response = await rawPost({ 
    url: "/api/create-customer", 
    body: {
      ...formValues
    }
  });
  return response.ok;
}
