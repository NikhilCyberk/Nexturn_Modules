const { toggleVisibility } = require("../src/domManipulator");

describe("toggleVisibility", () => {
  let element;
  let setDisplaySpy;

  beforeEach(() => {
    // Create a fresh DOM element before each test
    element = document.createElement("div");
    // Spy on the style object's display property setter
    setDisplaySpy = jest.spyOn(element.style, "display", "set");
  });

  afterEach(() => {
    // Clean up after each test
    jest.restoreAllMocks();
  });

  test("should change display from visible to none", () => {
    // Initial state: visible (empty string is default display)
    element.style.display = "";

    toggleVisibility(element);

    // Assert the style was changed
    expect(element.style.display).toBe("none");
    // Verify the spy was called with the correct value
    expect(setDisplaySpy).toHaveBeenCalledWith("none");
  });

  test("should change display from none to block", () => {
    // Initial state: hidden
    element.style.display = "none";

    toggleVisibility(element);

    // Assert the style was changed
    expect(element.style.display).toBe("block");
    // Verify the spy was called with the correct value
    expect(setDisplaySpy).toHaveBeenCalledWith("block");
  });

  test("should change display from block to none", () => {
    // Initial state: explicitly visible
    element.style.display = "block";

    toggleVisibility(element);

    // Assert the style was changed
    expect(element.style.display).toBe("none");
    // Verify the spy was called with the correct value
    expect(setDisplaySpy).toHaveBeenCalledWith("none");
  });

  test("should track number of style changes", () => {
    // Toggle multiple times
    toggleVisibility(element); // visible -> none
    toggleVisibility(element); // none -> block
    toggleVisibility(element); // block -> none

    // Verify the spy was called exactly 3 times
    expect(setDisplaySpy).toHaveBeenCalledTimes(3);
  });
});
