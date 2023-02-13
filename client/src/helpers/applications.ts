import { ApplicationType, AlertSchema } from "src/types";

export const getAppSchemaByType = (
  appType: ApplicationType
): AlertSchema | undefined => {
  if (appType === ApplicationType.OTHER) return;

  const applicationSchemasByType: Record<string, any> = {
    [ApplicationType.AIRBRAKE]: {
      Title: "error.project.name",
      Description: "error.error_message",
      Link: "airbrake_error_url",
    },
    [ApplicationType.NEWRELIC]: {
      Title: "event_type",
      Description: "details",
      Link: "incident_url",
    },
    [ApplicationType.SENTRY]: {
      Title: "data.error.title",
      Description: "data.error.metadata.value",
      Link: "data.error.issue_url",
    },
  };

  return applicationSchemasByType[appType];
};
