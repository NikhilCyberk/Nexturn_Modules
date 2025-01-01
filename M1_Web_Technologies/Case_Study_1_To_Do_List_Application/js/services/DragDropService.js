export class DragDropService {
  constructor(taskList, onReorder) {
    this.taskList = taskList;
    this.onReorder = onReorder;
    this.init();
  }

  init() {
    this.taskList.addEventListener(
      "dragstart",
      this.handleDragStart.bind(this)
    );
    this.taskList.addEventListener("dragend", this.handleDragEnd.bind(this));
    this.taskList.addEventListener("dragover", this.handleDragOver.bind(this));
  }

  handleDragStart(e) {
    if (e.target.classList.contains("task-item")) {
      e.target.classList.add("dragging");
      e.dataTransfer.setData("text/plain", e.target.dataset.id);
    }
  }

  handleDragEnd(e) {
    if (e.target.classList.contains("task-item")) {
      e.target.classList.remove("dragging");
    }
  }

  handleDragOver(e) {
    e.preventDefault();
    const afterElement = this.getDragAfterElement(e.clientY);
    const draggable = document.querySelector(".dragging");
    if (afterElement) {
      this.taskList.insertBefore(draggable, afterElement);
    } else {
      this.taskList.appendChild(draggable);
    }
  }

  getDragAfterElement(y) {
    const draggableElements = [
      ...this.taskList.querySelectorAll(".task-item:not(.dragging)"),
    ];
    return draggableElements.reduce(
      (closest, child) => {
        const box = child.getBoundingClientRect();
        const offset = y - box.top - box.height / 2;
        if (offset < 0 && offset > closest.offset) {
          return { offset: offset, element: child };
        }
        return closest;
      },
      { offset: Number.NEGATIVE_INFINITY }
    ).element;
  }
}
