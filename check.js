class InventoryList {
    constructor() {
      this.itemList = [];
    }
  
    add(name) {
      if (!this.itemList.includes(name)) {
        this.itemList.push(name);
      }
    }
  
    remove(name) {
      const index = this.itemList.indexOf(name);
      if (index !== -1) {
        this.itemList.splice(index, 1);
      }
    }
  
    getList() {
      return this.itemList;
    }
  }
  
  // Stubbed code for testing
  function testInventoryList(inputData) {
    const inventory = new InventoryList();
    const output = [];
  
    for (const line of inputData) {
      const [command, ...params] = line.trim().split(" ");
      if (command === "add") {
        inventory.add(params[0]);
      } else if (command === "remove") {
        inventory.remove(params[0]);
      } else if (command === "getList") {
        const items = inventory.getList();
        if (items.length === 0) {
          output.push("No Items");
        } else {
          output.push(items.join(","));
        }
      }
    }
  
    return output;
  }
  
  // Sample input data as an array of strings
  const input_data = [
    "add shirt",
    "add trouser",
    "getList",
    "remove shirt",
    "getList",
  ];
  
  // Testing the implementation with the sample input data
  const result = testInventoryList(input_data);
  for (const output of result) {
    console.log(output);
  }
  