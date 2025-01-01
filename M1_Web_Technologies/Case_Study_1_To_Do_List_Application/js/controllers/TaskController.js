import { Task } from "../models/Task.js";
import { TaskView } from "../views/TaskView.js";
import { StorageService } from "../services/StorageService.js";
import { DragDropService } from "../services/DragDropService.js";

export class TaskController {
  constructor() {
    this.storageService = new StorageService();
    this.tasks = this.loadTasks();

    this.taskList = document.getElementById("taskList");
    this.taskInput = document.getElementById("taskInput");
    this.addTaskForm = document.getElementById("addTaskForm");
    this.taskCount = document.querySelector(".task-count");

    this.dragDropService = new DragDropService(
      this.taskList,
      this.handleReorder.bind(this)
    );

    this.init();
  }

  init() {
    this.renderTasks();
    this.setupEventListeners();
  }

  loadTasks() {
    const savedTasks = this.storageService.getTasks();
    return savedTasks.map((taskData) => {
      const task = new Task(taskData.text);
      Object.assign(task, taskData);
      return task;
    });
  }

  setupEventListeners() {
    this.addTaskForm.addEventListener("submit", (e) => {
      e.preventDefault();
      this.addTask();
    });
  }

  addTask() {
    const taskText = this.taskInput.value.trim();
    if (taskText) {
      const task = new Task(taskText);
      this.tasks.push(task);
      this.saveTasks();
      this.renderTasks();
      this.taskInput.value = "";
    }
  }

  toggleTask(id) {
    const task = this.tasks.find((t) => t.id === id);
    if (task) {
      task.toggle();
      this.saveTasks();
      this.renderTasks();
    }
  }

  editTask(id) {
    const task = this.tasks.find((t) => t.id === id);
    if (task) {
      const newText = prompt("Edit task:", task.text);
      if (newText !== null && newText.trim() !== "") {
        task.update(newText.trim());
        this.saveTasks();
        this.renderTasks();
      }
    }
  }

  deleteTask(id) {
    this.tasks = this.tasks.filter((task) => task.id !== id);
    this.saveTasks();
    this.renderTasks();
  }

  handleReorder() {
    const newOrder = Array.from(this.taskList.children).map((item) =>
      this.tasks.find((task) => task.id === parseInt(item.dataset.id))
    );
    this.tasks = newOrder;
    this.saveTasks();
  }

  saveTasks() {
    this.storageService.saveTasks(this.tasks);
    this.updateTaskCount();
  }

  updateTaskCount() {
    const pendingCount = this.tasks.filter((task) => !task.completed).length;
    this.taskCount.textContent = `${pendingCount} task${
      pendingCount !== 1 ? "s" : ""
    } pending`;
  }

  renderTasks() {
    this.taskList.innerHTML = "";
    this.tasks.forEach((task) => {
      const taskView = new TaskView(
        task,
        this.toggleTask.bind(this),
        this.editTask.bind(this),
        this.deleteTask.bind(this)
      );
      this.taskList.appendChild(taskView.render());
    });

    this.updateTaskCount();
  }
}
