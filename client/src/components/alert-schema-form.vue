<script lang="ts" setup>
/**
 * 1. Take in object as prop that resembles an alert received for
 *    the current application.
 * 2. Recursively turn object into dom elements where object keys are buttons
 *    and values are spans.
 * 3. Define an emitter for the key buttons to call when clicked that will send
 *    their path to the root up as a concatenated string
 */

import AlertSchemaTreeBuilder from "./alert-schema-tree-builder.vue";

const obj: Record<string, any> = {
  airbrake_error_url: "https://airbrake.io/airbrake-error-url",
  error: {
    environment: "avocado",
    error_class: "KitchenException",
    error_message: "KitchenException: You are all out of bacon!",
    file: "[PROJECT_ROOT]/app/controllers/bacon_controller.rb",
    first_occurred_at: "2012-02-23T22:03:03Z",
    id: 37463546,
    last_notice: {
      backtrace: [
        "[PROJECT_ROOT]/app/controllers/bacon_controller.rb:35:in `cook'",
        "[PROJECT_ROOT]/app/middleware/kitchen.rb:19:in `oven'",
        "[PROJECT_ROOT]/app/middleware/kitchen.rb:33:in `chef'",
        "[PROJECT_ROOT]/app/middleware/salumi.rb:23:in `store'",
      ],
      id: 4505303522,
      request_method: null,
      request_url: "http://airbrake.io:445/bacon/cook",
    },
    last_occurred_at: "2012-03-21T08:37:15Z",
    line_number: 35,
    project: {
      id: 1111,
      name: "Baconator",
    },
    times_occurred: 118,
  },
};
</script>

<template>
  <div>
    <p>Below what you sent us for this application!</p>
    <alert-schema-tree-builder :schema-object="obj" />
  </div>
</template>
