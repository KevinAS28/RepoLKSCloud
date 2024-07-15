function sumNumbers(intArg, floatArg) {
    let hasil = intArg + floatArg
    return hasil;
  }
  
  
let sum = sumNumbers(5, 3.14);
console.log(`The sum is: ${sum}`);

// ================================================

function convertAndSum(dataList) {
    let totalSum = 0.0;
    for (let item of dataList) {
      let floatItem = parseFloat(item);
      totalSum += floatItem;
    }
    return totalSum;
  }

let mixedData = [10, 3.14, "5", 20];  
let total = convertAndSum(mixedData);

console.log("total:", total);
