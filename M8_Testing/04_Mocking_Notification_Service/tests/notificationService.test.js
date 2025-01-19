const { sendNotification } = require("../src/notificationService");

describe("sendNotification", () => {
  // Test case for successful notification
  test("should return success message when notification is sent successfully", () => {
    // Create a mock notification service
    const mockNotificationService = {
      send: jest.fn().mockReturnValue(true),
    };

    const message = "Hello, World!";
    const result = sendNotification(mockNotificationService, message);

    // Assertions
    expect(result).toBe("Notification Sent");
    expect(mockNotificationService.send).toHaveBeenCalledWith(message);
    expect(mockNotificationService.send).toHaveBeenCalledTimes(1);
  });

  // Test case for failed notification
  test("should return failure message when notification fails to send", () => {
    // Create a mock notification service that fails
    const mockNotificationService = {
      send: jest.fn().mockReturnValue(false),
    };

    const message = "Hello, World!";
    const result = sendNotification(mockNotificationService, message);

    // Assertions
    expect(result).toBe("Failed to Send");
    expect(mockNotificationService.send).toHaveBeenCalledWith(message);
    expect(mockNotificationService.send).toHaveBeenCalledTimes(1);
  });
});
