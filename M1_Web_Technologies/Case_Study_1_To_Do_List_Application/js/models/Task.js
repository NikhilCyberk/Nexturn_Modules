export class Task {
  constructor(text) {
    this.id = Date.now();
    this.text = text;
    this.completed = false;
    this.createdAt = new Date();
  }

  toggle() {
    this.completed = !this.completed;
  }

  update(text) {
    this.text = text;
  }
}