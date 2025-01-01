export class StorageService {
  constructor(key = "expenses") {
    this.storageKey = key;
  }

  getExpenses() {
    return JSON.parse(localStorage.getItem(this.storageKey)) || [];
  }

  saveExpenses(expenses) {
    localStorage.setItem(this.storageKey, JSON.stringify(expenses));
  }
}
