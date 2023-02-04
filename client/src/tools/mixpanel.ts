import mixpanel, { Mixpanel } from "mixpanel-browser";
import { config } from "src/config"; 

interface IMixpanelWrapper {
  setup: () => void;
  client: Mixpanel;
}

export const mixpanelWrapper: IMixpanelWrapper = {
  setup(): void {
    if (!config.mixpanelToken) { 
      console.log("Warning: No mixpanel token detected.")
      return;
    }
    mixpanel.init(config.mixpanelToken, { debug: config.mixpanelDebug });
    mixpanel.track("Application opened")
  },
  client: mixpanel,
}

