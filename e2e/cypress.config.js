const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    experimentalRunAllSpecs: true,
    setupNodeEvents(on, config) {
      if (config.env.baseUrl) {
        config.baseUrl = config.env.baseUrl;
      }
      return config;
    },
  },
});
