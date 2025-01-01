import { categories } from "../config/categories.js";

export class ExpenseListView {
  constructor(onDelete) {
    this.tableBody = document.getElementById("expenseTableBody");
    this.onDelete = onDelete;
  }

  render(expenses) {
    this.tableBody.innerHTML = "";
    expenses.forEach((expense) => {
      const row = this.createExpenseRow(expense);
      this.tableBody.appendChild(row);
    });
  }

  createExpenseRow(expense) {
    const row = document.createElement("tr");
    row.innerHTML = `
            <td>${expense.date}</td>
            <td>${expense.description}</td>
            <td>
                <span class="category-badge" style="background-color: ${
                  categories[expense.category]
                }20; 
                    color: ${categories[expense.category]}">
                    ${expense.category}
                </span>
            </td>
            <td>₹${expense.amount.toFixed(2)}</td>
            <td>
                <button class="delete-btn">Delete</button>
            </td>
        `;

    const deleteBtn = row.querySelector(".delete-btn");
    deleteBtn.addEventListener("click", () => this.onDelete(expense.id));

    return row;
  }
}
