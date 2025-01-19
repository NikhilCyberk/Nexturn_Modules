export function capitalize(word) {
  if (!word) return "";
  return word[0].toUpperCase() + word.slice(1);
}

export function reverseString(str) {
  return str.split("").reverse().join("");
}
