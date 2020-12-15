(() => {
  const solution1 = () => {
    const countLetter = (password: string, letter: string) => {
      let letterCount = 0;
      for (let i = 0; i < password.length; i++) {
        if (password[i] === letter) {
          letterCount++;
        }
      }

      return letterCount;
    };

    const content = Deno.readTextFileSync("input.txt");
    const lines = content.split("\n");
    const rules = lines
      .filter((l) => l !== "")
      .map((l) => {
        const split = l.split(" ");
        const minMaxOccurences = split[0].split("-");
        return {
          minOccurence: parseInt(minMaxOccurences[0]),
          maxOccurence: parseInt(minMaxOccurences[1]),
          letter: split[1].replace(":", ""),
          password: split[2],
        };
      });

    let passwordValidated = 0;
    for (const rule of rules) {
      const count = countLetter(rule.password, rule.letter);
      if (count >= rule.minOccurence && count <= rule.maxOccurence) {
        passwordValidated++;
      }
    }

    console.log(`Valid password count: ${passwordValidated}`);
  };

  const solution2 = () => {
    const isValid = (
      password: string,
      firstIndex: number,
      lastIndex: number,
      letter: string,
    ) => {
      const firstLetter = password[firstIndex];
      const lastLetter = password[lastIndex];

      console.log(firstLetter + ' ' + lastLetter);
      if (firstLetter === letter && lastLetter === letter) {
        return false;
      }

      if (firstLetter === letter || lastLetter == letter) {
        return true;
      }

      return false;
    };

    const content = Deno.readTextFileSync("input.txt");
    const lines = content.split("\n");
    const rules = lines
      .filter((l) => l !== "")
      .map((l) => {
        const split = l.split(" ");
        const indexes = split[0].split("-");
        return {
          firstIndex: parseInt(indexes[0]),
          lastIndex: parseInt(indexes[1]),
          letter: split[1].replace(":", ""),
          password: split[2],
        };
      });

    let passwordValidated = 0;
    for (const rule of rules) {
      console.log(rule);
      const isPasswordValid = isValid(
        rule.password,
        rule.firstIndex - 1,
        rule.lastIndex - 1,
        rule.letter,
      );
      if (
        isPasswordValid
      ) {
        console.log(isPasswordValid);
        passwordValidated++;
      }
    }

    console.log(`Valid password count: ${passwordValidated}`);
  };

  if (Deno.args[0] === "1") {
    solution1();
  } else if (Deno.args[0] === "2") {
    solution2();
  }

  Deno.exit(0);
})();
