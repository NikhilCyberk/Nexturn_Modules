module.exports = {
  testEnvironment: "jsdom",
  testMatch: ["**/tests/**/*.test.js"],
  collectCoverage: true,
  coverageDirectory: "coverage",
  coverageReporters: ["text", "lcov"],
};
