export class TaskView {
  constructor(task, onToggle, onEdit, onDelete) {
    this.task = task;
    this.onToggle = onToggle;
    this.onEdit = onEdit;
    this.onDelete = onDelete;
  }

  render() {
    const li = document.createElement("li");
    li.className = `task-item ${this.task.completed ? "completed" : ""}`;
    li.draggable = true;
    li.dataset.id = this.task.id;

    li.innerHTML = `
            <input type="checkbox" class="task-checkbox" ${
              this.task.completed ? "checked" : ""
            }>
            <span class="task-text">${this.task.text}</span>
            <div class="task-actions">
                <button class="edit-btn">Edit</button>
                <button class="delete-btn">Delete</button>
            </div>
        `;

    this.setupEventListeners(li);
    return li;
  }

  setupEventListeners(li) {
    const checkbox = li.querySelector(".task-checkbox");
    const editBtn = li.querySelector(".edit-btn");
    const deleteBtn = li.querySelector(".delete-btn");

    checkbox.addEventListener("change", () => this.onToggle(this.task.id));
    editBtn.addEventListener("click", () => this.onEdit(this.task.id));
    deleteBtn.addEventListener("click", () => this.onDelete(this.task.id));
  }
}
