export class Expense {
  constructor(amount, description, category) {
    this.id = Date.now();
    this.amount = parseFloat(amount);
    this.description = description;
    this.category = category;
    this.date = new Date().toLocaleDateString();
  }
}
