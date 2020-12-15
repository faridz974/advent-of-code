(() => {
  const content = Deno.readTextFileSync("input.txt");
  const lines = content.split("\n");
  const numbers = lines.map((l) => parseInt(l));

  const solution1 = () => {
    for (let i = 0; i < numbers.length - 1; i++) {
      for (let j = i + 1; j < numbers.length; j++) {
        if (numbers[i] + numbers[j] === 2020) {
          console.log(numbers[i] * numbers[j]);
          return;
        }
      }
    }
  };

  const solution2 = () => {
    for (let i = 0; i < numbers.length - 1; i++) {
      for (let j = i + 1; j < numbers.length; j++) {
        for (let k = j + 1; k < numbers.length; k++) {
          if (numbers[i] + numbers[j] + numbers[k] === 2020) {
            console.log(numbers[i] * numbers[j] * numbers[k]);
            return;
          }
        }
      }
    }
  };

  if (Deno.args[0] === "1") {
    solution1();
  } else if (Deno.args[0] === "2") {
    solution2();
  }

  Deno.exit(0);
})();
