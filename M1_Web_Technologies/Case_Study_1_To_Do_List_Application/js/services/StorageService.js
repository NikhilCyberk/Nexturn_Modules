export class StorageService {
  constructor(storageKey = "tasks") {
    this.storageKey = storageKey;
  }

  getTasks() {
    const tasks = localStorage.getItem(this.storageKey);
    return tasks ? JSON.parse(tasks) : [];
  }

  saveTasks(tasks) {
    localStorage.setItem(this.storageKey, JSON.stringify(tasks));
  }
}
