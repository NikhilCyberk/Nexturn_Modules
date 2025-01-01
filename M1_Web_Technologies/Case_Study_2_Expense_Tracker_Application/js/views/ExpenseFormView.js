import { categories } from "../config/categories.js";

export class ExpenseFormView {
  constructor(onSubmit) {
    this.form = document.getElementById("expenseForm");
    this.categorySelect = document.getElementById("category");
    this.initializeCategories();
    this.bindEvents(onSubmit);
  }

  initializeCategories() {
    Object.keys(categories).forEach((category) => {
      const option = document.createElement("option");
      option.value = category;
      option.textContent = category;
      this.categorySelect.appendChild(option);
    });
  }

  bindEvents(onSubmit) {
    this.form.addEventListener("submit", (e) => {
      e.preventDefault();
      const formData = this.getFormData();
      onSubmit(formData);
      this.form.reset();
    });
  }

  getFormData() {
    return {
      amount: document.getElementById("amount").value,
      description: document.getElementById("description").value,
      category: this.categorySelect.value,
    };
  }
}
