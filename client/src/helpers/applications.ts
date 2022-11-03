import { ApplicationType, IAlertSchema } from "src/types";

export const getAppSchemaByType = (appType: ApplicationType): IAlertSchema => {
  if (appType === ApplicationType.OTHER)
    throw new Error(
      'Galata Error: No pre-defined schema for "other" application type'
    );

  const applicationSchemasByType: Record<string, any> = {
    [ApplicationType.AIRBRAKE]: {
      Title: "error.project.name",
      Description: "error.error_message",
      Link: "airbrake_error_url",
    },
    [ApplicationType.NEWRELIC]: {
      Title: "",
      Description: "",
      Link: "",
    },
    [ApplicationType.SENTRY]: {
      Title: "",
      Description: "",
      Link: "",
    },
  };

  return applicationSchemasByType[appType];
};
