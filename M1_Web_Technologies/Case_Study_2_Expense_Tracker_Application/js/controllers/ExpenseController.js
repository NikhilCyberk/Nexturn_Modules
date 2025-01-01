import { Expense } from "../models/Expense.js";
import { StorageService } from "../services/StorageService.js";
import { ChartService } from "../services/ChartService.js";
import { ExpenseFormView } from "../views/ExpenseFormView.js";
import { ExpenseListView } from "../views/ExpenseListView.js";

export class ExpenseController {
  constructor() {
    this.storageService = new StorageService();
    this.chartService = new ChartService();
    this.expenses = this.storageService.getExpenses();

    this.expenseListView = new ExpenseListView(this.handleDelete.bind(this));

    this.expenseFormView = new ExpenseFormView(this.handleSubmit.bind(this));

    this.render();
  }

  handleSubmit({ amount, description, category }) {
    const expense = new Expense(amount, description, category);
    this.expenses.push(expense);
    this.saveAndRender();
  }

  handleDelete(id) {
    this.expenses = this.expenses.filter((expense) => expense.id !== id);
    this.saveAndRender();
  }

  calculateCategoryTotals() {
    return this.expenses.reduce((totals, expense) => {
      totals[expense.category] =
        (totals[expense.category] || 0) + expense.amount;
      return totals;
    }, {});
  }

  saveAndRender() {
    this.storageService.saveExpenses(this.expenses);
    this.render();
  }

  render() {
    this.expenseListView.render(this.expenses);
    const categoryTotals = this.calculateCategoryTotals();
    this.chartService.updateCharts(categoryTotals);
  }
}
