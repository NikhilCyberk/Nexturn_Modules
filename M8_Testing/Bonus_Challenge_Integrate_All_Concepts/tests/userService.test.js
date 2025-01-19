const { fetchAndDisplayUser } = require("../src/userService");

describe("fetchAndDisplayUser", () => {
  let element;
  let mockApiService;

  beforeEach(() => {
    // Set up DOM element
    element = document.createElement("div");

    // Create mock API service
    mockApiService = {
      getUser: jest.fn(),
    };
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("should display user name when API call is successful", async () => {
    // Arrange
    const userId = "123";
    const mockUser = { name: "John Doe" };
    mockApiService.getUser.mockResolvedValue(mockUser);

    // Act
    await fetchAndDisplayUser(mockApiService, userId, element);

    // Assert
    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(element.textContent).toBe("Hello, John Doe");
  });

  test("should display error when API call fails", async () => {
    // Arrange
    const userId = "123";
    const errorMessage = "API Error";
    mockApiService.getUser.mockRejectedValue(new Error(errorMessage));

    // Act
    await fetchAndDisplayUser(mockApiService, userId, element);

    // Assert
    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(element.textContent).toBe(errorMessage);
  });

  test("should display error when user data is invalid", async () => {
    // Arrange
    const userId = "123";
    const invalidUser = { name: "" }; // Invalid user with empty name
    mockApiService.getUser.mockResolvedValue(invalidUser);

    // Act
    await fetchAndDisplayUser(mockApiService, userId, element);

    // Assert
    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(element.textContent).toBe("Invalid user data");
  });

  test("should handle network timeout", async () => {
    // Arrange
    const userId = "123";
    mockApiService.getUser.mockImplementation(
      () =>
        new Promise((_, reject) =>
          setTimeout(() => reject(new Error("Network timeout")), 100)
        )
    );

    // Act
    await fetchAndDisplayUser(mockApiService, userId, element);

    // Assert
    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(element.textContent).toBe("Network timeout");
  });

  test("should handle multiple consecutive calls", async () => {
    // Arrange
    const userId = "123";
    mockApiService.getUser
      .mockResolvedValueOnce({ name: "John Doe" })
      .mockRejectedValueOnce(new Error("API Error"))
      .mockResolvedValueOnce({ name: "Jane Doe" });

    // Act
    await fetchAndDisplayUser(mockApiService, userId, element);
    await fetchAndDisplayUser(mockApiService, userId, element);
    await fetchAndDisplayUser(mockApiService, userId, element);

    // Assert
    expect(mockApiService.getUser).toHaveBeenCalledTimes(3);
    expect(element.textContent).toBe("Hello, Jane Doe");
  });
});
