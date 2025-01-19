const { getElement } = require("../src/arrayUtils");

describe("getElement", () => {
  const testArray = [1, 2, 3, 4, 5];

  // Tests for valid indices
  describe("valid indices", () => {
    test("should return first element with index 0", () => {
      expect(getElement(testArray, 0)).toBe(1);
    });

    test("should return last element with index length-1", () => {
      expect(getElement(testArray, testArray.length - 1)).toBe(5);
    });

    test("should return middle element", () => {
      expect(getElement(testArray, 2)).toBe(3);
    });
  });

  // Tests for invalid indices
  describe("invalid indices", () => {
    test("should throw error for negative index", () => {
      expect(() => {
        getElement(testArray, -1);
      }).toThrow("Index out of bounds");
    });

    test("should throw error for index equal to array length", () => {
      expect(() => {
        getElement(testArray, testArray.length);
      }).toThrow("Index out of bounds");
    });

    test("should throw error for index greater than array length", () => {
      expect(() => {
        getElement(testArray, testArray.length + 1);
      }).toThrow("Index out of bounds");
    });
  });

  // Edge cases
  describe("edge cases", () => {
    test("should handle empty array", () => {
      expect(() => {
        getElement([], 0);
      }).toThrow("Index out of bounds");
    });

    test("should handle very large indices", () => {
      expect(() => {
        getElement(testArray, 1000000);
      }).toThrow("Index out of bounds");
    });
  });
});
