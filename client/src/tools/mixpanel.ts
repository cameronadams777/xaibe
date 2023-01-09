import mixpanel, { Mixpanel } from "mixpanel-browser";

interface IMixpanelWrapper {
  setup: () => void;
  client: Mixpanel;
}

const token = import.meta.env.VITE_MIXPANEL_TOKEN;
const debug = import.meta.env.VITE_MIXPANEL_DEBUG;

export const mixpanelWrapper: IMixpanelWrapper = {
  setup(): void {
    if (!token.length) {
      console.log("Warning: No mixpanel token detected.")
      return;
    }
    mixpanel.init(token, { debug })
  },
  client: mixpanel,
}

