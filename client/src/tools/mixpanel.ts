import mixpanel, { Mixpanel } from "mixpanel-browser";
import { config } from "src/config"; 

interface IMixpanelWrapper {
  setup: () => void;
  client: Mixpanel;
}

export const mixpanelWrapper: IMixpanelWrapper = {
  setup(): void {
    const allowTelemetry: boolean = localStorage.getItem("allowTelemetry") === "true";
    if (!config.mixpanelToken) { 
      console.log("Warning: No mixpanel token detected.")
      return;
    }
    if (!allowTelemetry) {
      console.log("Telemetry tracking not enabled");
      return;
    }
    mixpanel.init(config.mixpanelToken, { debug: config.mixpanelDebug });
    mixpanel.track("Application opened")
  },
  client: {
    ...mixpanel,
    track: localStorage.getItem("allowTelemetry") === "true" 
      ? mixpanel.track 
      : (_message: string) => console.log("Telemetry tracking not enabled"), 
  },
}

