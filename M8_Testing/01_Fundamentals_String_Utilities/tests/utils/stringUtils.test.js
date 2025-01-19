import { capitalize, reverseString } from "../../src/utils/stringUtils";

describe("String Utilities", () => {
  describe("capitalize()", () => {
    test("should capitalize the first letter of a word", () => {
      expect(capitalize("hello")).toBe("Hello");
      expect(capitalize("world")).toBe("World");
      expect(capitalize("javascript")).toBe("Javascript");
    });

    test("should handle single-character words", () => {
      expect(capitalize("a")).toBe("A");
      expect(capitalize("z")).toBe("Z");
    });

    test("should handle empty strings", () => {
      expect(capitalize("")).toBe("");
    });

    test("should not modify already capitalized words", () => {
      expect(capitalize("Hello")).toBe("Hello");
      expect(capitalize("World")).toBe("World");
    });

    test("should handle strings with numbers and special characters", () => {
      expect(capitalize("123abc")).toBe("123abc");
      expect(capitalize("!hello")).toBe("!hello");
    });
  });

  describe("reverseString()", () => {
    test("should reverse a simple string", () => {
      expect(reverseString("hello")).toBe("olleh");
      expect(reverseString("world")).toBe("dlrow");
    });

    test("should handle palindromes correctly", () => {
      expect(reverseString("radar")).toBe("radar");
      expect(reverseString("level")).toBe("level");
      expect(reverseString("madam")).toBe("madam");
    });

    test("should handle empty strings", () => {
      expect(reverseString("")).toBe("");
    });

    test("should handle single-character strings", () => {
      expect(reverseString("a")).toBe("a");
    });

    test("should handle strings with spaces", () => {
      expect(reverseString("hello world")).toBe("dlrow olleh");
    });

    test("should handle strings with numbers and special characters", () => {
      expect(reverseString("hello123!")).toBe("!321olleh");
      expect(reverseString("12345")).toBe("54321");
    });
  });
});
