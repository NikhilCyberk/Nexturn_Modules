const { delayedGreeting } = require("../src/greetingService");

describe("delayedGreeting", () => {
  // Setup and teardown for timer mocks
  beforeEach(() => {
    jest.useFakeTimers();
  });

  afterEach(() => {
    jest.useRealTimers();
  });

  test("should resolve with correct greeting message", async () => {
    // Arrange
    const name = "Alice";
    const delay = 1000;
    const expectedGreeting = "Hello, Alice!";

    // Act
    const greetingPromise = delayedGreeting(name, delay);

    // Fast-forward time
    jest.advanceTimersByTime(delay);

    // Assert
    const result = await greetingPromise;
    expect(result).toBe(expectedGreeting);
  });

  test("should not resolve before the specified delay", async () => {
    // Arrange
    const name = "Bob";
    const delay = 2000;
    let resolved = false;

    // Act
    const promise = delayedGreeting(name, delay);
    promise.then(() => {
      resolved = true;
    });

    // Advance time partially
    jest.advanceTimersByTime(delay - 1);

    // Assert
    expect(resolved).toBe(false);
  });

  test("should resolve after exact delay time", async () => {
    // Arrange
    const name = "Charlie";
    const delay = 1500;
    let resolved = false;

    // Act
    const promise = delayedGreeting(name, delay);
    promise.then(() => {
      resolved = true;
    });

    // Advance time to exact delay
    jest.advanceTimersByTime(delay);

    // Ensure all timers are processed
    jest.runAllTimers();

    // Wait for the promise to resolve
    await promise;

    // Assert
    expect(resolved).toBe(true);
  });

  test("should handle multiple concurrent delays correctly", async () => {
    // Arrange
    const promises = [
      delayedGreeting("Dave", 1000),
      delayedGreeting("Eve", 2000),
    ];

    // Act & Assert
    jest.advanceTimersByTime(1000);
    const firstResult = await promises[0];
    expect(firstResult).toBe("Hello, Dave!");

    jest.advanceTimersByTime(1000);
    const secondResult = await promises[1];
    expect(secondResult).toBe("Hello, Eve!");
  });

  test("should handle zero delay", async () => {
    // Arrange
    const name = "Frank";
    const delay = 0;

    // Act
    const promise = delayedGreeting(name, delay);
    jest.advanceTimersByTime(0);

    // Assert
    const result = await promise;
    expect(result).toBe("Hello, Frank!");
  });
});
